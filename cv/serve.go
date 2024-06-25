package cv

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

type Host struct {
	capture *gocv.VideoCapture
	stream  *mjpeg.Stream
	device  any
}

func Serve(devices ...any) {
	var (
		err error
		url = "192.168.0.7:9000"
	)
	hosts := make([]Host, len(devices))

	for i, deviceID := range devices {
		h := &hosts[i]
		h.device = deviceID
		// open webcam
		h.capture, err = gocv.OpenVideoCapture(deviceID)
		if err != nil {
			fmt.Printf("Error opening capture device: %v\n", deviceID)
			return
		}
		defer h.capture.Close()

		// create the mjpeg stream
		h.stream = mjpeg.NewStream()

		go mjpegCapture(h.stream, h.capture, deviceID)
		if i == 0 {
			http.Handle("/", h.stream)
		} else {
			http.Handle(fmt.Sprintf("/%d/", i), h.stream)
		}
	}

	fmt.Println("Capturing. Point your browser to " + url)

	server := &http.Server{
		Addr:         url,
		ReadTimeout:  6000 * time.Second,
		WriteTimeout: 6000 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}

func mjpegCapture(stream *mjpeg.Stream, capture *gocv.VideoCapture, deviceID any) {
	var (
		img = gocv.NewMat()
		ok  bool
		err error
		buf *gocv.NativeByteBuffer
	)
	defer img.Close()

	for {
		time.Sleep(time.Millisecond)

		if ok = capture.Read(&img); !ok {
			log.Println("Device closed:", deviceID)
			continue
		}
		if img.Empty() {
			continue
		}

		buf, err = gocv.IMEncode(".jpg", img)
		if err != nil {
			log.Println("IMEncode", err)
			continue
		}

		stream.UpdateJPEG(buf.GetBytes())
		buf.Close()
	}
}

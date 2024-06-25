package cv

import (
	"log"
	"time"

	"gocv.io/x/gocv"
)

func Capture(quit <-chan int, deviceID any, hooks ...Hook) {
	var (
		err     error
		capture *gocv.VideoCapture
		img     gocv.Mat
	)

	capture, err = gocv.OpenVideoCapture(deviceID)
	if err != nil {
		log.Println(err, deviceID, "OpenVideoCapture")
		return
	}

	img = gocv.NewMat()
	defer func() {
		for _, hook := range hooks {
			hook.Quit(0)
		}
		img.Close()
		capture.Close()
	}()

	for {
		select {
		case <-quit:
			return
		default:

			if !capture.Read(&img) {
				log.Println("Device closed:", deviceID)
				continue
			}
			if img.Empty() {
				continue
			}

			for _, hook := range hooks {
				hook.Update(&img)
			}
		}

		time.Sleep(time.Millisecond)
	}

}

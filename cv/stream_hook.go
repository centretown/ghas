package cv

import (
	"log"

	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

type StreamHook struct {
	stream *mjpeg.Stream
}

func NewStreamHook() *StreamHook {
	sh := &StreamHook{}
	sh.stream = mjpeg.NewStream()
	return sh
}

var _ Hook = (*StreamHook)(nil)

func (sh *StreamHook) Update(img *gocv.Mat) {
	buf, err := gocv.IMEncode(".jpg", *img)
	if err != nil {
		log.Println("IMEncode", err)
		return
	}

	sh.stream.UpdateJPEG(buf.GetBytes())
	buf.Close()
}

func (sh *StreamHook) Quit(int) {}

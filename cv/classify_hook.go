package cv

import "gocv.io/x/gocv"

type ClassifyHook struct {
	classifier gocv.CascadeClassifier
}

var _ Hook = (*ClassifyHook)(nil)

func NewClassifyHook() *ClassifyHook {
	ch := &ClassifyHook{}
	ch.classifier = gocv.NewCascadeClassifier()
	return ch
}

func (ch *ClassifyHook) Update(img *gocv.Mat) {
}

func (ch *ClassifyHook) Quit(int) {
	ch.classifier.Close()
}

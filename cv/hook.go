package cv

import "gocv.io/x/gocv"

type Hook interface {
	Update(img *gocv.Mat)
	Quit(int)
}

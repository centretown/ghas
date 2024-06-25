package cv

import "gocv.io/x/gocv"

type Tracker struct {
	Handler  func(int)
	Pos      int
	Max      int
	trackBar *gocv.Trackbar
}

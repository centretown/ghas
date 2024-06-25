package cv

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gocv.io/x/gocv"
)

func Stream(device any, trackers ...*Tracker) {
	var (
		err     error
		capture *gocv.VideoCapture
		window  *gocv.Window
		mat     gocv.Mat
		// width     = 720
		// height    = 480
		blue      = color.RGBA{R: 0, G: 128, B: 255, A: 255}
		thickness = 2
	)

	capture, err = gocv.OpenVideoCapture(device)
	if err != nil {
		log.Println(err, "OpenVideoCapture")
		os.Exit(1)
	}
	window = gocv.NewWindow(fmt.Sprintf("Cam:%v", device))
	mat = gocv.NewMat()
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	log.Println(capture.CodecString())

	for _, tr := range trackers {
		tr.trackBar = window.CreateTrackbar("pan", tr.Max)
		tr.trackBar.SetPos(tr.Pos)
	}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("data/haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	capture.Read(&mat)
	log.Println(mat.Cols(), mat.Rows())
	window.ResizeWindow(mat.Cols(), mat.Rows())
	window.IMShow(mat)
	for {
		window.WaitKey(100)

		for _, tr := range trackers {
			newPos := tr.trackBar.GetPos()
			if tr.Pos != newPos {
				tr.Pos = newPos
				tr.Handler(tr.Pos)
			}
		}

		if !capture.Read(&mat) {
			// gocv.Circle(&mat, image.Point{X: width / 2, Y: height / 2}, height/3, colour, thickness)
			continue
		}

		if mat.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(mat)
		// fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image
		var r image.Rectangle
		// window.GetWindowProperty((go))
		for _, r = range rects {
			gocv.Rectangle(&mat, r, blue, thickness)
		}

		window.IMShow(mat)
	}
}

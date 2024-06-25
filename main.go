package main

import (
	"flag"
	"fmt"
	"ghas/cv"
	"ghas/hass"
	"log"
	"net/http"
	"time"
)

var (
	webCam any = 0
	ipCam  any = "http://192.168.0.25:8080"
)

func main() {
	flag.Parse()
	hass.Get("states/number.pan")
	run()
}

func run() {

	classify := cv.NewClassifyHook()

	webWin := cv.NewWinHook(webCam, &cv.Tracker{Handler: matrix_light, Title: "Brightness", Pos: 50, Max: 100})
	webStream := cv.NewStreamHook()
	http.Handle("/", webStream.Stream)
	webQuit := make(chan int)
	go cv.Capture(webQuit, webCam, classify, webWin, webStream)

	ipWin := cv.NewWinHook(ipCam, &cv.Tracker{Handler: pan, Title: "Pan Camera", Pos: 90, Max: 180})
	ipStream := cv.NewStreamHook()
	http.Handle("/1/", ipStream.Stream)
	ipQuit := make(chan int)
	go cv.Capture(ipQuit, ipCam, classify, ipWin, ipStream)

	url := "192.168.0.7:9000"
	fmt.Println("Capturing. Point your browser to " + url)

	server := &http.Server{
		Addr:         url,
		ReadTimeout:  6000 * time.Second,
		WriteTimeout: 6000 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func pan(value int) {
	cmd := fmt.Sprintf(`{"entity_id": "number.pan", "value": %d}`, value)
	hass.Post("services/number/set_value", cmd)
}

func matrix_light(value int) {
	cmd := fmt.Sprintf(`{"entity_id": "light.led_matrix_24","effect": "rainbow-vertical","brightness_pct": %d}`, value)
	hass.Post("services/light/turn_on", cmd)
}

// func mattype(mat *gocv.Mat) {
// 	image, _ := mat.ToImage()
// }

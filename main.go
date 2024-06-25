package main

// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// defer cancel()
// options := &websocket.DialOptions{}
// options.HTTPHeader = make(http.Header)
// options.HTTPHeader.Add("Authorization", "Bearer "+token)
// c, resp, err := websocket.Dial(ctx, "http://192.168.0.17:8123/api/", options)
// if err != nil {
// 	log.Println(err)
// 	log.Println(resp)
// 	return
// }

// var buffer = make([]byte, 1024)
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
	webcam any = 0
	ipcam  any = "http://192.168.0.25:8080"
	phcam  any = "http://192.168.0.177:8080/browserfs.html"
)

func main() {
	flag.Parse()

	var (
		url = "192.168.0.7:9000"
	)

	// panValue := 0
	// arg0 := flag.Arg(0)
	// log.Println("arg0 =", arg0)
	// fmt.Sscan(arg0, &panValue)
	// log.Println("pan =", panValue)

	// pan(panValue)
	hass.Get("states/number.pan")

	wh := cv.NewWinHook(0, &cv.Tracker{Handler: matrix_light, Pos: 50, Max: 100})
	quit := make(chan int)
	go cv.Capture(quit, 0, wh)

	// go cv.Stream(webcam, &cv.Tracker{Handler: matrix_light, Pos: 50, Max: 100})
	// go cv.Stream(ipcam, &cv.Tracker{Handler: pan, Pos: panValue, Max: 180})
	// // cv.Serve(ipcam, webcam)
	// cv.Serve(webcam)

	server := &http.Server{
		Addr:         url,
		ReadTimeout:  6000 * time.Second,
		WriteTimeout: 6000 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

	ch := make(chan int)
	<-ch
	// comm.Get("config")
	// comm.Get("events")
	// comm.Get("services")
	// comm.Get("calendars")
	// comm.Get("states/sun.sun")
	// comm.Get("states/sun.sun")

	// comm.Post("states/number.pan", `{"state": 15.0}`)

	// comm.Post("services/light/turn_on",
	// 	`{"entity_id": "light.led_strip_24"}`)

	// comm.Post("services/light/turn_on",
	// 	`{"entity_id": "light.led_matrix_24",
	// 	"effect": "rainbow-vertical",
	// 	"brightness_pct": 25}`)

	// comm.Post("services/number/set_value",
	// 	`{"entity_id": "number.pan", "value": "0"}`)
}

func pan(value int) {
	var (
		cmd = fmt.Sprintf(`{"entity_id": "number.pan", "value": %d}`, value)
	)

	hass.Post("services/number/set_value", cmd)
}

func matrix_light(value int) {
	var (
		// comm.Post("services/light/turn_on",
		// 	`{"entity_id": "light.led_matrix_24",
		// 	"effect": "rainbow-vertical",
		// 	"brightness_pct": 25}`)

		cmd = fmt.Sprintf(`{"entity_id": "light.led_matrix_24","effect": "rainbow-vertical","brightness_pct": %d}`, value)
	)

	hass.Post("services/light/turn_on", cmd)
}

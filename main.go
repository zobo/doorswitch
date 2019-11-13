package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime"
	"time"
)

func loopSwitch(pin GPIO_Pin, c <-chan struct{}) {
	for {
		<-c

		pin.Low()
		time.Sleep(time.Second * 3)
		pin.High()
	}
}

func handleSwitch(c chan<- struct{}, res http.ResponseWriter, req *http.Request) {
	// whatever we get

	select {
	case c <- struct{}{}:
		// sent OK
		res.Write([]byte("OK"))
	default:
		// sent error
		res.Write([]byte("Busy"))
	}
}

func main() {

	gpio := GPIO{}

	pin, err := gpio.Pin("2")
	if runtime.GOOS == "windows" {
		log.Println("On windows, ignoring GPIO errors")
	} else {
		if err != nil {
			log.Fatal(err)
		}
		err = pin.Output()
		if err != nil {
			log.Fatal(err)
		}
		// close door default
		err = pin.High()
		if err != nil {
			log.Fatal(err)
		}
	}

	c := make(chan struct{})

	go loopSwitch(pin, c)

	r := mux.NewRouter()
	r.HandleFunc("/switch", func(res http.ResponseWriter, req *http.Request) { handleSwitch(c, res, req) }).Methods("POST")
	r.Handle("/", http.FileServer(http.Dir("."))).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", r))
}

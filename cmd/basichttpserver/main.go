package main

import (
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/goog-lukemc/gotrain"
)

func main() {
	//done := make(chan bool)
	// handle the request to a route
	http.HandleFunc("/ascii/", asciiHandler)
	// catchall default route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up")) // What about err this do we not care?
		return
	})

	// A little golang worker here to help us know what the Http
	// server is ready.  It serves no other purpose other than notification
	go checkhttp()

	// A single line in the http package starts an http server.

	// Set a delay timer so we can see the http server check
	time.Sleep(time.Second * 3)

	// If we wanted we could listen on 2 or more ports at the same time.
	// go http.ListenAndServe(":8081", nil)

	http.ListenAndServe(":8080", nil) // Handle this error

}

// This function handles http request for the json file
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	// Lets get the string at the end of the URL
	scount := path.Clean(path.Base(r.URL.Path))

	// Try to convert it to and int
	count, err := strconv.Atoi(scount)

	// if it fails lets notify the user by send it to the Http
	// response body
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// we use our custom package to get us the data
	data := gotrain.ASCIICoolness{
		Length: count + 1,
	}
	result, err := data.MarshalJSON()
	if err != nil {
		http.Error(w,"Internal Server", http.StatusInternalServerError)
		log.Println(err.Error())
		//w.Write([]byte(err.Error()))
		return
	}

	// Set the header so the call knows its json data
	w.Header().Set("Content-Type", "application/json")

	// write the body
	w.Write(result) // 
	return
}

// Check our HTTP server to see if it up
func checkhttp() {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:8080/")
		if resp != nil {
			if resp.StatusCode == 200 {
				log.Println("Server Is up!")
				return
			}
		}
		if err != nil {
			log.Printf("http check error:%v", err.Error())
		}
		time.Sleep(time.Millisecond * 500)
	}

}

// Exercise 10 minutes:
// What are the problems with this code? Swallowed error etc...?

// Bonus: The golang built in http lib is extremely robust but, over the years, many
// external helper libs have been created. gorilla mux is a good example. Any
// why this happen? How do you go about evaluating external open source libs?

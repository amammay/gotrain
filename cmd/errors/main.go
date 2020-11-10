package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// These are the  urls
var urls = []string{"http://www.google.com", "http://www.walmart.com", "http://fake12wxxxyh.us/"}

func main() {
	// Setup out logger
	log.SetOutput(os.Stdout)

	// Gos For loop - evaluate our urls array
	for i, v := range urls {

		// Display our progress to the ui
		log.Printf("Position:%v Url:%v", i, v)
		resp, err := http.Get(v)

		// TODO: Check for problems - Panic at first we will modify this.
		if err != nil {
			log.Printf("Url %s returns error:%s", v, err.Error())
			panic(err)
			// continue
		}

		// Check our webserver response status
		if resp.StatusCode != 200 {
			log.Printf("Response is not 200. Error is %v", resp.Status)
		}

		// Check that we have an actual response body
		if resp.Body == nil {
			log.Printf("Response Body is Nil for url:%v", v)
		}

		// Defer statements push to the stack ans scheduled to execute after the function completes.
		defer resp.Body.Close()

		u, _ := url.Parse(v)

		// Dynamically create a file name to write to
		fileName := fmt.Sprintf("result-%v.txt", u.Host)

		// Use the os lib to open a write handle to the file.  The call will Create
		// it if it doesn't exist or overwrite it if it does.
		file, err := os.Create(fileName)

		// If we can write the a file we should quit
		if err != nil {
			panic(err)
		}

		// Again we clean up by defering the close
		defer file.Close()

		// Finally copy the results of our http response to the file.
		// os.Copy can handle very ver large files cause it copies the
		// file in sections.
		io.Copy(file, resp.Body)

	}

}

// Exercise: 5 Minutes
// There are 1 or more common go programming mistakes in that code that is not already commented on.
// Can you find them?

// In more complex functions with many defer statements,
// it is important to note that defers are schedeled in last in first out order.

// Bonus: https://golang.org/pkg/errors
// Let's look at and discuss some of the new Error methods added in
// go 1.13

// b := new(bob)
// Simple(b)

// type bob struct {
// }

// func (b bob) Error() string {
// 	return ""
// }

// func Simple(e error) string {
// 	return "I Am An Error indeed"
// }

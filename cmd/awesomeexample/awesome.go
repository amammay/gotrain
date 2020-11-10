// Package awesomeexample has some awesome examples!
package awesomeexample

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, 👋 %s", name)
}

// HelloEverybody will greet everybody you pass in
func HelloEverybody(names []string) []string {
	var greeted []string
	for _, name := range names {
		greeted = append(greeted, fmt.Sprintf("Hello, 👋 %s", name))
	}
	return greeted
}

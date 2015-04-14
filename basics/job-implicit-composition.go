// Go Bootcamp: 3.6

package main

import (
	"log"
	"os"
)

type Job struct {
	Command     string
	*log.Logger // implicit composition here
}

func main() {
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Print("starting now...")
}

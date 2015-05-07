// Building Web Apps with Go: Render Package

package main

import (
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

}

// Building Web Apps with Go: Ex 1

package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./html")))
}

package main

import "net/http"

func main() {
	yoadmin.InitStore()
	defer yoads.CloseStore()
	http.Handle("/protected", httphandler.Protected())
	http.ListenAndServe(":8080", nil)
}

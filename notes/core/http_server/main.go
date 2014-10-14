package main

import (
	"io"
	"net/http"
)

func hello(response http.ResponseWriter, request *http.Request) {
	response.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		response,
		`<doctype html>
		<html>
		<head>
		</head>
		<body>
		Hello World!
		</body>
		</html>
		`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/asserts",
			http.FileServer(http.Dir("assets")),
		),
	)

	http.ListenAndServe(":9000", nil)
}

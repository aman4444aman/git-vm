package main

import (
	"fmt"
	"net/http"
	"os"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hostname</title>
			<style>
				body {
					background-color: lightblue;
					font-family: Arial, sans-serif;
					text-align: center;
					padding-top: 50px;
				}
			</style>
		</head>
		<body>
			<h1>Hostname: %s</h1>
		</body>
		</html>
	`, hostname)
}

func main() {
	http.HandleFunc("/", hostnameHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}

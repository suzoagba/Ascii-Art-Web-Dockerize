package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

	// Check if the number of arguments is correct
if len(os.Args) != 1 {
	fmt.Println("No arguments are needed.")
	fmt.Println("Usage: go run .")
	os.Exit(1)
}

// Catch interrupt signal and exit gracefully
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
go func() {
	sig := <-sigChan
	log.Printf("Signal received (%v). Exiting...\n", sig)
	os.Exit(0)
}()

// Serve home page using GET method
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// Check if the path is correct
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	// Check if the method is correct
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Serve the home page using the static file
	http.ServeFile(w, r, "static/index.html")
})

// Serve ascii-art using POST method
http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
	// Check if the path is correct
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	// Check if the method is correct
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Validate the banner name
	bannerName := r.FormValue("banner")
	if bannerName == "" {
		http.Error(w, "Missing banner", http.StatusBadRequest)
		return
	}
	if bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy" {
		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}

	// Create a new banner from file
	b, err := NewBanner(bannerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the text from the form and unescape it
	inputText := r.FormValue("text")
	if inputText == "" && !r.Form.Has("text") {
		http.Error(w, "Missing text field", http.StatusBadRequest)
		return
	}

	// Replace the Windows line endings with escaped Unix line endings
	inputText = strings.ReplaceAll(inputText, "\r\n", "\\n")

	// Print the text as ASCII art
	asciiArt, err := b.PrintText(inputText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the ascii art to the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(asciiArt))
})

// Start the server
log.Println("Starting server on port 8080...")
err := http.ListenAndServe(":8080", nil)
if err != nil {
	log.Fatal(err)
}
}
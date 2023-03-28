package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var (
	errCannotBeEmpty = errors.New("input cannot be empty")
)

type Ascii struct {
	Art string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandlerWithCode(w, fmt.Errorf("cannot parse input form: %w", err), 400)

		return
	}
	input := r.FormValue("input")
	style := r.FormValue("styles")
	outType := r.FormValue("outputs")
	// return error if input is empty
	if input == "" {
		errorHandlerWithCode(w, errCannotBeEmpty, 400)

		return
	}
	// transform input into output
	out, err := PrintAscii(input, style)
	if err != nil {
		errorHandler(w, err)

		return
	}
	// handle download request
	if outType == "export" {
		exportHandler(w, r, out)

		return
	}
	// print results in website
	output := Ascii{Art: out}
	tmpl := template.Must(template.ParseFiles("static/result.html"))
	err = tmpl.Execute(w, output)
	if err != nil {
		errorHandler(w, err)

		return
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request, out string) {
	// Set the Content-Disposition header to specify that the response should be treated as a file attachment
	w.Header().Set("Content-Disposition", "attachment; filename=example.txt")
	// Set content length
	w.Header().Set("Content-Length", strconv.Itoa(len(out)))
	// Set the Content-Type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Write the file contents to the response
	w.Write([]byte(out))
}

// errorHandler dynamically sets code 404 if NotExist, 400 if Invalid or 500 else.
// Parses error template and responds.
// Return after calling this method.
func errorHandler(w http.ResponseWriter, er error) {
	code := 500
	if errors.Is(er, os.ErrNotExist) {
		code = 404
	} else if errors.Is(er, os.ErrInvalid) {
		code = 400
	}

	errorHandlerWithCode(w, er, code)
}

// errorHandlerWithCode parses error template and responds with given data.
// Return after calling this method.
func errorHandlerWithCode(w http.ResponseWriter, er error, code int) {
	var errMap map[string]any = map[string]any{
		"ErrMsg":     er.Error(),
		"StatusCode": code,
	}
	w.WriteHeader(code)
	html, err := template.ParseFiles("static/error.html")
	if err != nil {
		http.Error(w, "cannot parse error.html template", 404)
		return
	}
	err = html.Execute(w, errMap)
	if err != nil {
		http.Error(w, "cannot map data to template", 500)
		return
	}
}
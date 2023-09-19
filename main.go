package main

import (
	asciiArtWeb "asciiArtWeb/asciiArtFuncs"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	fileserver := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileserver)
	http.HandleFunc("/ascii-art", resultAscii)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	text := r.FormValue("text")
// 	banner := r.FormValue("banner")

// 	bannerResult := asciiArtWeb.ReadBanner(banner)
// 	if strings.Contains(bannerResult, "Error") {
// 		w.WriteHeader(404)
// 		w.Write([]byte(bannerResult))
// 		return
// 	}
// 	formatedToAscii := asciiArtWeb.FormatTextToAscii(text)

// 	fmt.Fprintf(w, "Chosen banner style: %s\n", banner)
// 	fmt.Fprintf(w, "Original text: %s\n", text)
// 	fmt.Fprintf(w, "Text in Ascii Art form: \n%s\n", formatedToAscii)
// }


func resultAscii(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	bannerResult := asciiArtWeb.ReadBanner(banner)
	if strings.Contains(bannerResult, "Error") {
		w.WriteHeader(404)
		w.Write([]byte(bannerResult))
		return
	}
	formatedToAscii := asciiArtWeb.FormatTextToAscii(text)

	fp := path.Join("templates", "ascii_art.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, formatedToAscii); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

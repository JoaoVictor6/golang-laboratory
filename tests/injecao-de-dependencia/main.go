package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func Cumprimenta(w io.Writer, nome string) {
	fmt.Fprintf(w, "Olá, %v", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "mundo")
}

func main() {
	fmt.Printf(NoticeColor, "Server is running in http://localhost:5000")
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		fmt.Println(ErrorColor, err)
	}
}

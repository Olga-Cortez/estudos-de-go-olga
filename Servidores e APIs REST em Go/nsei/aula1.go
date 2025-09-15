package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ola\n")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalhos)

	http.ListenAndServe(":8090", nil)
}
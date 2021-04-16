package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	h.l.Println("Hello!")
	data, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		//rw.WriteHeader(http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s", data)
	fmt.Fprintf(rw, "Hello %s", data)
}

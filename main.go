package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calc", handleCalc)
	http.ListenAndServe(":8090", nil)

}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	num := r.FormValue("n")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	io.WriteString(w, strconv.Itoa(factorial(n)))

}

func factorial(n int) int {
	return n * n
}

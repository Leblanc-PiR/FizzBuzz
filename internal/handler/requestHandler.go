package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Leblanc-PiR/FizzBuzz/internal/data"
	"github.com/Leblanc-PiR/FizzBuzz/internal/service"
)

// Health allows user to check at any time that the server is up and running (httpCode 200)
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetFizzBuzz receives and gives the awaited type to forwarded params
func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	// init
	var (
		int1, int2, lim int
		err             error
	)

	// Get ints
	int1, int2, lim = GetUrlInt(w, r, "int1"), GetUrlInt(w, r, "int2"), GetUrlInt(w, r, "limit")

	if (int1 < 0 || int1 > lim) ||
		(int2 < 0 || int2 > lim) ||
		lim < 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)

		err = fmt.Errorf(`
		unusable values given.
		Reminder: lim > str1, str2 and all of them must be greater than 0
		Usage: fizzbuzz?int1=2&int2=3&lim=100&str1=fizz&str2=buzz`)

		w.Write([]byte(err.Error()))
		return
	}

	// Get strings
	str1, str2 := r.URL.Query().Get("str1"), r.URL.Query().Get("str2")

	res := service.FizzBuzz(int1, int2, lim, str1, str2)
<<<<<<< HEAD

	// Record call
	data.RecordFizzBuzzCall(int1, int2, lim, str1, str2)
=======
	spew.Dump(res)
>>>>>>> 42ea955 (feat: :hammer: fizzbuzz ok, WIP stats)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

// GetUrlInt gives int from URL param
func GetUrlInt(w http.ResponseWriter, r *http.Request, param string) int {
	strParam := r.URL.Query().Get(param)

	res, err := strconv.Atoi(strParam)
	if err != nil {
		w.Write([]byte(err.Error()))
		res = -1
	}

	return res
}

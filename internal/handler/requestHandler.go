package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Leblanc-PiR/FizzBuzz/internal/data"
	"github.com/Leblanc-PiR/FizzBuzz/internal/service"
)

// exposedHighestHitsRequest respresent the answer of /stats endpoint
type exposedHighestHitsRequest struct {
	FormattedRequest string `json:"formattedRequest"`
	Hits             int    `json:"hits"`
}

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
	int1, int2, lim = getUrlInt(w, r, data.Int1ParamStr),
		getUrlInt(w, r, data.Int2ParamStr),
		getUrlInt(w, r, data.LimitParamStr)

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
	str1, str2 := r.URL.Query().Get(data.Str1ParamStr),
		r.URL.Query().Get(data.Str2ParamStr)

	res := service.FizzBuzz(int1, int2, lim, str1, str2)

	// Record call
	data.RecordFizzBuzzCall(int1, int2, lim, str1, str2)

	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(res)
}

// GetStats return param of request with most hits
func GetStats(w http.ResponseWriter, r *http.Request) {

	res := []exposedHighestHitsRequest{}
	for _, foundRequests := range data.GetHighestHitsRequestParams() {
		newEHHR := exposedHighestHitsRequest{
			FormattedRequest: foundRequests.FormattedRequest,
			Hits:             foundRequests.Hits,
		}
		res = append(res, newEHHR)
	}

	w.WriteHeader(http.StatusOK)

	exposeEncodedValues(w, res)
}

// getUrlInt obtains int from URL param
func getUrlInt(w http.ResponseWriter, r *http.Request, param string) int {
	strParam := r.URL.Query().Get(param)

	if param == "" {
		exposeEncodedValues(w, "no param given")
		return -1
	}

	res, err := strconv.Atoi(strParam)
	if err != nil {
		exposeEncodedValues(w, fmt.Sprintf("could not parse %s to int", strParam))
		return -1
	}

	return res
}

// exposeEncodedValues encodes and translates any ascii code that might happen
func exposeEncodedValues(w http.ResponseWriter, res any) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(res)
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleCalc(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "zero",
			num:  5,
			want: 25,
		},
		{
			name: "one",
			num:  12,
			want: 144,
		},
		{
			name: "two",
			num:  100,
			want: 10000,
		},
	}

	handler := http.HandlerFunc(handleCalc)

	for _, cases := range testCases {
		t.Run(cases.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", fmt.Sprintf("/calc?n=%d", cases.num), nil)
			handler.ServeHTTP(rec, request)
			got, _ := strconv.Atoi(string(rec.Body.Bytes()))
			assert.Equal(t, cases.want, got)
		})
	}

}

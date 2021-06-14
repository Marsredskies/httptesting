package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleCalc(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want []byte
	}{
		{
			name: "zero",
			num:  5,
			want: []byte("25"),
		},
		{
			name: "one",
			num:  12,
			want: []byte("144"),
		},
		{
			name: "two",
			num:  100,
			want: []byte("10000"),
		},
	}

	handler := http.HandlerFunc(handleCalc)

	for _, cases := range testCases {
		t.Run(cases.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", fmt.Sprintf("/calc?n=%d", cases.num), nil)
			handler.ServeHTTP(rec, request)
			assert.Equal(t, cases.want, rec.Body.Bytes())
		})
	}

}

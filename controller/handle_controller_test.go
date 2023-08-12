package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestShowFibNumber(t *testing.T) {
	router := gin.Default()
	router.GET("/fib", ShowFibNumber)

	test := []struct {
		query    string
		expected string
		status   int
	}{
		{"n=0", "", 400},
		{"n=1", "1", 200},
		{"n=5", "5", 200},
		{"n=10", "55", 200},
		{"n=-1", "", 400},
		{"n=0.5", "", 400},
		{"n=250001", "", 400},
	}

	for _, tc := range test {
		t.Run(tc.query, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/fib?"+tc.query, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tc.status {
				t.Errorf("Expected status %d; got %d", tc.status, rec.Code)
			}

			if rec.Code == 200 && rec.Body.String() != `{"result":`+tc.expected+`}` {
				t.Errorf("Expected response %s; got %s", `{"result":`+tc.expected+`}`, rec.Body.String())
			}
		})
	}
}

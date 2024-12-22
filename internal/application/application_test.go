package application

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name          string
		expression    string
		expectedCode  int
		expectedBody  string
	}{
		{
			name:          "Valid Expression",
			expression:    `{"expression": "2+2"}`,
			expectedCode:  http.StatusOK,
			expectedBody:  `{"result":4}`,
		},
		{
			name:          "Invalid Expression",
			expression:    `{"expression": "2+2*"}`,
			expectedCode:  http.StatusUnprocessableEntity,
			expectedBody:  `{"error":"Expression is not valid"}`,
		},
		{
			name:          "Internal Server Error - Invalid JSON",
			expression:    `{"expresion": "2+2"`,
			expectedCode:  http.StatusInternalServerError,
			expectedBody:  `{"error":"Internal server error"}`,
		},
		{
			name:          "Internal Server Error - Empty Body",
			expression:    ``,
			expectedCode:  http.StatusInternalServerError,
			expectedBody:  `{"error":"Internal server error"}`,
		},
		{
			name:          "Complex Expression",
			expression:    `{"expression": "2+2*2/2"}`,
			expectedCode:  http.StatusOK,
			expectedBody:  `{"result":4}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(tt.expression)))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedCode)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/truth", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	type Response struct {
		Ans int `json:"ans"`
		Message string `json:"message"`
	}
	var data Response
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Error(err)
	}
	truthMessage := data.Message
	assert.Equal(t, "2 + 2 = 4. (if it is not 4 then we are in trouble)", truthMessage)

	cases := [] struct {
		Name string
		Code int
		N1 string
		N2 string
		Ans int
		Message string
	}{
		{
			Name: "1+1",
			Code: 200,
			N1: "1",
			N2: "1",
			Ans: 2,
		},
		{
			Name: "Invalid values in input 1",
			Code: 422,
			N1: "a",
			N2: "1",
			Message: "Invalid arguments",
		},
		{
			Name: "Invalid values in input 2",
			Code: 422,
			N1: "1",
			N2: "a",
			Message: "Invalid arguments",
		},
		{
			Name: "Invalid values in input 3",
			Code: 422,
			N1: "a",
			N2: "a",
			Message: "Invalid arguments",
		},
		{
			Name: "Invalid values in input 4",
			Code: 422,
			N1: "",
			N2: "",
			Message: "Invalid arguments",
		},
	}
	
	for _, tt := range cases {
		t.Run(tt.Name, func(t2 *testing.T) {
			t2.Parallel()
			w := httptest.NewRecorder()
			PostBody := strings.NewReader("{n1:" +tt.N1+ ",n2:" +tt.N2+ "}")
			req, _ := http.NewRequest("POST", "/addTwoNumbers", PostBody)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.Code, w.Code)

			var data Response
			err := json.Unmarshal(w.Body.Bytes(), &data)
			if err != nil {
				t.Error(err)
			}
			ans := data.Ans
			message := data.Message
			assert.Equal(t, ans, tt.Ans)
			assert.Equal(t, message, tt.Message)
		})
	}
}

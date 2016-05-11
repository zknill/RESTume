package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestHobbies(t *testing.T) {

	var responseRecorder *httptest.ResponseRecorder
	var request http.Request

	responseRecorder = httptest.NewRecorder()

	Hobbies(responseRecorder, &request)

	body := responseRecorder.Body.Bytes()
	response := map[string]interface{}{}
	json.Unmarshal(body, &response)

	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, response["Hobbies"], []interface{}{"Running", "Marathons", "Code"})

}

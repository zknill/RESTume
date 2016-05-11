package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbout(t *testing.T) {

	var responseRecorder *httptest.ResponseRecorder
	var request http.Request

	responseRecorder = httptest.NewRecorder()

	About(responseRecorder, &request)

	body := responseRecorder.Body.Bytes()
	response := map[string]interface{}{}
	json.Unmarshal(body, &response)

	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, response["Job"], "Developer")
	assert.Equal(t, response["Name"], "Zak")
}

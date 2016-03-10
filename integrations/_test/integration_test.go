package integration_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/cloudnativego/secureservice/server"
	"github.com/codegangsta/negroni"
)

var (
	server   *negroni.Negroni
	recorder *httptest.ResponseRecorder
)

const (
	validKey   = "FLUFFY"
	invalidKey = "INVALID_KEY"
	postAPI    = "/api/post"
	getAPI     = "/api/get"
)

func TestIntegration(t *testing.T) {

	// Set API_KEY environment variable
	os.Setenv(APIKey, validKey)
	server = NewServer()

	// Home page can be accessed without API_KEY
	getHomePageRequest, _ := http.NewRequest("GET", "/", nil)
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, getHomePageRequest)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected response code to be %d, received: %d", http.StatusOK, recorder.Code)
	}

	// GET /api/get without API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("GET", getAPI, false, false))
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code to be %d, received %d", http.StatusUnauthorized, recorder.Code)
	}

	// GET /api/get with invalid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("GET", getAPI, true, false))
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code to be %d, received %d", http.StatusUnauthorized, recorder.Code)
	}

	// GET /api/get with valid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("GET", getAPI, true, true))
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected response code to be %d, received %d", http.StatusOK, recorder.Code)
	}

	// POST /api/get with invalid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("POST", getAPI, true, false))
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code to be %d, received %d", http.StatusUnauthorized, recorder.Code)
	}

	// POST /api/get with valid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("POST", getAPI, true, true))
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected response code to be %d, received %d", http.StatusNotFound, recorder.Code)
	}

	// POST /api/post without API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("POST", postAPI, false, false))
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code to be %d, received %d", http.StatusUnauthorized, recorder.Code)
	}

	// POST /api/post with invalid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("POST", postAPI, true, false))
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code to be %d, received %d", http.StatusUnauthorized, recorder.Code)
	}

	// POST /api/post with valid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("POST", postAPI, true, true))
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected response code to be %d, received %d", http.StatusOK, recorder.Code)
	}

	// GET /api/post with invalid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("GET", postAPI, true, false))
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected response code to be %d, received %d", http.StatusUnauthorized, recorder.Code)
	}

	// GET /api/post with valid API_KEY
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, createRequest("GET", postAPI, true, true))
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected response code to be %d, received %d", http.StatusNotFound, recorder.Code)
	}

}

func createRequest(method, api string, hasKey, valid bool) (req *http.Request) {
	req, _ = http.NewRequest(method, api, nil)
	if hasKey {
		if valid {
			req.Header.Add(APIKey, validKey)
		} else {
			req.Header.Add(APIKey, invalidKey)
		}
	}
	return
}

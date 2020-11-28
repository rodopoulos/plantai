package server //nolint:testpackage

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
)

var server *httptest.Server

func Test_GetHealthcheck(t *testing.T) {
	res, err := http.Get(server.URL + "/healthcheck") //nolint:noctx
	if err != nil {
		t.Errorf("Failed to request server, got: %v", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected HTTP code 200, but got %v", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Errorf("Failed to read server response, got: %v", err)
	}
	payload := string(body)
	if payload != "WORKING" {
		t.Errorf("Expected response string WORKING, but got %s", payload)
	}
}

func TestMain(m *testing.M) {
	log, _ := test.NewNullLogger()
	router := buildRouter(queue)
	server = httptest.NewServer(router)
	status := m.Run()
	server.Close()
	os.Exit(status)
}

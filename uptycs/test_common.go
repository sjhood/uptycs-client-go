package uptycs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jarcoal/httpmock"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func RespFromFixture(fixtureFile string) (*http.Response, error) {
	var fixture map[string]interface{}
	fixture_str, err := ioutil.ReadFile(fixtureFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fixture_str, &fixture)
	if err != nil {
		return nil, err
	}

	resp, err := httpmock.NewJsonResponse(200, fixture)
	return resp, err
}

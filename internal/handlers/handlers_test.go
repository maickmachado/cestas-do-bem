package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	//TODO: preencher com o resto
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"beneficiary-register-get", "/beneficiary-register", "GET", []postData{}, http.StatusOK},
	{"beneficiary-register-post", "/beneficiary-register", "POST", []postData{
		{key: "first_name", value: "maick"},
		{key: "email", value: "maick.armachado@gmail.com"},
	}, http.StatusOK},
}

func TestHandler(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %v, expected %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, v := range e.params {
				values.Add(v.key, v.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %v, expected %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

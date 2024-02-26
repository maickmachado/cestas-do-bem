package setup_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/maickmachado/cestas-do-bem/internal/middlewares"
)

func TestNoSurf(t *testing.T) {
	var myH MyHandler

	h := middlewares.NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Printf("TestNoSurf, Don't know type %T\n", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH MyHandler

	h := middlewares.SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Printf("TestSessionLoad, Don't know type %T\n", v))
	}
}

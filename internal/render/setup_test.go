package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/models"
)

var sessionManager *scs.SessionManager
var testApp config.AppConfig

// TestMain runs in the main goroutine and can do whatever setup and teardown is necessary around a call to m.Run. It should then call os.Exit with the result of m.Run
// could be done as the "getRoutes" function at handler test
func TestMain(m *testing.M) {

	//what am I going to put in session
	gob.Register(models.Register{})

	//change it to true when in production
	testApp.InProduction = false

	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false
	// add the session to the AppConfig
	testApp.Session = sessionManager

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/markbates/pkger/cmd/pkger"
	"github.com/pkg/browser"
	log "github.com/sirupsen/logrus"
)

var directoryPath string
var listeningAddr string

func init() {
    directoryPath = "./web/dist"
    listeningAddr = ":8080"

    // set up logrus configuration
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
}

func Serve() {

	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	router := newRouter()

    server := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(WebLogger),
	)
	server.UseHandler(router)

    log.Infof("Listening on address: %s", listeningAddr)
	log.Infof("Serving Path: %s", directoryPath)

	browser.OpenURL("http://localhost"+listeningAddr)

	log.Fatal(http.ListenAndServe(listeningAddr, server))

	//http.ListenAndServe(":8080", server)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := pkger.Dir(directoryPath) // http.Dir(directoryPath)
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}

func WebLogger(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, req)
	res := rw.(negroni.ResponseWriter)
	defer func() {
		elapsed := time.Since(start)
		log.WithFields(log.Fields{
			"elapsed": elapsed,
			"method":  req.Method,
			"host":    req.URL.Host,
			"path":    req.URL.Path,
			"query":   req.URL.RawQuery,
			"status":  res.Status(),
			"size":    res.Size(),
		}).Info(req.Method + " " + req.URL.Path)
	}()
}
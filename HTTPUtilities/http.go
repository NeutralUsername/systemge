package HTTPUtilities

import (
	"fmt"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
}

func CreateHttpServer(port string, mux *http.ServeMux) *HTTPServer {
	server := &HTTPServer{
		server: &http.Server{
			Addr:    port,
			Handler: mux,
		},
	}
	return server
}

func (server *HTTPServer) SetHandlerFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	server.server.Handler.(*http.ServeMux).HandleFunc(pattern, handler)
}

func (server *HTTPServer) StartHTTP() {
	server.server.ListenAndServe()
}

func (server *HTTPServer) StartHTTPS(certPath string, keyPath string) {
	err := server.server.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		fmt.Println(err)
	}
}

func (server *HTTPServer) Shutdown() {
	server.server.Shutdown(nil)
	server = nil
}

func (server *HTTPServer) Close() {
	server.server.Close()
	server.server = CreateHttpServer(server.server.Addr, server.server.Handler.(*http.ServeMux)).server
}

func SendDirectory(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(path)).ServeHTTP(w, r)
	}
}

func Redirect(toURL string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, toURL, http.StatusMovedPermanently)
	}
}

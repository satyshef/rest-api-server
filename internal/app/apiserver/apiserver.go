package apiserver

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *http.Server
}

// Create new APIServer instance
func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
		logger: logrus.New(),
		router: &http.Server{},
	}
}

// Run server
func (a *APIServer) Run() error {
	if err := a.initLogger(); err != nil {
		return err
	}
	a.logger.Info("Starting API server...")
	a.initRouter()
	return a.router.ListenAndServe()
}

// Configure logger
func (a *APIServer) initLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)
	return nil
}

// Configure router
func (a *APIServer) initRouter() {
	a.router.Addr = a.config.Addr
	a.router.ReadTimeout = time.Duration(a.config.ReadTimeout) * time.Second
	a.router.WriteTimeout = time.Duration(a.config.WriteTimeout) * time.Second
	a.router.MaxHeaderBytes = 1 << 20 // 1 MB
	a.router.Handler = a.handlerMain()
}

// Main handler
func (a *APIServer) handlerMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := fmt.Sprintf("Method: %s\n", r.Method)
		result += fmt.Sprintf("Request: %s\n", r.URL.Path)
		result += fmt.Sprintf("Header: %#v\n", r.Header)
		result += "Params:\n"
		m, _ := url.ParseQuery(r.URL.RawQuery)
		for key, value := range m {
			result += fmt.Sprintf("        %s = %s\n", key, value)
		}
		io.WriteString(w, result)
	}
}

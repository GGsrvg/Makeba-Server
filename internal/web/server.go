package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	dir    string
	config *Config
	logger *logrus.Logger
	router *mux.Router
	// store  *store.Store
}

func New(dir string, config *Config) *Server {
	return &Server{
		dir:    dir,
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting web server")

	return http.ListenAndServe(s.config.HttpAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLebel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	// return static files
	// html, css, js, pdf, mp4 and more
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(s.dir+"/internal/web/static"))))

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "./internal/web/static/html/main.html")
	}).Methods("GET")
}

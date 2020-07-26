package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"makeba/internal/util"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	dir string
)

func init() {
	cdir, err := util.CurrentDir()

	if err != nil {
		log.Fatal(err)
	}

	dir = cdir
}

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	// store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
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

	// if err := s.configureStore(); err != nil {
	// 	return err
	// }

	s.logger.Info("starting api server")

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
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static"))))

	// return web page
	s.router.HandleFunc("/", s.getIndexHandle()).Methods("GET")
	// s.router.HandleFunc("/", s.postIndexHandle()).Methods("POST")
}

func (s *Server) apiHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\r%s\r", b)
		io.WriteString(w, "Api")
	}
}

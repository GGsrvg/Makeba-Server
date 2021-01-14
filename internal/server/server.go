package server

import (
	"log"
	"net"
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

	s.logger.Info("starting api server")
	s.logger.Info("IP: " + getOutboundIP().String())

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
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(s.dir+"/static"))))

	// return api

	// return stats
	s.router.HandleFunc("/", s.getIndexHandle()).Methods("GET")

	// return required headers
	// s.router.HandleFunc("/headers", s.getHeadersHandle()).Methods("GET")

	// return settings
	// s.router.HandleFunc("/settings", s.getSettingsHandle()).Methods("GET")

	// return token authorize if need request success
	s.router.HandleFunc("/authorize/", s.postAuthorizeHandle()).Methods("POST")
}

// func (s *Server) apiHandle() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		b, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("\r%s\r", b)
// 		io.WriteString(w, "Api")
// 	}
// }

func getOutboundIP() *net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Panic(err)
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return &localAddr.IP
}

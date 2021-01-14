package server

import (
	"encoding/json"
	"fmt"
	"io"
	"makeba/internal/server/mocks"
	"makeba/internal/server/models"
	"makeba/internal/server/models/authorization"
	"makeba/internal/server/models/headers"
	"net/http"
)

var (
	token string
)

func init() {
	token = "V3R1_S3CR3T_T0K3N"
}

// return list data
func (s *Server) getIndexHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := headers.DetectBaseHeader(&r.Header)

		var response *models.Response

		if err != nil {
			errorMessage := err.Error()
			response = models.NewRespose(nil, &errorMessage, nil)

			w.WriteHeader(http.StatusUnauthorized)
		} else {
			arrayStatsContainer := []*models.StatsContainer{
				models.New(
					"Disk space",
					mocks.GetDiskSpace(),
				),
				models.New(
					"Diograms",
					mocks.GetDiagram(),
				),
			}

			response = models.NewRespose(nil, nil, arrayStatsContainer)
		}

		byte, err := json.Marshal(response)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(byte))
	}
}

// P.S. NOT NEED
// func (s *Server) getHeadersHandle() http.HandlerFunc {
// 	headers := []string{
// 		"TimeZone",
// 		"Locale",
// 		"Authorization",
// 	}

// 	byte, err := json.Marshal(headers)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, string(byte))
// 	}
// }

// func (s *Server) getSettingsHandle() http.HandlerFunc {

// 	settings := make(map[string]string)

// 	settings[""] = ""

// 	byte, err := json.Marshal(settings)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, string(byte))
// 	}
// }

func (s *Server) postAuthorizeHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var authorize authorization.AuthorizationRequest
		err := decoder.Decode(&authorize)
		if err != nil {
			fmt.Println(err)
		}

		if authorize.Login != "I" {
			io.WriteString(w, "Login not correct")
		} else if authorize.Password != "123" {
			io.WriteString(w, "Login or Password not correct")
		} else {
			response := authorization.NewAuthorizationResponse(token)
			byte, err := json.Marshal(response)

			if err != nil {
				fmt.Println(err)
			}

			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w, string(byte))
		}
	}
}

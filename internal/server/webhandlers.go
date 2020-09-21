package server

import (
	"encoding/json"
	"fmt"
	"io"
	"makeba/internal/diskhelper"
	"makeba/internal/server/models"
	"net/http"
	"strconv"
)

// return list data
func (s *Server) getIndexHandle() http.HandlerFunc {

	disk := diskhelper.DiskUsage("/")
	// fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(diskhelper.GB))
	// fmt.Printf("Avail: %.2f GB\n", float64(disk.Avail)/float64(diskhelper.GB))
	// fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(diskhelper.GB))

	return func(w http.ResponseWriter, r *http.Request) {
		GB := " GB"
		allSpace := strconv.FormatFloat(float64(disk.All)/float64(diskhelper.GB), 'f', 2, 64) + GB
		availSpace := strconv.FormatFloat(float64(disk.Avail)/float64(diskhelper.GB), 'f', 2, 64) + GB
		usedSpace := strconv.FormatFloat(float64(disk.Used)/float64(diskhelper.GB), 'f', 2, 64) + GB

		arrayStatsContainer := []*models.StatsContainer{
			models.New(
				"Disk space",
				[]models.Stat{
					*models.NewStatText(
						"All",
						*models.NewText(
							allSpace,
							models.Horizontal,
							models.Center,
						),
					),
					*models.NewStatText(
						"Avail",
						*models.NewText(
							availSpace,
							models.Horizontal,
							models.Center,
						),
					),
					*models.NewStatText(
						"Used",
						*models.NewText(
							usedSpace,
							models.Horizontal,
							models.Center,
						),
					),
				},
			),
		}

		response := models.NewRespose(nil, nil, arrayStatsContainer)

		byte, err := json.Marshal(response)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(byte))
	}
}

func (s *Server) getHeadersHandle() http.HandlerFunc {
	headers := []string{
		"TimeZone",
		"CurrentDate",
		"Locale",
		"Authorization",
	}

	byte, err := json.Marshal(headers)

	if err != nil {
		fmt.Println(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string(byte))
	}
}

func (s *Server) getSettingsHandle() http.HandlerFunc {

	settings := make(map[string]string)

	settings[""] = ""

	byte, err := json.Marshal(settings)

	if err != nil {
		fmt.Println(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string(byte))
	}
}

func (s *Server) postAuthorizeHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Alpha")
	}
}

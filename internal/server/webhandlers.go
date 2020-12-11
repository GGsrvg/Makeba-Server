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
						models.NewText(
							allSpace,
							// models.Horizontal,
							// models.Center,
						),
					),
					*models.NewStatText(
						"Avail",
						models.NewText(
							availSpace,
							// models.Horizontal,
							// models.Center,
						),
					),
					*models.NewStatText(
						"Used",
						models.NewText(
							usedSpace,
							// models.Horizontal,
							// models.Center,
						),
					),
				},
			),
			models.New(
				"Diograms",
				[]models.Stat{
					*models.NewStatDiogram(
						"TIOBE Index for November 2020",
						models.NewDiagram(
							*models.NewDiagramItem(
								"C",
								*models.NewColor(149, 219, 72),
								10,
							),
							*models.NewDiagramItem(
								"Python",
								*models.NewColor(68, 238, 242),
								10,
							),
							*models.NewDiagramItem(
								"Java",
								*models.NewColor(242, 177, 68),
								10,
							),
							*models.NewDiagramItem(
								"C++",
								*models.NewColor(232, 65, 123),
								10,
							),
							*models.NewDiagramItem(
								"Other",
								*models.NewColor(123, 49, 245),
								10,
							),
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

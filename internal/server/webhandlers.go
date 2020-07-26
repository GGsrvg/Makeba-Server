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

		arrayStatsContainer := []*models.StatsContainer{
			models.New(
				"Disk space",
				[]models.Stat{
					*models.NewStatText(
						"All",
						*models.NewText(
							strconv.FormatFloat(float64(disk.All)/float64(diskhelper.GB), 'f', -1, 64),
							models.Horizontal,
							models.Center,
						),
					),
					*models.NewStatText(
						"Avail",
						*models.NewText(
							strconv.FormatFloat(float64(disk.Avail)/float64(diskhelper.GB), 'f', -1, 64),
							models.Horizontal,
							models.Center,
						),
					),
					*models.NewStatText(
						"Used",
						*models.NewText(
							strconv.FormatFloat(float64(disk.Used)/float64(diskhelper.GB), 'f', -1, 64),
							models.Horizontal,
							models.Center,
						),
					),
				},
			),
		}

		byte, err := json.Marshal(arrayStatsContainer)

		if err != nil {
			fmt.Println(err)
		}

		io.WriteString(w, string(byte))
	}
}

func (s *Server) postIndexHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "index")
	}
}

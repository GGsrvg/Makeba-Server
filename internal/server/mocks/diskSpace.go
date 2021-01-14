package mocks

import (
	"makeba/internal/diskhelper"
	"makeba/internal/server/models"
	"strconv"
)

func GetDiskSpace() []models.Stat {
	disk := diskhelper.DiskUsage("/")
	GB := " GB"

	allSpace := strconv.FormatFloat(float64(disk.All)/float64(diskhelper.GB), 'f', 2, 64) + GB
	availSpace := strconv.FormatFloat(float64(disk.Avail)/float64(diskhelper.GB), 'f', 2, 64) + GB
	usedSpace := strconv.FormatFloat(float64(disk.Used)/float64(diskhelper.GB), 'f', 2, 64) + GB

	stats := []models.Stat{
		*models.NewStatText(
			"All",
			models.NewText(
				allSpace,
			),
		),
		*models.NewStatText(
			"Avail",
			models.NewText(
				availSpace,
			),
		),
		*models.NewStatText(
			"Used",
			models.NewText(
				usedSpace,
			),
		),
	}

	return stats
}

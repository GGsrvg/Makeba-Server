package mocks

import (
	"makeba/internal/diskhelper"
	"makeba/internal/server/models"
	"strconv"
)

func GetDiagramSausage() []models.Stat {

	disk := diskhelper.DiskUsage("/")
	GB := " GB"

	availSpace := float64(disk.Avail) / float64(diskhelper.GB)
	availSpaceUInt := uint(availSpace)

	usedSpace := float64(disk.Used) / float64(diskhelper.GB)
	usedSpaceUInt := uint(usedSpace)

	allSpaceString := strconv.FormatFloat(float64(disk.All)/float64(diskhelper.GB), 'f', 2, 64) + GB
	availSpaceString := strconv.FormatFloat(availSpace, 'f', 2, 64) + GB
	usedSpaceString := strconv.FormatFloat(usedSpace, 'f', 2, 64) + GB

	stats := []models.Stat{
		*models.NewStatDiogram(
			"All disk space "+allSpaceString,
			models.NewDiagram(
				models.SAUSAGE,
				*models.NewDiagramItem(
					usedSpaceString,
					*models.NewColor(86, 200, 86),
					usedSpaceUInt,
				),
				*models.NewDiagramItem(
					availSpaceString,
					*models.NewColor(210, 210, 210),
					availSpaceUInt,
				),
			),
		),
	}

	return stats
}

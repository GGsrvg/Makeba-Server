package mocks

import "makeba/internal/server/models"

func GetDiagram() []models.Stat {
	stats := []models.Stat{
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
	}

	return stats
}

package main

import "time"

const DateFormat = "2006-01-02"

func sampleData(from time.Time) []Swimlane {
	return []Swimlane{
		{
			Name: "Backend",
			Items: []Bar{
				{
					Start: from.Format(DateFormat),
					End:   from.Add(time.Hour * 24 * 14).Format(DateFormat),
					GanttBarConfig: GanttBarConfig{
						Id:         "item-3",
						Label:      "Milestone support",
						HasHandles: true,
					},
				},
				{
					Start: from.Add(time.Hour * 24 * 20).Format(DateFormat),
					End:   from.Add(time.Hour * 24 * 34).Format(DateFormat),
					GanttBarConfig: GanttBarConfig{
						Id:         "item-4",
						Label:      "Move across lanes",
						HasHandles: true,
					},
				},
			},
		},
		{
			Name: "Frontend",
			Items: []Bar{
				{
					Start: from.Add(time.Hour * 24 * 5).Format(DateFormat),
					End:   from.Add(time.Hour * 24 * 14).Format(DateFormat),
					GanttBarConfig: GanttBarConfig{
						Id:         "item-1",
						Label:      "Context menu",
						HasHandles: true,
					},
				},
				{
					Start: from.Add(time.Hour * 24 * 21).Format(DateFormat),
					End:   from.Add(time.Hour * 24 * 36).Format(DateFormat),
					GanttBarConfig: GanttBarConfig{
						Id:         "item-2",
						Label:      "Adjust daterange",
						HasHandles: true,
					},
				},
			},
		},
	}
}

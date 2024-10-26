package main

type Swimlane struct {
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
	Items []Bar  `json:"items"`
}

type Bar struct {
	Start          string         `json:"start"`
	End            string         `json:"end"`
	GanttBarConfig GanttBarConfig `json:"ganttBarConfig"`
}

type GanttBarConfig struct {
	Id         string `json:"id"`
	Label      string `json:"label"`
	HasHandles bool   `json:"hasHandles"`
}

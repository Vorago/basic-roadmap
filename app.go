package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx   context.Context
	state []Swimlane
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		state: []Swimlane{
			{
				Name: "Backend",
				Items: []Bar{
					{
						Start: "2024-11-01",
						End:   "2024-11-15",
						GanttBarConfig: GanttBarConfig{
							Id:         "item-3",
							Label:      "Milestone support",
							HasHandles: true,
						},
					},
					{
						Start: "2024-10-01",
						End:   "2024-10-15",
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
						Start: "2024-10-16",
						End:   "2024-10-30",
						GanttBarConfig: GanttBarConfig{
							Id:         "item-1",
							Label:      "Context menu",
							HasHandles: true,
						},
					},
					{
						Start: "2024-11-01",
						End:   "2024-11-15",
						GanttBarConfig: GanttBarConfig{
							Id:         "item-2",
							Label:      "Adjust daterange",
							HasHandles: true,
						},
					},
				},
			},
		},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) UpdateSwimlanes(swimlanes []Swimlane) {
	a.state = swimlanes
	a.saveToFile()
}

func (a *App) LoadFromFile() []Swimlane {
	usr, _ := user.Current()
	dir := usr.HomeDir
	dialog, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: dir,
		DefaultFilename:  "roadmap.json",
		Title:            "Load roadmap from file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "JSON files",
				Pattern:     "*.json",
			},
		},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	content, err := os.ReadFile(dialog)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []Swimlane
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	a.state = payload
	return a.state
}

func (a *App) saveToFile() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	dialog, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           dir,
		DefaultFilename:            "roadmap.json",
		Title:                      "Save roadmap to file",
		Filters:                    nil,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: false,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	file, _ := json.MarshalIndent(a.state, "", " ")

	_ = os.WriteFile(path.Join(dialog), file, 0644)
}

func (a *App) GetSwimlanes() []Swimlane {
	return a.state
}

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

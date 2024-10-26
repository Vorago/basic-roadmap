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
	state *State
}

type State struct {
	Swimlanes []Swimlane `json:"swimlanes"`
	Config    config     `json:"config"`
}

// NewApp creates a new App application struct
func NewApp(conf *config) *App {
	return &App{
		state: &State{
			Swimlanes: sampleData(conf.requireStartDate()),
			Config:    *conf,
		},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) UpdateSwimlanes(swimlanes []Swimlane) {
	a.state.Swimlanes = swimlanes
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

	a.state.Swimlanes = payload
	return a.state.Swimlanes
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
	return a.state.Swimlanes
}

func (a *App) GetState() *State {
	return a.state
}

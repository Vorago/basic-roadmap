package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type config struct {
	StartDate string `toml:"start_date" json:"startDate"`
	EndDate   string `toml:"end_date" json:"endDate"`
}

func defaultConfig() *config {
	return &config{
		StartDate: "2024-10-01",
		EndDate:   "2024-12-01",
	}
}

func (c *config) requireStartDate() time.Time {
	t, err := time.Parse("2006-01-02", c.StartDate)
	if err != nil {
		log.Fatalf("start date %s is invalid: %s", c.StartDate, err.Error())
	}

	return t
}

func tryLoadConfigFromFile() (*config, error) {
	usr, _ := user.Current()
	currentDir, _ := os.Getwd()

	homeCfg := filepath.Join(usr.HomeDir, ".basic-roadmap.toml")
	currentCfg := filepath.Join(currentDir, ".basic-roadmap.toml")
	if _, err := os.Stat(homeCfg); err == nil {
		return readConfigFromPath(homeCfg)
	} else if _, err := os.Stat(currentCfg); err == nil {
		return readConfigFromPath(currentCfg)
	}

	return defaultConfig(), nil
}

func readConfigFromPath(path string) (*config, error) {
	cfg := defaultConfig()
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *config) RequireEndDate() time.Time {
	t, err := time.Parse("2006-01-02", c.EndDate)
	if err != nil {
		log.Fatalf("end date %s is invalid: %s", c.StartDate, err.Error())
	}

	return t
}

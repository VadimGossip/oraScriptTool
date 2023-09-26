package app

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

type App struct {
	name         string
	configDir    string
	appStartedAt time.Time
}

func NewApp(name, configDir string) *App {
	return &App{
		name:      name,
		configDir: configDir,
	}
}

func (app *App) Run() {
	dbAdapter := NewDBAdapter()
	if err := dbAdapter.Connect(); err != nil {
		logrus.Fatalf("Fail to connect git %s", err)
	}
}

package service

import "fmt"

var (
	CommitHash    = "commitHash"
	VersionNumber = "versionNumber"
	BuildDateTime = "buildDateTime"
)

type Application struct {
	ApplicationName string
	CommitHash      string
	VersionNumber   string
	BuildDateTime   string
}

func New(name string) *Application {
	return &Application{
		ApplicationName: name,
		CommitHash:      CommitHash,
		VersionNumber:   VersionNumber,
		BuildDateTime:   BuildDateTime,
	}
}

func (a *Application) BuildInfo() string {
	return fmt.Sprintf("%s %s\nBuild-Date: %s\nRef: %s", a.ApplicationName, a.VersionNumber, a.BuildDateTime, a.CommitHash)
}

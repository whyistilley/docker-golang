package service

import "fmt"

var (
	CommitHash    = "commitHash"
	VersionNumber = "versionNumber"
	BuildDateTime = "buildDateTime"
)

type Service struct {
	ServiceName   string
	CommitHash    string
	VersionNumber string
	BuildDateTime string
}

func New(name string) *Service {
	return &Service{
		ServiceName:   name,
		CommitHash:    CommitHash,
		VersionNumber: VersionNumber,
		BuildDateTime: BuildDateTime,
	}
}

func (a *Service) BuildInfo() string {
	return fmt.Sprintf("%s %s\nBuild-Date: %s\nRef: %s", a.ServiceName, a.VersionNumber, a.BuildDateTime, a.CommitHash)
}

package main

import (
	"fmt"
	"os/exec"

	"github.com/martinsirbe/services-status-notifier-bitbar/internal/models"
	"github.com/martinsirbe/services-status-notifier-bitbar/internal/service"
	log "github.com/sirupsen/logrus"
)

const goodServiceStatus = "All Systems Operational"

var services = []models.ServiceStatusPage{
	{Name: "CircleCI", ID: "6w4r0ttlx5ft"},
	{Name: "Trello", ID: "h5frqhb041yq"},
	{Name: "Atlassian", ID: "x67gp49yvrzv"},
}

func main() {
	servicesChecker := service.NewServicesChecker(services)

	services := servicesChecker.Check()

	gotBadStatus := checkForBadStatus(services)
	if gotBadStatus {
		fmt.Println("🔥\n---")
	} else {
		fmt.Println("👌\n---")
	}

	for _, s := range services {
		serviceStatus := s.Status.Description

		if serviceStatus != goodServiceStatus {
			cmd := exec.Command("say", fmt.Sprintf("%s reports %s", s.Page.Name, s.Status.Description))
			if err := cmd.Run(); err != nil {
				log.WithError(err).Error("failed execute say command")
			}
		}

		fmt.Printf("%s | href=%s color=#ffffff size=13\n", s.Page.Name, s.Page.URL)
		fmt.Printf("%s %s | color=%s size=9\n", getStatusSymbol(serviceStatus), serviceStatus, getStatusColor(serviceStatus))
		fmt.Println("---")
	}
}

func checkForBadStatus(statuses []*models.StatusPage) bool {
	for _, s := range statuses {
		if s.Status.Description != goodServiceStatus {
			return true
		}
	}
	return false
}

func getStatusSymbol(status string) string {
	if status == goodServiceStatus {
		return "👌"
	}
	return "🔥"
}

func getStatusColor(status string) string {
	if status == goodServiceStatus {
		return "#3fff00"
	}
	return "#ff2817"
}

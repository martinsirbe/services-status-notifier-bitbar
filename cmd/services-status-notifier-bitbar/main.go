package main

import (
	"fmt"
	"os/exec"

	"github.com/martinsirbe/services-status-notifier-bitbar/internal/models"
	"github.com/martinsirbe/services-status-notifier-bitbar/internal/service"
	log "github.com/sirupsen/logrus"
)

var services = []models.ServiceStatusPage{
	{Name: "CircleCI", ID: "6w4r0ttlx5ft"},
	{Name: "Trello", ID: "h5frqhb041yq"},
	{Name: "Atlassian", ID: "x67gp49yvrzv"},
}

func main() {
	servicesChecker := service.NewServicesChecker(services)

	var isBadStatus bool
	for _, r := range servicesChecker.Check() {
		if r.Status.Description != "All Systems Operational" {
			isBadStatus = true
			cmd := exec.Command("say", fmt.Sprintf("%s reports %s", r.Page.Name, r.Status.Description))
			if err := cmd.Run(); err != nil {
				log.WithError(err).Error("failed execute say command")
			}
		}
	}

	if isBadStatus {
		fmt.Println("🔥\n---")
		return
	}

	fmt.Println("👌\n---")
}

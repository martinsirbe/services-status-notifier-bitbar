package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/martinsirbe/services-status-notifier-bitbar/internal/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const statuspageURL = "https://%s.statuspage.io/api/v2/status.json"

// ServicesChecker defines services status checker functions
type ServicesChecker interface {
	Check() []*models.StatusPage
}

// Checker used to perform service status check
type Checker struct {
	services   []models.ServiceStatusPage
	httpClient http.Client
}

// NewServicesChecker used to initialise a new instance of services checker
func NewServicesChecker(services []models.ServiceStatusPage) *Checker {
	return &Checker{
		services: services,
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Check performs services status checks
func (c *Checker) Check() []*models.StatusPage {
	var statuses []*models.StatusPage
	for _, s := range c.services {
		status, err := c.checkStatus(s.ID)
		if err != nil {
			log.WithError(err).Error("service status check failed")
		}
		statuses = append(statuses, status)
	}

	return statuses
}

func (c *Checker) checkStatus(statusPageID string) (*models.StatusPage, error) {
	requestURL := fmt.Sprintf(statuspageURL, statusPageID)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create http get request, statuspageURL - %s", requestURL)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http get request execution failed")
	}

	var statusPage models.StatusPage
	if err := json.NewDecoder(res.Body).Decode(&statusPage); err != nil {
		return nil, errors.Wrap(err, "failed to decode response body to status page struct")
	}

	return &statusPage, nil
}

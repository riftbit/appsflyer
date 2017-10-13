package rawdata

import (
	"github.com/riftbit/go-appsflyer/dispatcher"
	"github.com/riftbit/go-appsflyer/model/rawdata"
)

const endpointOrganicInAppEventsReport = "organic_in_app_events_report/v5"

func GetOrganicInAppEventsReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointOrganicInAppEventsReport, client)
}

func GetEachOrganicInAppEventsReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointOrganicInAppEventsReport, client, f)
}

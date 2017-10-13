package rawdata

import (
	"github.com/riftbit/go-appsflyer/dispatcher"
	"github.com/riftbit/go-appsflyer/model/rawdata"
)

const endpointInAppEventsReport = "in_app_events_report/v5"

func GetInAppEventsReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointInAppEventsReport, client)
}

func GetEachInAppEventsReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointInAppEventsReport, client, f)
}

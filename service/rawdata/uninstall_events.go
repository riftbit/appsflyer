package rawdata

import (
	"github.com/riftbit/go-appsflyer/dispatcher"
	"github.com/riftbit/go-appsflyer/model/rawdata"
)

const endpointUninstallReport = "uninstall_events_report/v5"

func GetUninstallReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointUninstallReport, client)
}

func GetEachUninstallReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointUninstallReport, client, f)
}

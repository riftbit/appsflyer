package rawdata

import (
	"github.com/riftbit/go-appsflyer/dispatcher"
	"github.com/riftbit/go-appsflyer/model/rawdata"
)

const endpointInstallReport = "installs_report/v5"

func GetInstallReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointInstallReport, client)
}

func GetEachInstallReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointInstallReport, client, f)
}
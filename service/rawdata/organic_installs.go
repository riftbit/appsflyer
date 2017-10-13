package rawdata

import (
	"github.com/riftbit/go-appsflyer/dispatcher"
	"github.com/riftbit/go-appsflyer/model/rawdata"
)

const endpointOrganicInstallReport = "organic_installs_report/v5"

func GetOrganicInstallReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointOrganicInstallReport, client)
}

func GetEachOrganicInstallReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointOrganicInstallReport, client, f)
}

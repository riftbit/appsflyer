package rawdata

import (
	"testing"

	"github.com/riftbit/go-appsflyer/dispatcher"
	"github.com/riftbit/go-appsflyer/model/rawdata"
)

func TestGetEachInAppEventsReport(t *testing.T) {
	var (
		appID    = "jp.eure.android.pairs_tw"
		fromDate = "2016-08-23 00:00"
		toDate   = "2016-08-24 23:59"
	)
	client := dispatcher.NewClient(appID, fromDate, toDate)
	client.SetOptionalParameter(dispatcher.OptionalParameter{
		Reattr: "true",
	})
	if err := GetEachInAppEventsReport(client, func(r rawdata.Report) {
		t.Logf("Report = %v", r)
	}); err != nil {
		t.Error("Failed", err)
	}
}

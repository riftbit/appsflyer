# go-appsflyer
 
AppsFlyer API client library for Go.
  
https://support.appsflyer.com/hc/en-us/articles/207034346-Pull-APIs-Pulling-AppsFlyer-Reports-by-APIs

Original source: https://github.com/eure/appsflyer/ - but it not updates for a long time
  
## Installation

```bash
$ go get github.com/riftbit/go-appsflyer
```

## Usage
  
```go
import (
	"fmt"

	appsflyerDispatcher "github.com/riftbit/go-appsflyer/dispatcher"
	appsflyerRawdataSVC "github.com/riftbit/go-appsflyer/service/rawdata"
)

func GetInAppEventsReports() error {
	var (
		apiToken = "xxx-xxx-xxx"
		appID    = "xxx-xxx-xxx"
		fromDate = "2016-08-20"
		toDate   = "2016-08-21"
	)
	client := appsflyerDispatcher.NewClientWithParam(apiToken, appID, fromDate, toDate)

	// export APPSFLYER_API_TOKEN=""
	// client := dispatcher.NewClient(appID, fromDate, toDate)
	
	reports, err := appsflyerRawdataSVC.GetInAppEventsReports(client)
	if err != nil {
		return err
	}
	fmt.Println(reports)
}
```

### RawData

https://support.appsflyer.com/hc/en-us/articles/208387843-Raw-Data-Reports-V5-

This library supports below header names.

```go
type Report struct {
	AttributedTouchType string `json:"attributed_touch_type" csv:"Attributed Touch Type"`
	AttributedTouchTime string `json:"attributed_touch_time" csv:"Attributed Touch Time"`
	InstallTime         string `json:"install_time" csv:"Install Time"`
	EventTime           string `json:"event_time" csv:"Event Time"`
	EventName           string `json:"event_name" csv:"Event Name"`
	MediaSource         string `json:"media_source" csv:"Media Source"`
	Channel             string `json:"channel" csv:"Channel"`
	Campaign            string `json:"campaign" csv:"Campaign"`
	Ad                  string `json:"ad" csv:"Ad"`
	AdvertisingID       string `json:"advertising_id" csv:"Advertising ID"`
	IDFA                string `json:"idfa" csv:"IDFA"`
	CustomerUserID      string `json:"customer_user_id" csv:"Customer User ID"`
	IsRetargeting       string `json:"is_retargeting" csv:"Is Retargeting"`
	IP                  string `json:"ip" csv:"IP"`
	AppsflyerID         string `json:"appsflyer_id" csv:"AppsFlyer ID"`
	AndroidID           string `json:"android_id" csv:"Android ID"`
	OSVersion           string `json:"os_version" csv:"OS Version"`
	AppVersion          string `json:"app_version" csv:"App Version"`
	SDKVersion          string `json:"sdk_version" csv:"SDK Version"`
	UserAgent           string `json:"user_agent" csv:"User Agent"`
	OriginalURL         string `json:"original_url" csv:"Original URL"`
}
```

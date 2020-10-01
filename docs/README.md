# ohdear
--
    import "github.com/VictorAvelar/goh-dear/ohdear"

Package ohdear is a Golang SDK to interact with oh-dear REST api.

The Oh Dear API lets you configure everything about our application through a
simple, structured Application Programming Interface (API). Everything you see
in your dashboard can be controlled with the API. And as a bonus, all changes
you make with the API will be visible in realtime on your dashboard.

The full api documentation can be found in
https://ohdear.app/docs/general/welcome

When instantiating a new client, you can provide and API token using
OHDEAR_API_TOKEN which is the default environment variable and it will be looked
up by the library automatically. This is strongly recommended as it is the most
secure way to deal with your key.

If the library cannot resolve your API token from the environment you can
provide it when instantiating a new ohdear client using the NewClient
constructor.

## Usage

```go
const (
	BaseURL             string = "https://ohdear.app/api/"
	TokenType           string = "Bearer"
	APITokenEnv         string = "OHDEAR_API_TOKEN"
	ContentExchangeType string = "application/json"
	AuthHeader          string = "Authorization"
)
```
Oh-dear package level constants

```go
const SitesBasePath string = "/sites"
```
SitesBasePath is the resource path prefix.

```go
var (
	ErrEmptyAPIToken  error = fmt.Errorf("your api token is empty, please provide a non-empty token")
	ErrInvalidBaseURL error = fmt.Errorf("your base url must contain a trailing slash")
)
```
Oh-dear package level errors

#### func  CheckResponse

```go
func CheckResponse(r *http.Response) error
```
CheckResponse checks the API response for errors, and returns them if present. A
response is considered an error if it has a status code outside the 200 range.
API's error responses must have either no response body, or a JSON response
body.

#### type BaseClient

```go
type BaseClient interface {
	NewAPIRequest(method, uri string, body interface{}) (req *http.Request, err error)
	Do(ctx context.Context, req *http.Request) (res *Response, err error)
}
```

BaseClient interface describe an oh-dear API implementation.

#### type BrokenLinksSettingsRequest

```go
type BrokenLinksSettingsRequest struct {
	BrokenLinksCheckIncludeExternalLinks bool   `json:"broken_links_check_include_external_links,omitempty"`
	BrokenLinksWhitelistedURLS           string `json:"broken_links_whitelisted_urls,omitempty"`
}
```

BrokenLinksSettingsRequest describes the request body required to update a site
broken links settings.

Values are not required and should only be sent when updated.

#### type Client

```go
type Client struct {
	BaseURL *url.URL

	// Services
	Sites *SitesSrv
}
```

Client is the main API caller.

#### func  NewClient

```go
func NewClient(baseClient *http.Client, baseURL, apiToken string) (dear *Client, err error)
```
NewClient returns a new Oh-Dear HTTP API client. You can pass a previously build
http client, if none is provided then http.DefaultClient will be used.

NewClient will lookup the environment for values to assign to the API token
(`OHDEAR_API_TOKEN`) to be used as authentication.

#### func (*Client) Do

```go
func (c *Client) Do(req *http.Request) (*Response, error)
```
Do sends an API request and returns the API response or returned as an error if
an API error has occurred.

#### func (*Client) NewAPIRequest

```go
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error)
```
NewAPIRequest is a wrapper around the http.NewRequest function.

It will setup the authentication headers/parameters according to the client
config.

#### type CustomDate

```go
type CustomDate struct {
	time.Time
}
```


#### func (*CustomDate) UnmarshalJSON

```go
func (d *CustomDate) UnmarshalJSON(b []byte) error
```

#### type DowntimePeriods

```go
type DowntimePeriods struct {
	StartedAt *time.Time `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at"`
}
```

DowntimePeriods describes the individual values returned for site downtime
responses.

#### type DowntimeRequestFilters

```go
type DowntimeRequestFilters struct {
	StartedAt string `url:"filter[started_at]"`
	EndedAt   string `url:"filter[ended_at]"`
}
```

DowntimeRequestFilters adds the required filters to retrieve a window of
downtime values. The specified dates should be represented as follows:
20200801000000

Both values are required for downtime requests.

#### type DowntimeResponse

```go
type DowntimeResponse struct {
	Data []*DowntimePeriods `json:"data"`
}
```

DowntimeResponse is an array of values inside an outer data wrapper.

#### type Error

```go
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Response *http.Response `json:"response"` // the full response that produced the error
}
```

Error maps a standard error to a more useful data structure which is enriched
with the failing request pointer.

#### func (*Error) Error

```go
func (e *Error) Error() string
```
Error function complies with the error interface

#### type ListSitesRequestFilters

```go
type ListSitesRequestFilters struct {
	PageSize       uint   `url:"page[size],omitempty"`
	PageNumber     uint   `url:"page[number],omitempty"`
	SortBy         string `url:"sort,omitempty"`
	FilterByTeamID uint   `url:"filter[team_id],omitempty"`
}
```

ListSitesRequestFilters adds the required query string parameters to control the
response values of a list sites requests.

When filtering by a property you can prefix with a `-` sign to indicate that you
want results descending.

Non of the values are required.

#### type Response

```go
type Response struct {
	*http.Response
}
```

Response wraps the standard http.Response returned from oh-dear and provides
non-blocking access to the request content.

#### type Site

```go
type Site struct {
	ID                                   uint        `json:"id,omitempty"`
	URL                                  string      `json:"url,omitempty"`
	SortURL                              string      `json:"sort_url,omitempty"`
	Label                                string      `json:"label,omitempty"`
	TeamID                               uint        `json:"team_id,omitempty"`
	LatestRunDate                        *CustomDate `json:"latest_run_date,omitempty"`
	CreatedAt                            *CustomDate `json:"created_at,omitempty"`
	UpdatedAt                            *CustomDate `json:"updated_at,omitempty"`
	Checks                               []struct{}  `json:"checks,omitempty"`
	SummarizedChecksResult               string      `json:"summarized_checks_result,omitempty"`
	FriendlyName                         string      `json:"friendly_name,omitempty"`
	UsesHTTPS                            bool        `json:"uses_https,omitempty"`
	BrokenLinksCheckIncludeExternalLinks bool        `json:"broken_links_check_include_external_links,omitempty"`
	BrokenLinksWhitelistedURLS           []*url.URL  `json:"broken_links_whitelisted_urls,omitempty"`
}
```

Site represents a monitored website and its properties.

#### type SitesSrv

```go
type SitesSrv srv
```

SitesSrv operates over the site resource

#### func (*SitesSrv) AddToBrokenLinkWhitelist

```go
func (ss *SitesSrv) AddToBrokenLinkWhitelist(id uint, url string) (site *Site, err error)
```
AddToBrokenLinkWhitelist extends the whitelist of a given site.

See:
https://ohdear.app/docs/integrations/api/sites#adding-urls-to-the-broken-links-whitelist

#### func (*SitesSrv) Create

```go
func (ss *SitesSrv) Create(s Site) (site *Site, err error)
```
Create adds a new site to your account.

See: https://ohdear.app/docs/integrations/api/sites#add-a-site-through-the-api

#### func (*SitesSrv) Delete

```go
func (ss *SitesSrv) Delete(id uint) (err error)
```
Delete removes a site from your account.

See: https://ohdear.app/docs/integrations/api/sites#deleting-a-site

#### func (*SitesSrv) Get

```go
func (ss *SitesSrv) Get(id uint) (site *Site, err error)
```
Get retrieves a specific site by its ID.

See:
https://ohdear.app/docs/integrations/api/sites#get-a-specific-site-via-the-api

#### func (*SitesSrv) GetByURL

```go
func (ss *SitesSrv) GetByURL(url string) (site *Site, err error)
```
GetByURL returns a site by its url value.

See: https://ohdear.app/swagger#/sites/get_sites_url__siteUrl_

#### func (*SitesSrv) GetDowntimePeriods

```go
func (ss *SitesSrv) GetDowntimePeriods(id uint, filters DowntimeRequestFilters) (dr *DowntimeResponse, err error)
```
GetDowntimePeriods retrieves a collection of downtime periods.

See: https://ohdear.app/swagger#/sites/get_sites__siteId__downtime

#### func (*SitesSrv) GetUptimePercentage

```go
func (ss *SitesSrv) GetUptimePercentage(id uint, filters UptimeRequestFilters) (ur *UptimeResponse, err error)
```
GetUptimePercentage returns the uptime percentage per date.

See: https://ohdear.app/swagger#/sites/get_sites__siteId__uptime

#### func (*SitesSrv) List

```go
func (ss *SitesSrv) List(filters ListSitesRequestFilters) (sites []*Site, err error)
```
List returns all the sites in your account.

See:
https://ohdear.app/docs/integrations/api/sites#get-all-sites-in-your-account

#### func (*SitesSrv) UpdateBrokenLinksSettings

```go
func (ss *SitesSrv) UpdateBrokenLinksSettings(id uint, body BrokenLinksSettingsRequest) (site *Site, err error)
```
UpdateBrokenLinksSettings changes the configuration for broken links.

See: https://ohdear.app/docs/integrations/api/sites#broken-links-settings

#### type SplitValue

```go
type SplitValue string
```

SplitValue provides an aggregation criteria for requests.

```go
const (
	SplitByDay   SplitValue = "day"
	SplitByHour  SplitValue = "hour"
	SplitByMonth SplitValue = "month"
)
```
Supported SplitValue values.

#### type UptimePerDatetime

```go
type UptimePerDatetime struct {
	Datetime         *time.Time `json:"datetime"`
	UptimePercentage float64    `json:"uptime_percentage"`
}
```

UptimePerDatetime describes the individual values returned for site uptime
responses.

#### type UptimeRequestFilters

```go
type UptimeRequestFilters struct {
	StartedAt string     `url:"filter[started_at]"`
	EndedAt   string     `url:"filter[ended_at]"`
	Split     SplitValue `url:"split"`
}
```

UptimeRequestFilters adds the required filters to retrieve a window of uptime
values. The specified dates should be represented as follows: 20200801000000.

The three values are required for uptime requests.

#### type UptimeResponse

```go
type UptimeResponse struct {
	Data []*UptimePerDatetime `json:"data"`
}
```

UptimeResponse is an array of values per date inside an outer data wrapper.

#### type WhitelistURLRequest

```go
type WhitelistURLRequest struct {
	WhitelistURL string `json:"whitelistUrl,omitempty"`
}
```

WhitelistURLRequest generates the correct json body to add a new site to the
whitelist.

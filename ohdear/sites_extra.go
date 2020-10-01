package ohdear

import "time"

// ListSitesRequestFilters adds the required query string
// parameters to control the response values of a list sites
// requests.
//
// When filtering by a property you can prefix with a `-` sign
// to indicate that you want results descending.
//
// Non of the values are required.
type ListSitesRequestFilters struct {
	PageSize       uint   `url:"page[size],omitempty"`
	PageNumber     uint   `url:"page[number],omitempty"`
	SortBy         string `url:"sort,omitempty"`
	FilterByTeamID uint   `url:"filter[team_id],omitempty"`
}

// WhitelistURLRequest generates the correct json body
// to add a new site to the whitelist.
type WhitelistURLRequest struct {
	WhitelistURL string `json:"whitelistUrl,omitempty"`
}

// UptimeResponse is an array of values per date
// inside an outer data wrapper.
type UptimeResponse struct {
	Data []*UptimePerDatetime `json:"data"`
}

// UptimePerDatetime describes the individual values returned for
// site uptime responses.
type UptimePerDatetime struct {
	Datetime         *time.Time `json:"datetime"`
	UptimePercentage float64    `json:"uptime_percentage"`
}

// UptimeRequestFilters adds the required filters to
// retrieve a window of uptime values.
// The specified dates should be represented as follows:
// 20200801000000.
//
// The three values are required for uptime requests.
type UptimeRequestFilters struct {
	StartedAt string     `url:"filter[started_at]"`
	EndedAt   string     `url:"filter[ended_at]"`
	Split     SplitValue `url:"split"`
}

// DowntimePeriods describes the individual values returned for
// site downtime responses.
type DowntimePeriods struct {
	StartedAt *time.Time `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at"`
}

// DowntimeResponse is an array of values inside an outer data wrapper.
type DowntimeResponse struct {
	Data []*DowntimePeriods `json:"data"`
}

// DowntimeRequestFilters adds the required filters to
// retrieve a window of downtime values.
// The specified dates should be represented as follows:
// 20200801000000
//
// Both values are required for downtime requests.
type DowntimeRequestFilters struct {
	StartedAt string `url:"filter[started_at]"`
	EndedAt   string `url:"filter[ended_at]"`
}

// SplitValue provides an aggregation criteria for requests.
type SplitValue string

// Supported SplitValue values.
const (
	SplitByDay   SplitValue = "day"
	SplitByHour  SplitValue = "hour"
	SplitByMonth SplitValue = "month"
)

// BrokenLinksSettingsRequest describes the request body required
// to update a site broken links settings.
//
// Values are not required and should only be sent when updated.
type BrokenLinksSettingsRequest struct {
	BrokenLinksCheckIncludeExternalLinks bool   `json:"broken_links_check_include_external_links,omitempty"`
	BrokenLinksWhitelistedURLS           string `json:"broken_links_whitelisted_urls,omitempty"`
}

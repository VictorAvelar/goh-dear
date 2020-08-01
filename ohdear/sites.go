package ohdear

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// SitesBasePath is the resource path prefix.
const SitesBasePath string = "/sites"

// SitesSrv operates over the site resource
type SitesSrv srv

// Site represents a monitored website and its properties.
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

// List returns all the sites in your account.
//
// See: https://ohdear.app/docs/integrations/api/sites#get-all-sites-in-your-account
func (ss *SitesSrv) List(filters ListSitesRequestFilters) (sites []*Site, err error) {
	q, _ := query.Values(filters)
	req, err := ss.client.NewAPIRequest(
		http.MethodGet,
		fmt.Sprintf("%s?%s", SitesBasePath, q.Encode()),
		nil,
	)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sites); err != nil {
		return
	}

	return
}

// Create adds a new site to your account.
//
// See: https://ohdear.app/docs/integrations/api/sites#add-a-site-through-the-api
func (ss *SitesSrv) Create(s Site) (site *Site, err error) {
	req, err := ss.client.NewAPIRequest(http.MethodPost, SitesBasePath, s)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &site); err != nil {
		return
	}

	return
}

// Get retrieves a specific site by its ID.
//
// See: https://ohdear.app/docs/integrations/api/sites#get-a-specific-site-via-the-api
func (ss *SitesSrv) Get(id uint) (site *Site, err error) {
	req, err := ss.client.NewAPIRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%d", SitesBasePath, id),
		nil,
	)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &site); err != nil {
		log.Println(err)
		return
	}

	return
}

// Delete removes a site from your account.
//
// See: https://ohdear.app/docs/integrations/api/sites#deleting-a-site
func (ss *SitesSrv) Delete(id uint) (err error) {
	req, err := ss.client.NewAPIRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/%d", SitesBasePath, id),
		nil,
	)
	if err != nil {
		return
	}

	_, err = ss.client.Do(req)
	if err != nil {
		return
	}

	return
}

// GetByURL returns a site by its url value.
//
// See: https://ohdear.app/swagger#/sites/get_sites_url__siteUrl_
func (ss *SitesSrv) GetByURL(url string) (site *Site, err error) {
	req, err := ss.client.NewAPIRequest(
		http.MethodGet,
		fmt.Sprintf("%s/url/%s", SitesBasePath, url),
		nil,
	)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &site); err != nil {
		return
	}

	return
}

// GetDowntimePeriods retrieves a collection of downtime periods.
//
// See: https://ohdear.app/swagger#/sites/get_sites__siteId__downtime
func (ss *SitesSrv) GetDowntimePeriods(id uint, filters DowntimeRequestFilters) (dr *DowntimeResponse, err error) {
	q, _ := query.Values(filters)

	req, err := ss.client.NewAPIRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%d/downtime?%s", SitesBasePath, id, q.Encode()),
		nil,
	)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &dr); err != nil {
		return
	}

	return
}

// GetUptimePercentage returns the uptime percentage per date.
//
// See: https://ohdear.app/swagger#/sites/get_sites__siteId__uptime
func (ss *SitesSrv) GetUptimePercentage(id uint, filters UptimeRequestFilters) (ur *UptimeResponse, err error) {
	q, _ := query.Values(filters)

	req, err := ss.client.NewAPIRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%d/uptime?%s", SitesBasePath, id, q.Encode()),
		nil,
	)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ur); err != nil {
		return
	}

	return
}

// AddToBrokenLinkWhitelist extends the whitelist of a given site.
//
// See: https://ohdear.app/docs/integrations/api/sites#adding-urls-to-the-broken-links-whitelist
func (ss *SitesSrv) AddToBrokenLinkWhitelist(id uint, url string) (site *Site, err error) {
	body := WhitelistURLRequest{
		WhitelistURL: url,
	}

	req, err := ss.client.NewAPIRequest(http.MethodPost, fmt.Sprintf("%s/%d/add-to-broken-links-whitelist", SitesBasePath, id), body)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &site); err != nil {
		return
	}

	return
}

// UpdateBrokenLinksSettings changes the configuration for broken links.
//
// See: https://ohdear.app/docs/integrations/api/sites#broken-links-settings
func (ss *SitesSrv) UpdateBrokenLinksSettings(id uint, body BrokenLinksSettingsRequest) (site *Site, err error) {
	req, err := ss.client.NewAPIRequest(
		http.MethodPut,
		fmt.Sprintf("%s/%d/update-broken-links-settings", SitesBasePath, id),
		body,
	)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &site); err != nil {
		return
	}

	return
}

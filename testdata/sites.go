package testdata

const MultipleSitesResponse = `{
  "data": [
    {
      "id": 1,
      "url": "http://yoursite.tld",
      "sort_url": "yoursite.tld",
      "label": "your-site",
      "team_id": 1,
      "latest_run_date": "2019-09-16 07:29:02",
      "summarized_check_result": "succeeded",
      "created_at": "20171106 07:40:49",
      "updated_at": "20171106 07:40:49",
      "checks": [
        {
          "id": 100,
          "type": "uptime",
          "label": "Uptime",
          "enabled": true,
          "latest_run_ended_at": "2019-09-16 07:29:02",
          "latest_run_result": "succeeded"
        },
        {
          "id": 101,
          "type": "broken_links",
          "label": "Broken links",
          "enabled": true,
          "latest_run_ended_at": "2019-09-16 07:29:05",
          "latest_run_result": "succeeded"
        },
      ]
    },
    {
      "id": 2,
      "url": "https://yourothersite.tld",
      "sort_url": "yourothersite.tld",
      "label": "my-site",
      "team_id": 1,
      "latest_run_date": "2019-09-16 07:29:02",
      "summarized_check_result": "failed",
      "created_at": "20171108 07:40:16",
      "updated_at": "20171108 07:40:16",
      "checks": [
        {
          "id": 1,
          "type": "uptime",
          "label": "Uptime",
          "enabled": true,
          "latest_run_ended_at": "2019-09-16 07:29:02",
          "latest_run_result": "succeeded"
        },
        {
          "id": 2,
          "type": "broken_links",
          "label": "Broken links",
          "enabled": true,
          "latest_run_ended_at": "2019-09-16 07:29:05",
          "latest_run_result": "failed"
        },
        {
          "id": 3,
          "type": "mixed_content",
          "label": "Mixed content",
          "enabled": true,
          "latest_run_ended_at": "2019-09-16 07:29:05",
          "latest_run_result": "succeeded"
        },
        {
          "id": 4,
          "type": "certificate_health",
          "label": "Certificate health",
          "enabled": true,
          "latest_run_ended_at": "2019-09-16 07:29:02",
          "latest_run_result": "failed"
        },
        {
          "id": 5,
          "type": "certificate_transparency",
          "label": "Certificate transparency",
          "enabled": true,
          "latest_run_ended_at": null,
          "latest_run_result": null
        }
      ]
    }
  ]
}`

const SingleSiteResponse = `{
  "id": 1,
  "url": "http://yoursite.tld",
  "sort_url": "yoursite.tld",
  "label": "your-site",
  "team_id": 1,
  "latest_run_date": "2019-09-16 07:29:02",
  "summarized_check_result": "succeeded",
  "created_at": "20171106 07:40:49",
  "updated_at": "20171106 07:40:49",
  "checks": [
	{
	  "id": 100,
	  "type": "uptime",
	  "label": "Uptime",
	  "enabled": true,
	  "latest_run_ended_at": "2019-09-16 07:29:02",
	  "latest_run_result": "succeeded"
	},
	{
	  "id": 101,
	  "type": "broken_links",
	  "label": "Broken links",
	  "enabled": true,
	  "latest_run_ended_at": "2019-09-16 07:29:05",
	  "latest_run_result": "succeeded"
	},
  ]
}`

const UptimeResponse = `{
data	[
	{
		datetime: 2018-09-22 12:00:00,
		uptime_percentage:	99.98
	},
	{
		datetime: 2018-09-23 12:00:00,
		uptime_percentage:	98.00
	}
]
}`

package main

import (
	"os"

	"github.com/cloudflare/cloudflare-go"
)

type cf struct {
	api *cloudflare.API
}

func (cf *cf) New() (*cf, error) {
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		return nil, err
	}
	cf.api = api
	return cf, nil
}

func (cf *cf) updateAllRecords(ip string) error {
	// Fetch list of zones for given domain
	zones, err := cf.api.ListZones(os.Getenv("CF_ZONE_DOMAIN"))
	if err != nil {
		return err
	}

	for _, zone := range zones {
		// List zone records
		records, err := cf.api.DNSRecords(zone.ID, cloudflare.DNSRecord{})
		if err != nil {
			return err
		}
		for _, record := range records {
			updatedRecord := cloudflare.DNSRecord{
				Type:    "A",
				Name:    record.Name,
				Content: ip,
				Proxied: true,
			}
			err := cf.api.UpdateDNSRecord(zone.ID, record.ID, updatedRecord)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

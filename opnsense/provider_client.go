package opnsense

import (
	opn_api "github.com/eugenmayer/opnsense-cli/opnsense/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"net/url"
)

type ProviderClient struct {
	Url         *url.URL
	ApiKey      string
	ApiSecret   string
	NoSslVerify bool
	Conn        *opn_api.OPNsense
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	opnUrl := d.Get(OPN_URL).(string)
	if opnUrl == "" {
		log.Println("Defaulting environment in URL config to use API default version...")
	}

	opnApiKey := d.Get(OPN_APIKEY).(string)
	if opnApiKey == "" {
		log.Println("Defaulting environment in URL config to use API default hostname...")
		opnApiKey = "localhost"
	}

	opnApiSecret := d.Get(OPN_APISECRET).(string)
	if opnApiSecret == "" {
		log.Println("Defaulting environment in URL config to use API default hostname...")
		opnApiSecret = "localhost"
	}

	opnNoSslVerify := d.Get(OPN_NOSSLVERIFY).(bool)

	return newProviderClient(opnUrl, opnApiKey, opnApiSecret, opnNoSslVerify)
}

// newProviderClient is a factory for creating ProviderClient structs
func newProviderClient(address string, apiKey string, apiSecret string, noSslVerify bool) (*ProviderClient, error) {
	u, err := url.Parse(address)
	if err == nil {
		p := &ProviderClient{
			Url:         u,
			ApiKey:      apiKey,
			ApiSecret:   apiSecret,
			NoSslVerify: noSslVerify,
		}
		p.Conn = &opn_api.OPNsense{
			BaseUrl:     *u,
			ApiKey:      apiKey,
			ApiSecret:   apiSecret,
			NoSslVerify: noSslVerify,
		}
		return p, nil
	}
	return nil, err
}

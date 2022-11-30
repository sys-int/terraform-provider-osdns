package opnsense

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const OPN_URL = "opn_url"
const OPN_NOSSLVERIFY = "opn_nosslverify"
const OPN_APIKEY = "opn_apikey"
const OPN_APISECRET = "opn_apisecret"

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,

		Schema: map[string]*schema.Schema{
			OPN_URL: {
				Type:     schema.TypeString,
				Required: true,
			},

			OPN_APIKEY: {
				Type:     schema.TypeString,
				Required: true,
			},

			OPN_APISECRET: {
				Type:     schema.TypeString,
				Required: true,
			},

			OPN_NOSSLVERIFY: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"osdns_override": readHostOverride(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"osdns_override": hostOverride(),
		},
	}
}

func hostOverride() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHostOverrideCreate,
		ReadContext:   resourceHostOverrideRead,
		UpdateContext: resourceHostOverrideUpdate,
		DeleteContext: resourceHostOverrideDelete,
		Schema:        hostOverrideSchema(),
	}
}

func readHostOverride() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceHostOverrideRead,
		Schema:      hostOverrideSchema(),
	}
}

func hostOverrideSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"host": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ip": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rr": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mx": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mxprio": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

package opnsense

import (
	opn_unbound "github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func CreateOverrideHost(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	host := unmarshalHost(d)
	uuid, err := api.HostOverrideCreate(host)
	if err == nil {
		host.Uuid = uuid
		marshalHost(d, host)
	}
	return err
}

func ReadOverrideHost(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	host, err := api.HostEntryGetByFQDN(d.Get("host").(string), d.Get("domain").(string))
	if err == nil {
		marshalHost(d, host)
	}
	return nil
}

func marshalHost(d *schema.ResourceData, host opn_unbound.HostOverride) {
	d.Set("enabled", host.Enabled)
	d.Set("host", host.Host)
	d.Set("domain", host.Domain)
	d.Set("ip", host.Ip)
	d.Set("rr", host.Rr)
	d.Set("mxprio", host.Mxprio)
	d.Set("mx", host.Mx)
	d.Set("description", host.Description)
	d.Set("uuid", host.Uuid)
	d.SetId(host.Uuid)
}

func unmarshalHost(d *schema.ResourceData) opn_unbound.HostOverride {
	host := opn_unbound.HostOverride{
		Enabled:     d.Get("enabled").(string),
		Host:        d.Get("host").(string),
		Domain:      d.Get("domain").(string),
		Ip:          d.Get("ip").(string),
		Rr:          d.Get("rr").(string),
		Mxprio:      d.Get("mxprio").(string),
		Mx:          d.Get("mx").(string),
		Description: d.Get("description").(string),
		Uuid:        d.Get("uuid").(string),
	}
	return host
}

func DeleteOverrideHost(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	err = api.HostEntryRemove(d.Get("uuid").(string))
	return err
}

func UpdateOverrideHost(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	host := unmarshalHost(d)
	uuid, err := api.HostOverrideUpdate(host)
	host.Uuid = uuid
	marshalHost(d, host)
	return err
}

func ExistsOverrideHost(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(ProviderClient)
	api := opn_unbound.UnboundApi{client.Conn}
	host, err := api.HostEntryGetByFQDN(d.Get("host").(string), d.Get("domain").(string))
	if err != nil {
		return false, nil
	}
	if &host != nil {
		return true, nil
	}
	return false, nil
}

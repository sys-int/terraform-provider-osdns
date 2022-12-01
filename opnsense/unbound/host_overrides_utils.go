package unbound

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sys-int/opnsense-api/api/unbound"
)

func marshalHost(ctx context.Context, d *schema.ResourceData, host unbound.HostOverride) {
	enabled := true
	if host.Enabled != "" {
		if host.Enabled == "1" {
			enabled = true
		} else {
			enabled = false
		}
	}
	d.Set("enabled", enabled)
	d.Set("host", host.Host)
	d.Set("domain", host.Domain)
	d.Set("ip", host.Ip)
	d.Set("mxprio", host.Mxprio)
	d.Set("mx", host.Mx)
	d.Set("description", host.Description)
}

func unmarshalHost(ctx context.Context, d *schema.ResourceData, host *unbound.HostOverride) *unbound.HostOverride {
	enabled := "1"
	if d.Get("enabled").(bool) {
		enabled = "1"
	} else {
		enabled = "0"
	}
	host.Enabled = enabled
	host.Host = d.Get("host").(string)
	host.Domain = d.Get("domain").(string)
	host.Ip = d.Get("ip").(string)
	host.Mxprio = d.Get("mxprio").(string)
	host.Mx = d.Get("mx").(string)
	host.Description = d.Get("description").(string)
	host.Uuid = d.Id()

	tflog.Debug(ctx, "post-unmarshalHost", map[string]interface{}{
		"enabled":     host.Enabled,
		"uuid":        host.Uuid,
		"host":        host.Host,
		"domain":      host.Domain,
		"ip":          host.Ip,
		"mx":          host.Mx,
		"mxprio":      host.Mxprio,
		"description": host.Description,
	})
	return host
}

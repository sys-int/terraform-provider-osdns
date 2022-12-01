package opnsense

import "github.com/sys-int/opnsense-api/api"

type IProviderClient interface {
	GetConn() *api.OPNsense
}

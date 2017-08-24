package driver

import (
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/mo"
)

type Host struct {
	driver *Driver
	host *object.HostSystem
}

func (d *Driver) NewHost(ref *types.ManagedObjectReference) *Host {
	return &Host{
		host: object.NewHostSystem(d.client.Client, *ref),
		driver: d,
	}
}

func (h *Host) Info(params ...string) (*mo.HostSystem, error){
	var p []string
	if len(params) == 0 {
		p = []string{"*"}
	} else {
		p = params
	}
	var info mo.HostSystem
	err := h.host.Properties(h.driver.ctx, h.host.Reference(), p, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

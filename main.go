package main

import (
	"git.openstack.org/openstack/docker-machine-openstack/driver"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(openstack.NewDriver("default", "path"))
}

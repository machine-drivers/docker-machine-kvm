package main

import (
	"github.com/code-ready/machine-driver-libvirt"
	"github.com/code-ready/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(libvirt.NewDriver("default", "path"))
}

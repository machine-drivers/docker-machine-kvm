package libvirt

const (
	DriverName         = "libvirt"
	DriverVersion      = "0.9.1"

	connectionString   = "qemu:///system"
	dnsmasqLeases      = "/var/lib/libvirt/dnsmasq/%s.leases"
	dnsmasqStatus      = "/var/lib/libvirt/dnsmasq/%s.status"
)
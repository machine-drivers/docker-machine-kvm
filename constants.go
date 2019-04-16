package libvirt

const (
	DriverName         = "libvirt"
	DriverVersion      = "0.9.0"

	connectionString   = "qemu:///system"
	privateNetworkName = "crc"
	dnsmasqLeases      = "/var/lib/libvirt/dnsmasq/%s.leases"
	dnsmasqStatus      = "/var/lib/libvirt/dnsmasq/%s.status"
	defaultSSHUser     = "core"

	domainXMLTemplate = `<domain type='kvm'>
  <name>{{.MachineName}}</name>
  <memory unit='MB'>{{.Memory}}</memory>
  <vcpu>{{.CPU}}</vcpu>
  <features><acpi/><apic/><pae/></features>
  <cpu mode='host-passthrough'></cpu>
  <os>
    <type>hvm</type>
    <boot dev='hd'/>
    <bootmenu enable='no'/>
  </os>
  <devices>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='{{.CacheMode}}' io='{{.IOMode}}' />
      <source file='{{.DiskPath}}'/>
      <target dev='hda' bus='ide'/>
    </disk>
    <graphics type='vnc' autoport='yes' listen='127.0.0.1'>
      <listen type='address' address='127.0.0.1'/>
    </graphics>
    <interface type='network'>
      <source network='{{.Network}}'/>
    </interface>
    <interface type='network'>
      <source network='{{.PrivateNetwork}}'/>
    </interface>
  </devices>
</domain>`
	networkXML = `<network>
  <name>%s</name>
  <ip address='%s' netmask='%s'>
    <dhcp>
      <range start='%s' end='%s'/>
    </dhcp>
  </ip>
</network>`
)
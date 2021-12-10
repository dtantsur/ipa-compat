package ipa

type BootMode string

const (
	UEFI   BootMode = "uefi"
	Legacy BootMode = "bios"
)

type BootInfo struct {
	// Boot mode the machine is currently in.
	CurrentBootMode BootMode `json:"current_boot_mode"`

	// MAC address of the PXE booting interface (if known).
	PXEInterface string `json:"pxe_interface,omitempty"`

	// The current host name.
	HostName string `json:"hostname"`
}

type BlockDevice struct {
	// Block device name (path).
	Name string `json:"name"`

	// Size in bytes.
	Size int64 `json:"size"`

	// Whether the device is a rotational disk.
	IsRotational bool `json:"rotational"`

	// Model (optional).
	Model string `json:"model,omitempty"`

	// Serial number (optional).
	Serial string `json:"serial,omitempty"`

	// WWN (optional).
	WWN string `json:"wwn,omitempty"`
}

type TLV [2]string

type NetworkInterface struct {
	// Interface name.
	Name string `json:"name"`

	// MAC address.
	MACAddress string `json:"mac_address"`

	// The current IPv4 address (if any).
	IPv4Address string `json:"ipv4_address,omitempty"`

	// The current IPv6 address (if any).
	IPv6Address string `json:"ipv6_address,omitempty"`

	// Raw LLDP information (optional).
	LLDP []TLV `json:"lldp,omitempty"`
}

type Memory struct {
	// Total memory in bytes as seen by the OS.
	Total int64 `json:"total"`

	// Physical memory in MiB as reported by DMI.
	PhysicalMiB int64 `json:"physical_mb"`
}

type Processors struct {
	// Core count.
	Count int `json:"count"`

	// Architecture.
	Architecture string `json:"architecture"`

	// Frequence in MHz.
	FrequencyMHz int `json:"frequency"`

	// Model name (optional).
	ModelName string `json:"model_name"`

	// CPU flags.
	Flags []string `json:"flags"`
}

type SystemVendor struct {
	ProductName  string `json:"product_name"`
	SerialNumber string `json:"serial_number"`
	Manufacturer string `json:"manufacturer"`
}

// Shorter version of the IPA inventory defined in
// https://docs.openstack.org/ironic-python-agent/latest/admin/how_it_works.html#hardware-inventory
type Inventory struct {
	// Information about the current boot.
	BootInfo BootInfo `json:"boot"`

	// List of block devices.
	BlockDevices []BlockDevice `json:"disks"`

	// List of network interfaces.
	NetworkInterfaces []NetworkInterface `json:"interfaces"`

	// Memory information
	Memory Memory `json:"memory"`

	// Processors information.
	Processors Processors `json:"cpu"`

	// System vendor information.
	SystemVendor SystemVendor `json:"system_vendor"`

	// BMC address (e.g. via ipmitool) - v4.
	BMCAddressV4 string `json:"bmc_address,omitempty"`

	// BMC address (e.g. via ipmitool) - v6.
	BMCAddressV6 string `json:"bmc_v6address,omitempty"`
}

type ConfigDrive struct {
	// User data (typically an Ignition script).
	UserData string
}

// Short version of an Ironic node.
type Node struct {
	// Unique ID.
	UUID string `json:"uuid"`

	// Unique name (optional).
	Name string `json:"name,omitempty"`

	// Instance information. Use GetConfigDrive to extract configuration
	// drive information.
	InstanceInfo map[string]interface{} `json:"instance_info"`
}

// Short version of an Ironic port.
type Port struct {
	// Unique ID.
	UUID string `json:"uuid"`

	// MAC address
	MACAddress string `json:"address"`
}

// An executable step.
type Step struct {
	// Interface name (use "deploy" if unsure").
	Interface string `json:"interface"`

	// Step name.
	Name string `json:"step"`

	// Priority. 0 to disable by default.
	Priority int `json:"priority"`

	// Function to execute.
	Execute func(Node, []Port) error `json:"-"`
}

type IronicAgent interface {
	GetInventory() (Inventory, error)

	ListDeploySteps() ([]Step, error)

	GetDeployStep(interfaceName, stepName string) (Step, error)

	ListCleanSteps() ([]Step, error)

	GetCleanStep(interfaceName, stepName string) (Step, error)
}

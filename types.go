package sonic

import "strings"

type (
	Table string
	Key   struct {
		keys []string
	}
)

func AsKey(k ...string) Key {
	return Key{
		keys: k,
	}
}
func (k *Key) String() string {
	return strings.Join(k.keys, keySeparator)
}

const (
	keySeparator       = "|"
	tableNameSeparator = "|"
)

const (
	AutoTechsupportFeatureTable Table = "AUTO_TECHSUPPORT_FEATURE"
	BreakoutConfigTable         Table = "BREAKOUT_CFG"
	ConfigDBInitializedTable    Table = "CONFIG_DB_INITIALIZED"
	DeviceMetadataTable         Table = "DEVICE_METADATA"
	FeatureTable                Table = "FEATURE"
	FlexCounterTable            Table = "FLEX_COUNTER_TABLE"
	InterfaceTable              Table = "INTERFACE"
	LoopbackInterfaceTable      Table = "LOOPBACK_INTERFACE"
	LLDPTable                   Table = "LLDP"
	MgmtVRFConfigTable          Table = "MGMT_VRF_CONFIG"
	NTPTable                    Table = "NTP"
	NTPServerTable              Table = "NTP_SERVER"
	PortTable                   Table = "PORT"
	VlanTable                   Table = "VLAN"
	VlanInterfaceTable          Table = "VLAN_INTERFACE"
	VlanMemberTable             Table = "VLAN_MEMBER"
	VrfTable                    Table = "VRF"
	VersionsTable               Table = "VERSIONS"
	XcvrdLogTable               Table = "XCVRD_LOG"
	VxlanTunnelMapTable         Table = "VXLAN_TUNNEL_MAP"
)

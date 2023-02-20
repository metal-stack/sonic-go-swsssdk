package sonic

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redismock/v9"
)

func TestVlanMember_SetMember(t *testing.T) {
	table := VlanMemberTable
	ctx := context.Background()

	db, mock := redismock.NewClientMock()
	c := &configDB{
		client: db,
	}

	tests := []struct {
		name          string
		interfaceName string
		vlan          uint16
		tagged        bool
		wantErr       bool
	}{
		{
			name:          "set vlan membership untagged",
			interfaceName: "Ethernet0",
			vlan:          4001,
			tagged:        false,
			wantErr:       false,
		},
		{
			name:          "set vlan membership tagged",
			interfaceName: "Ethernet0",
			vlan:          4001,
			tagged:        true,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := fmt.Sprintf("%s|Vlan%d|%s", VlanMemberTable, tt.vlan, tt.interfaceName)
			mock.ExpectHGetAll(key).SetVal(nil)
			if !tt.tagged {
				mock.ExpectHSet(key, map[string]any{"tagging_mode": "untagged"}).SetVal(1)
			} else {
				mock.ExpectHSet(key).SetVal(1)
			}
			v := &VlanMember{
				db: c,
				t:  table,
			}
			if err := v.SetMember(ctx, tt.interfaceName, tt.vlan, tt.tagged); (err != nil) != tt.wantErr {
				t.Errorf("VlanMember.SetMember() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

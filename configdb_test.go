package sonic

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redismock/v9"
)

func Test_configDB_Get(t *testing.T) {
	tests := []struct {
		name    string
		table   Table
		key     Key
		has     map[string]string
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "get Interface MTU",
			table:   InterfaceTable,
			key:     AsKey("Ethernet0"),
			has:     map[string]string{"mtu": "9000"},
			want:    map[string]string{"mtu": "9000"},
			wantErr: false,
		},
		{
			name:    "Get VXlan Tunnel Map",
			table:   InterfaceTable,
			key:     AsKey("vtep", "map_104000_Vlan4000"),
			has:     map[string]string{"vlan": "Vlan4000", "vni": "104000"},
			want:    map[string]string{"vlan": "Vlan4000", "vni": "104000"},
			wantErr: false,
		},
	}
	db, mock := redismock.NewClientMock()

	c := &configDB{
		client: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectHGetAll(k(tt.table, tt.key)).SetVal(tt.has)

			got, err := c.Get(context.Background(), tt.table, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("configDB.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDB.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

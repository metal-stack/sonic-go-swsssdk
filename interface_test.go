package sonic

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redismock/v9"
)

func TestInterface_MTU(t *testing.T) {
	table := InterfaceTable
	ctx := context.Background()

	db, mock := redismock.NewClientMock()
	c := &configDB{
		client: db,
	}

	tests := []struct {
		name          string
		interfaceName string
		vrf           uint16
		mock          redismock.ClientMock
		wantErr       bool
	}{
		{
			name:          "Set MTU",
			interfaceName: "Ethernet0",
			vrf:           30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectHSet(fmt.Sprintf("%s|%s", InterfaceTable, tt.interfaceName), map[string]string{"vrf_name": fmt.Sprintf("Vrf%d", tt.vrf)}).SetVal(int64(1))
			i := &Interface{
				db:         c,
				t:          table,
				interfaces: []string{tt.interfaceName},
			}
			if err := i.Vrf(ctx, tt.vrf); (err != nil) != tt.wantErr {
				t.Errorf("Interface.MTU() error = %q, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

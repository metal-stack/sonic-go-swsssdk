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
		mtu           uint16
		mock          redismock.ClientMock
		wantErr       bool
	}{
		{
			name:          "Set MTU",
			interfaceName: "Ethernet0",
			mtu:           9000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectHSet(fmt.Sprintf("%s|%s", InterfaceTable, tt.interfaceName), map[string]string{"mtu": fmt.Sprintf("%d", tt.mtu)}).SetVal(int64(1))
			i := &Interface{
				db: c,
				t:  table,
			}
			if err := i.MTU(ctx, tt.interfaceName, tt.mtu); (err != nil) != tt.wantErr {
				t.Errorf("Interface.MTU() error = %q, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

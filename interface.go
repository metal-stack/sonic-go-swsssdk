package sonic

import (
	"context"
	"fmt"
)

type Interface struct {
	db ConfigDB
	t  Table
}

func (i *Interface) MTU(ctx context.Context, interfaceName string, mtu uint16) error {
	return i.db.Set(ctx, i.t, AsKey(interfaceName), map[string]any{"mtu": fmt.Sprintf("%d", mtu)})
}

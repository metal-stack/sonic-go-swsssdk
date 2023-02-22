package sonic

import (
	"context"
	"fmt"
)

type Port struct {
	db   ConfigDB
	t    Table
	port string
}

func (p *Port) MTU(ctx context.Context, mtu uint16) error {
	existing, err := p.db.Get(ctx, p.t, AsKey(p.port))
	if err != nil {
		return err
	}

	existing["mtu"] = fmt.Sprintf("%d", mtu)

	return p.db.Set(ctx, p.t, AsKey(p.port), existing)
}

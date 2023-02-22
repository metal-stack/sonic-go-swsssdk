package sonic

import (
	"context"
	"fmt"

	"github.com/metal-stack/sonic-go-swsssdk/api"
	"github.com/openconfig/ygot/ygot"
)

type Port struct {
	db   ConfigDB
	t    Table
	port string
}

func (p *Port) MTU(ctx context.Context, port api.SonicPort_SonicPort_PORT_PORT_LIST) error {
	existing, err := p.db.Get(ctx, p.t, AsKey(p.port))
	if err != nil {
		return err
	}

	old := &api.SonicPort_SonicPort_PORT_PORT_LIST{}
	if err := api.Unmarshal([]byte(existing), old); err != nil {
		return fmt.Errorf("Cannot unmarshal JSON: %v", err)
	}

	newPort := old
	newPort.Mtu = port.Mtu

	js, err := ygot.EmitJSON(newPort, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		return err
	}

	return p.db.Set(ctx, p.t, AsKey(p.port), js)
}

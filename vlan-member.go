package sonic

import (
	"context"
	"fmt"
)

const (
	taggingMode = "tagging_mode"
	untagged    = "untagged"
)

type VlanMember struct {
	db ConfigDB
	t  Table
}

func (v *VlanMember) SetMember(ctx context.Context, interfaceName string, vlan uint16, tagged bool) error {
	key := AsKey(fmt.Sprintf("Vlan%d", vlan), interfaceName)

	existing, err := v.db.Get(ctx, v.t, key)
	if err != nil {
		return err
	}

	if len(existing) == 1 && existing[taggingMode] == untagged {
		return nil
	}

	var value map[string]string
	if !tagged {
		value = map[string]string{
			taggingMode: untagged,
		}
	} else {
		value = nil
	}

	return v.db.Set(ctx, v.t, key, value)
}

func (v *VlanMember) DelMember(ctx context.Context, interfaceName, vlan string) error {
	key := AsKey(fmt.Sprintf("Vlan%s", vlan), interfaceName)
	existing, err := v.db.Get(ctx, v.t, key)
	if err != nil {
		return err
	}
	if len(existing) == 0 {
		return nil
	}
	return v.db.Delete(ctx, v.t, key)
}

package sonic

import (
	"context"
	"errors"
	"fmt"
)

type Interface struct {
	db         ConfigDB
	t          Table
	interfaces []string
}

func (i *Interface) Vrf(ctx context.Context, vrf uint16) error {
	if err := i.Validate(); err != nil {
		return err
	}
	var errs []error
	for _, ifaceName := range i.interfaces {
		err := i.db.Set(ctx, i.t, AsKey(ifaceName), map[string]any{"vrf_name": fmt.Sprintf("Vrf%d", vrf)})
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (i *Interface) Validate() error {
	if len(i.interfaces) == 0 {
		return fmt.Errorf("at least one interface must be given")
	}
	return nil
}

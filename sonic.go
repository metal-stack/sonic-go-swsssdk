package sonic

import (
	"github.com/redis/go-redis/v9"
)

type Sonic interface {
	ConfigDB() ConfigDB
	Interface() *Interface
	Vlan() *VlanMember
}

type sonic struct {
	client *redis.Client
	db     ConfigDB
}

func New(client *redis.Client) (Sonic, error) {
	db, err := NewConfigDB(client)
	if err != nil {
		return nil, err
	}
	return &sonic{
		client: client,
		db:     db,
	}, nil
}

// ConfigDB implements Sonic
func (s *sonic) ConfigDB() ConfigDB {
	return s.db
}

// Interface implements Sonic
func (s *sonic) Interface() *Interface {
	return &Interface{
		db: s.db,
		t:  PortTable,
	}
}
func (s *sonic) Vlan() *VlanMember {
	return &VlanMember{
		db: s.db,
		t:  VlanTable,
	}
}

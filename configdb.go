package sonic

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// ConfigDB presents the low level interface to the redis CONFIG_DB of a SONiC switch.
type ConfigDB interface {
	Set(ctx context.Context, table Table, key Key, data map[string]string) error
	Get(ctx context.Context, table Table, key Key) (map[string]string, error)
	Delete(ctx context.Context, table Table, key Key) error
	Update(ctx context.Context, table Table, key Key, data map[string]string) error
	GetKeys(ctx context.Context, table Table, split bool) ([]Key, error)
}

type configDB struct {
	client *redis.Client
}

func NewConfigDB(client *redis.Client) (ConfigDB, error) {
	initialized := client.Get(context.Background(), string(ConfigDBInitializedTable))
	if initialized == nil {
		return nil, fmt.Errorf("redis database is not ready yet")
	}
	return &configDB{
		client: client,
	}, nil
}

// Get implements ConfigDB
func (c *configDB) Get(ctx context.Context, table Table, key Key) (map[string]string, error) {
	result, err := c.client.HGetAll(ctx, k(table, key)).Result()
	return result, err
}

// Set implements ConfigDB
func (c *configDB) Set(ctx context.Context, table Table, key Key, data map[string]string) error {
	return c.client.HSet(ctx, k(table, key), data).Err()
}

// Update implements ConfigDB
func (c *configDB) Update(ctx context.Context, table Table, key Key, data map[string]string) error {
	return c.client.HSet(ctx, k(table, key), data).Err()
}

// Delete implements ConfigDB
func (c *configDB) Delete(ctx context.Context, table Table, key Key) error {
	return c.client.Del(ctx, k(table, key)).Err()
}

// GetKeys implements ConfigDB
func (c *configDB) GetKeys(ctx context.Context, table Table, split bool) ([]Key, error) {
	result, err := c.client.Keys(ctx, fmt.Sprintf("%s*", table)).Result()
	if err != nil {
		return nil, err
	}

	var keys []Key
	for _, r := range result {
		keys = append(keys, AsKey(r))
	}
	return keys, nil
}

func k(table Table, key Key) string {
	return string(table) + tableNameSeparator + key.String()
}

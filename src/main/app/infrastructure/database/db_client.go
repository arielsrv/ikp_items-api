package database

import (
	"sync"

	"ikp_items-api/src/main/app/ent"

	_ "github.com/go-sql-driver/mysql"
)

type IDbClient interface {
	Context() *ent.Client
	Close() error
}

type DBClient struct {
	dbClient IDbClient
	mutex    sync.Once
	*ent.Client
}

func NewDBClient(client IDbClient) *DBClient {
	return &DBClient{
		dbClient: client,
	}
}

// Context template method, used by concrete impl.
func (d *DBClient) Context() *ent.Client {
	d.mutex.Do(func() {
		d.Client = d.dbClient.Context()
	})

	return d.Client
}

func (d *DBClient) Close() error {
	return d.Client.Close()
}

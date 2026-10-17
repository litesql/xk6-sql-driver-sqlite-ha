// Package sqliteha contains SQLite HA SQL driver registration for xk6-sql.
package sqliteha

import (
	"github.com/grafana/xk6-sql/sql"

	// Blank import required for initialization of driver.
	_ "github.com/litesql/go-remote-ha"
)

func init() {
	sql.RegisterModule("remote-ha")
}

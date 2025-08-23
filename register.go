// Package snowflake contains snowflake driver registration for xk6-sql.
package snowflake

import (
	"github.com/grafana/xk6-sql/sql"

	// Blank import required for initialization of driver.
	_ "github.com/snowflakedb/gosnowflake"
)

func init() {
	sql.RegisterModule("snowflake")
}

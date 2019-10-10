package postgres

// 只是举个例子，实际不能执行
import (
	"database/sql"
)

func init() {
	sql.Register("postgres", new(PostgresDriver))
}

package main

// 只是举个例子，实际不能执行
// 3-6 导入时使用空白标识符作为包的别名
import (
	"database/sql"

	_ "/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}

// demo1
package main

func main() {
	// 5-1 声明一个结构类型 user 在程序里定义一个用户类型
	type user struct {
		name       string
		email      string
		ext        int
		privileged bool
	}

	// 5-2 使用结构类型声明变量，并初始化为其零值声明 user 类型的变量
	var bill user

	// 5-3 使用结构字面量来声明一个结构类型的变量声明 user 类型的变量，并初始化所有字段
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// 5-4 使用结构字面量创建结构类型的值（5-3）里部分
	user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// 5-5 不使用字段名，创建结构类型的值声明 user 类型的变量
	lisa := user{"Lisa", "lisa@email.com", 123, true}

	// 5-6 使用其他结构类型声明字段 admin 需要一个 user 类型作为管理者，并附加权限
	type admin struct {
		person user
		level  string
	}

	// 5-7 使用结构字面量来创建字段的值声明 admin 类型的变量
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true, 33},
		level: "super",
	}

	// 5-8 基于 int64 声明一个新类型
	type Duration int64

	var dur Duration
	dur = int64(1000)
	// prog.go:7: cannot use int64(1000) (type int64) as type Duration in assignment
}

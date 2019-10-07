// 7-21 pool/main/main.go
// 这个示例程序展示如何使用 pool 包
// 来共享一组模拟的数据库连接
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"pool"
)

const (
	maxGoroutines   = 25 // 要使用的 goroutine 的数量
	pooledResources = 2  // 池中的资源的数量
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close 实现了 io.Closer 接口，以便 dbConnection
// 可以被池管理。 Close 用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的 id
var idCounter int32

// createConnection 是一个工厂函数，
// 当需要一个新连接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

// main 是所有 Go 程序的入口
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		// 每个 goroutine 需要自己复制一份要
		// 查询值的副本，不然所有的查询会共享同一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// 等待 goroutine 结束
	wg.Wait()
	// 关闭池
	log.Println("Shutdown Program.")
	p.Close()
}

// performQueries 用来测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	// 从池里请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将该连接释放回池里
	defer p.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

// G:\GitHub\learnGo\Go语言实战\ch07_并发模式\runner\src\runner\main>set GOPATH=G:\GitHub\learnGo\Go语言实战\ch07_并发模式\pool
// G:\GitHub\learnGo\Go语言实战\ch07_并发模式\pool\src\pool\main>go run main.go
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 2
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 3
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 4
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 6
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 7
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 8
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 9
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 10
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 11
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 12
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 13
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 14
// 2019/10/07 18:39:46 Create: New Connection 1
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 15
// 2019/10/07 18:39:46 Create: New Connection 5
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 16
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 17
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 18
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 19
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 20
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 21
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 22
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 23
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 24
// 2019/10/07 18:39:46 Acquire: New Resource
// 2019/10/07 18:39:46 Create: New Connection 25
// 2019/10/07 18:39:46 QID[10] CID[6]
// 2019/10/07 18:39:46 Release: In Queue
// 2019/10/07 18:39:46 QID[15] CID[25]
// 2019/10/07 18:39:46 Release: In Queue
// 2019/10/07 18:39:46 QID[24] CID[2]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 2
// 2019/10/07 18:39:46 QID[12] CID[7]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 7
// 2019/10/07 18:39:46 QID[4] CID[15]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 15
// 2019/10/07 18:39:46 QID[1] CID[20]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 20
// 2019/10/07 18:39:46 QID[3] CID[1]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 1
// 2019/10/07 18:39:46 QID[13] CID[17]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 17
// 2019/10/07 18:39:46 QID[21] CID[19]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 19
// 2019/10/07 18:39:46 QID[9] CID[16]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 16
// 2019/10/07 18:39:46 QID[2] CID[24]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 24
// 2019/10/07 18:39:46 QID[20] CID[12]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 12
// 2019/10/07 18:39:46 QID[14] CID[8]
// 2019/10/07 18:39:46 Release: Closing
// 2019/10/07 18:39:46 Close: Connection 8
// 2019/10/07 18:39:47 QID[0] CID[9]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 9
// 2019/10/07 18:39:47 QID[18] CID[11]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 11
// 2019/10/07 18:39:47 QID[17] CID[18]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 18
// 2019/10/07 18:39:47 QID[19] CID[22]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 22
// 2019/10/07 18:39:47 QID[23] CID[14]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 14
// 2019/10/07 18:39:47 QID[11] CID[21]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 21
// 2019/10/07 18:39:47 QID[16] CID[10]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 10
// 2019/10/07 18:39:47 QID[6] CID[23]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 23
// 2019/10/07 18:39:47 QID[22] CID[13]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 13
// 2019/10/07 18:39:47 QID[8] CID[5]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 5
// 2019/10/07 18:39:47 QID[7] CID[4]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 4
// 2019/10/07 18:39:47 QID[5] CID[3]
// 2019/10/07 18:39:47 Release: Closing
// 2019/10/07 18:39:47 Close: Connection 3
// 2019/10/07 18:39:47 Shutdown Program.
// 2019/10/07 18:39:47 Close: Connection 6
// 2019/10/07 18:39:47 Close: Connection 25

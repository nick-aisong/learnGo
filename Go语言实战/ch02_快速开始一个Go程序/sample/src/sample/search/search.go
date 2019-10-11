package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run performs the search logic.
// Run 执行搜索逻辑
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	// 构造一个 waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	// 设置需要等待处理每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// Wait for everything to be processed.
		// 等候所有任务完成
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// 用关闭通道的方式，通知 Display 函数
		// function that we can exit the program.
		// 可以退出程序了
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	// 启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

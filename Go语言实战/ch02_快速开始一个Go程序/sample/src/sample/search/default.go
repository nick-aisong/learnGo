package search

// defaultMatcher implements the default matcher.
// defaultMatcher 实现了默认匹配器
type defaultMatcher struct{}

// init registers the default matcher with the program.
// defaultMatcher 实现了默认匹配器
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
// Search 实现了默认匹配器的行为
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

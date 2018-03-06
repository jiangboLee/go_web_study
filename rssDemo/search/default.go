package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (d defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

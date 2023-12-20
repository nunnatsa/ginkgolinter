package custommatcher

type customMatcher struct{}

func (customMatcher) Match(actual interface{}) (success bool, err error) {
	return true, nil
}
func (customMatcher) FailureMessage(actual interface{}) (message string) {
	return "FailureMessage"
}
func (customMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "NegatedFailureMessage"
}

func CustomMatcher() customMatcher {
	return customMatcher{}
}

func CustomMatcherP() *customMatcher {
	return &customMatcher{}
}

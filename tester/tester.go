package tester

import "testing"

type Tester struct {
	T *testing.T
}

// The fail message will be red and have a ✕
func (t Tester) Fail(msg string, variables ...interface{}) {
	coloredMsg := "\033[31m \u2715 "+msg
	t.T.Errorf(coloredMsg, variables...)
}

// The success message will be green and have a ✓
func (t Tester) Sucess(msg string, variables ...interface{}) {
	coloredMsg := "\033[32m \u2713 "+msg
	t.T.Logf(coloredMsg, variables...)
}
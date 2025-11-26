package wrappers

import (
	gtypes "github.com/onsi/gomega/types"

	g "github.com/onsi/gomega"
)

var Expect = g.Expect

func MyExpect(value any) g.Assertion {
	return g.Expect(value)
}

func MyEventually(value any) g.AsyncAssertion {
	return g.Eventually(value)
}

func MyConsistently(value any) g.AsyncAssertion {
	return g.Consistently(value)
}

type WrapAssertion struct {
	g.Assertion
}

type WrapAsyncAssertion struct {
	g.AsyncAssertion
}

func NewAssertion(value any) WrapAssertion {
	a := g.Expect(value) // Okay!
	return WrapAssertion{a}
}

func NewAsyncAssertion(value any) WrapAsyncAssertion {
	a := g.Eventually(value) // Okay!
	return WrapAsyncAssertion{a}
}

var Equal = g.Equal

func MyEqual(value any) gtypes.GomegaMatcher {
	return g.Equal(value)
}

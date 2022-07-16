package noginkgo

// this test checks that files with no ginkgo import, does not trigger warning.
// It mimics the gomega API, so the pattern should be the same.

type fakeAssertion struct{}

func (fakeAssertion) Should(_ bool) {

}

func Expect(i int) fakeAssertion {
	return fakeAssertion{}
}

func Equal(_ int) bool {
	return true
}

func test() {
	x := []int{1, 2, 3, 4, 5}
	Expect(len(x)).Should(Equal(5))
}

package ginkgolinter

const doc = `enforces standards of using ginkgo and gomega

or
       ginkgolinter version

version: %s

currently, the linter searches for following:
* trigger a warning when using Eventually or Constantly with a function call. This is in order to prevent the case when 
  using a function call instead of a function. Function call returns a value only once, and so the original value
  is tested again and again and is never changed. [Bug]

* trigger a warning when comparing a pointer to a value. [Bug]

* trigger a warning for missing assertion method: [Bug]
	Eventually(checkSomething)

* trigger a warning when a ginkgo focus container (FDescribe, FContext, FWhen or FIt) is found. [Bug]

* validate the MatchError gomega matcher [Bug]

* trigger a warning when using the Equal or the BeIdentical matcher with two different types, as these matchers will
  fail in runtime.

* wrong length assertions. We want to assert the item rather than its length. [Style]
For example:
	Expect(len(x)).Should(Equal(1))
This should be replaced with:
	Expect(x)).Should(HavelLen(1))
	
* wrong nil assertions. We want to assert the item rather than a comparison result. [Style]
For example:
	Expect(x == nil).Should(BeTrue())
This should be replaced with:
	Expect(x).Should(BeNil())

* wrong error assertions. For example: [Style]
	Expect(err == nil).Should(BeTrue())
This should be replaced with:
	Expect(err).ShouldNot(HaveOccurred())

* wrong boolean comparison, for example: [Style]
	Expect(x == 8).Should(BeTrue())
This should be replaced with:
	Expect(x).Should(BeEqual(8))

* replaces Equal(true/false) with BeTrue()/BeFalse() [Style]

* replaces HaveLen(0) with BeEmpty() [Style]

* replaces Expect(...).Should(...) with Expect(...).To() [stype]
`

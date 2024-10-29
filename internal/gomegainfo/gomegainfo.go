package gomegainfo

const ( // gomega actual method names
	expect                 = "Expect"
	expectWithOffset       = "ExpectWithOffset"
	omega                  = "Ω"
	eventually             = "Eventually"
	eventuallyWithOffset   = "EventuallyWithOffset"
	consistently           = "Consistently"
	consistentlyWithOffset = "ConsistentlyWithOffset"
)

const ( // assertion methods
	to        = "To"
	toNot     = "ToNot"
	notTo     = "NotTo"
	should    = "Should"
	shouldNot = "ShouldNot"
)

var funcOffsetMap = map[string]int{
	expect:                 0,
	expectWithOffset:       1,
	omega:                  0,
	eventually:             0,
	eventuallyWithOffset:   1,
	consistently:           0,
	consistentlyWithOffset: 1,
}

func IsActualMethod(name string) bool {
	_, found := funcOffsetMap[name]
	return found
}

func ActualArgOffset(methodName string) int {
	funcOffset, ok := funcOffsetMap[methodName]
	if !ok {
		return -1
	}
	return funcOffset
}

func GetAllowedAssertionMethods(actualMethodName string) string {
	switch actualMethodName {
	case expect, expectWithOffset:
		return `"To()", "ToNot()" or "NotTo()"`

	case eventually, eventuallyWithOffset, consistently, consistentlyWithOffset:
		return `"Should()" or "ShouldNot()"`

	case omega:
		return `"Should()", "To()", "ShouldNot()", "ToNot()" or "NotTo()"`

	default:
		return ""
	}
}

var asyncFuncSet = map[string]struct{}{
	eventually:             {},
	eventuallyWithOffset:   {},
	consistently:           {},
	consistentlyWithOffset: {},
}

func IsAsyncActualMethod(name string) bool {
	_, ok := asyncFuncSet[name]
	return ok
}

func IsAssertionFunc(name string) bool {
	switch name {
	case to, toNot, notTo, should, shouldNot:
		return true
	}
	return false
}

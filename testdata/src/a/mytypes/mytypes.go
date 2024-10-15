package mytypes

type MyErr struct{}

func (e MyErr) Error() string {
	return "My Error"
}

func (e MyErr) String() string {
	return "MyErr"
}

var MyError = &MyErr{}

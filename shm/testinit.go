package shm

var (
	isTest = false
)

func setupTest() {
	isTest = true
}

func teardownTest() {
	isTest = false
}

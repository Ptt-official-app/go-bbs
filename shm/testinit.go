package shm

var (
	IsTest = false
)

func setupTest() {
	IsTest = true
}

func teardownTest() {
	IsTest = false
}

package bbs

var (
	isTest = false
)

func setupTest() {
	isTest = true
	initTestVars()
}

func teardownTest() {
	isTest = false
	freeTestVars()
}

func initTestVars() {
	if testOpenUserecFile1 == nil {
		testOpenUserecFile1 = make([]*Userec, 50)
		for i := 0; i < 50; i++ {
			testOpenUserecFile1[i] = testUserecEmpty
		}
		testOpenUserecFile1[0] = testUserec1
		testOpenUserecFile1[1] = testUserec2
		testOpenUserecFile1[2] = testUserec3
		testOpenUserecFile1[3] = testUserec4
		testOpenUserecFile1[4] = testUserec5
	}
}

func freeTestVars() {
	testOpenUserecFile1 = nil
}

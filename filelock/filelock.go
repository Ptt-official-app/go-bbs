package filelock

type File interface{}

func Lock(f File) error {
	return nil
}

func RLock(f File) error {
	return nil
}

func IsLock(f File) bool {
	// TODO: Need to complete this function, it's just return false now.
	return false
}

func Unlock(f File) error {
	return nil
}

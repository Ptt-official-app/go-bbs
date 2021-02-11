package filelock

type File interface{}

func Lock(f File) error {
	return nil
}

func RLock(f File) error {
	return nil
}

func Unlock(f File) error {
	return nil
}

package cache

type Cache interface {
	Open()
	Close()
}

func NewCache(connectionString string) {

}

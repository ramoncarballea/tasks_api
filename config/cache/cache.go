package cache

type Storage interface {
	Add(key string)
	TryGet(key string) (*Collection, error)
	Remove(key string)
	Clear()
}

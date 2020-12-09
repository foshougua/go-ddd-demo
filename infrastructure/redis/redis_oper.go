package redis

type RedisOperInterfaces interface {
	Set(string) error
	Get(string) (string, error)
}

type RedisOperator struct{}

var _ RedisOperInterfaces = &RedisOperator{}

func NewRedisOperator() *RedisOperator {
	return &RedisOperator{}
}
func (ro *RedisOperator) Set(info string) error {
	return nil
}

func (ro *RedisOperator) Get(key string) (string, error) {
	return "", nil
}

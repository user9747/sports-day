package conf

/*
Redis Configurations
*/
type CacheConf struct {
	EnableSSL bool
	Addr      string
}

func getRedisAddr() string {
	return "localhost:6379"
	// return "host.docker.internal:6379"
}

var RedisConf = CacheConf{
	Addr:      getRedisAddr(),
	EnableSSL: false,
}

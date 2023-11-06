package conf

/*
Redis Configurations
*/
type CacheConf struct {
	EnableSSL bool
	Addr      string
}

func getRedisAddr() string {
	if MODE == MODE_LOCAL_DOCKER {
		return "host.docker.internal:6379"
	}
	return "localhost:6379"
}

var RedisConf = CacheConf{
	Addr:      getRedisAddr(),
	EnableSSL: false,
}

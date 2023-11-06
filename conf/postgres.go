package conf

type DBConf struct {
	Driver           string
	ConnectionString string
}

func getPostgresConnectionString() string {
	if MODE == MODE_LOCAL_DOCKER {
		return "postgres://postgres:postgres@host.docker.internal:5432/sports_day?sslmode=disable"
	}
	return "postgres://postgres:postgres@localhost:5432/sports_day?sslmode=disable"
}

var PostgresConf = DBConf{
	Driver:           "postgres",
	ConnectionString: getPostgresConnectionString(),
}

package conf

type DBConf struct {
	Driver           string
	ConnectionString string
}

var PostgresConf = DBConf{
	Driver:           "postgres",
	ConnectionString: "postgres://postgres:postgres@localhost:5432/sports_day?sslmode=disable",
}

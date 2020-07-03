package auth

// DbConfig структура хранит конфиги для подключения к БД
type DbConfig struct {
	Development struct {
		Dialect    string
		Datasource string
	}
}

var dbConfigFile = "dbconfig.yml"

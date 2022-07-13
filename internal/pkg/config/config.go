package config

// Config represents a link record.
type Config struct {
	SQLType   string `json:"sql_type"`
	NoSQLType string `json:"nosql_type"`
	HTTPPort  string `json:"http_port"`
	LogRoutes bool   `json:"log_routes"`
	Migrate   bool   `json:"migrate"`
	Debug     bool   `json:"debug"`
}

// NewConfig creates a new config struct.
// TODO: change this to be possible custom modification by user.
func NewConfig() (c Config) {
	c.SQLType = "sqlite"
	c.NoSQLType = "badger"
	c.HTTPPort = "8080"
	c.LogRoutes = true
	c.Migrate = true
	c.Debug = false
	return
}

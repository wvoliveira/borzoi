package config

// Config represents a link record.
type Config struct {
	HTTPPort string `json:"http_port"`
}

// NewConfig creates a new config struct.
func NewConfig() (c Config) {
	c.HTTPPort = "8080"
	return
}

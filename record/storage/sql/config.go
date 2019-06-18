package sqlstore

// Config is the configuration for the elasticsearch storage client
type Config struct {
	DriverName string
	Host string
	Port string
	DBName string
	Username string
	Password string
	RecordTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		DriverName: "",
		Host:      "",
		Port:       "",
		DBName:       "",
		Username:       "",
		Password:       "",
		RecordTableName: "",
	}
}
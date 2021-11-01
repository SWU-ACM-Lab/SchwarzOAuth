package application_module

type RedisConfig struct {
	Address  string // RedisAddress
	Password string // RedisPassword
	Port     string // RedisPort
	Database int    // RedisDatabaseIndex
}
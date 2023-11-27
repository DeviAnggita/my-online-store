package config

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	// 	RedisAddr  string
	// 	RedisPass  string
}

var AppConfig = Config{
	DBUsername: "root",
	DBPassword: "",
	DBHost:     "127.0.0.1",
	DBPort:     "3306",
	DBName:     "my-online-store",

	// RedisAddr: "localhost:6379",
	// RedisPass: "", // Set your Redis password if any
}

package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type KafkaConfig struct {
	Host  string
	Port  int
	Topic string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Service  ServiceConfig
	Kafka    KafkaConfig
	Postgres PostgresConfig
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{

		Service: ServiceConfig{
			Host: viper.GetString("service.host"),
			Port: viper.GetInt("service.port"),
		},
		Kafka: KafkaConfig{
			Host:  viper.GetString("kafka.host"),
			Port:  viper.GetInt("kafka.port"),
			Topic: viper.GetString("kafka.topic"),
		},
		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Database: viper.GetString("postgres.database"),
		},
	}
	return &cfg, nil
}

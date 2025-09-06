package config

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

type MongoConfig struct {
	Host       string
	Port       int
	Database   string
	Collection string
}

type Config struct {
	Service ServiceConfig
	Kafka   KafkaConfig
	Mongo   MongoConfig
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

		Mongo: MongoConfig{
			Host:       viper.GetString("mongodb.host"),
			Port:       viper.GetInt("mongodb.port"),
			Database:   viper.GetString("mongodb.database"),
			Collection: viper.GetString("mongodb.collection"),
		},
	}
	return &cfg, nil
}

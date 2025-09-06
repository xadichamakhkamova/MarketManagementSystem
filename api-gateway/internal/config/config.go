package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type MinioConfig struct {
	Host   string
	Port   int
	Id     string
	Secret string
}

type Config struct {
	ApiGateway       ServiceConfig
	ProductService   ServiceConfig
	DashboardService ServiceConfig
	DebtService      ServiceConfig
	Minio            MinioConfig
	FrontHost        string
	FrontPort        int
	TokenKey         string
	CertFile         string
	KeyFile          string
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		ApiGateway: ServiceConfig{
			Host: viper.GetString("api-gateway.host"),
			Port: viper.GetInt("api-gateway.port"),
		},
		ProductService: ServiceConfig{
			Host: viper.GetString("services.product-service.host"),
			Port: viper.GetInt("services.product-service.port"),
		},
		DashboardService: ServiceConfig{
			Host: viper.GetString("services.dashboard-service.host"),
			Port: viper.GetInt("services.dashboard-service.port"),
		},
		DebtService: ServiceConfig{
			Host: viper.GetString("services.debt-service.host"),
			Port: viper.GetInt("services.debt-service.port"),
		},
		Minio: MinioConfig{
			Host:   viper.GetString("minio.host"),
			Port:   viper.GetInt("minio.port"),
			Id:     viper.GetString("minio.id"),
			Secret: viper.GetString("minio.secret"),
		},
		FrontHost: viper.GetString("frontend.host"),
		FrontPort: viper.GetInt("frontend.port"),

		TokenKey: viper.GetString("token.key"),

		CertFile: viper.GetString("file.cert"),
		KeyFile:  viper.GetString("file.key"),
	}
	return &cfg, nil
}

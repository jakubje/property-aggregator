package util

import "github.com/spf13/viper"

type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBName            string `mapstructure:"DB_NAME"`
	CollectionName    string `mapstructure:"COLLECTION_NAME"`
	DBMongoUri        string `mapstructure:"DB_MONGO_URI"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	DiscordWebhook    string `mapstructure:"DISCORD_WEBHOOK"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		return
	}

	err = viper.Unmarshal(&config)
	return
}

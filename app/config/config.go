package config

import "github.com/spf13/viper"

type Config struct {
	IncludeModels           []string `mapstructure:"INCLUDE_MODELS"`
	CopyTo                  string   `mapstructure:"COPY_TO"`
	CopyIncludeRepositories []string `mapstructure:"COPY_INCLUDE_REPOSITORIES"`
	ConfirmOverwrite        bool     `mapstructure:"CONFIRM_OVERWRITE"`
	ConfirmAllFiles         bool     `mapstructure:"CONFIRM_ALL_FILES"`
}

func NewConfig() (*Config, error) {
	// conf := Config{}
	// if err := env.Parse(&conf); err != nil {
	// 	return nil, err
	// }

	// return &conf, nil

	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{}
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

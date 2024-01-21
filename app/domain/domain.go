package domain

import "github.com/spf13/viper"

type Domain struct {
	ServiceName   string   `json:"service_name" mapstructure:"service_name"`
	AggregateName string   `json:"aggregate_name" mapstructure:"aggregate_name"`
	DomainName    string   `json:"domain_name" mapstructure:"domain_name"`
	Methods       []Method `json:"methods" mapstructure:"methods"`
}

type Method struct {
	MethodName string   `json:"method_name" mapstructure:"method_name"`
	Request    *Payload `json:"request" mapstructure:"request"`
	Response   *Payload `json:"response" mapstructure:"response"`
}

type Payload struct {
	PayloadName string         `json:"payload_name" mapstructure:"payload_name"`
	Data        []*PayloadData `json:"data" mapstructure:"data"`
	//EmbeddedPayload  *Payload      `json:"embedded_payload"`
	//EmbeddedPayloads []*Payload    `json:"embedded_payloads"`
}

type PayloadData struct {
	Name      string `json:"name" mapstructure:"name"`
	Type      string `json:"type" mapstructure:"type"`
	IsNotNull bool   `json:"is_not_null" mapstructure:"is_not_null"`
}

func NewDomain(dir string, fileName string) (*Domain, error) {
	v := viper.New()
	v.AddConfigPath(dir)
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var dml Domain
	if err := v.Unmarshal(&dml); err != nil {
		return nil, err
	}

	return &dml, nil
}

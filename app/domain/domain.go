package domain

import (
	"vl-template/app/config"
)

type Domain struct {
	ServiceName     string    `json:"service_name" mapstructure:"service_name"`
	AggregateName   string    `json:"aggregate_name" mapstructure:"aggregate_name"`
	DomainName      string    `json:"domain_name" mapstructure:"domain_name"`
	QueryMethods    []*Method `json:"query_methods" mapstructure:"query_methods"`
	MutationMethods []*Method `json:"mutation_methods" mapstructure:"mutation_methods"`
}

type Method struct {
	Name         string     `json:"name" mapstructure:"name"`
	RequestName  string     `json:"request_name" mapstructure:"request_name"`
	ResponseName string     `json:"response_name" mapstructure:"response_name"`
	Requests     []*Payload `json:"requests" mapstructure:"requests"`
	Responses    []*Payload `json:"responses" mapstructure:"responses"`
}

type Payload struct {
	PayloadName string         `json:"name" mapstructure:"name"`
	Data        []*PayloadData `json:"data" mapstructure:"data"`
	//EmbeddedPayload  *Payload       `json:"embedded_payload"`
	//EmbeddedPayloads []*Payload     `json:"embedded_payloads"`
}

type PayloadData struct {
	Name      string `json:"name" mapstructure:"name"`
	Type      string `json:"type" mapstructure:"type"`
	IsNotNull bool   `json:"is_not_null" mapstructure:"is_not_null"`
}

type DomainReader struct {
	Conf       *config.Config
	DomainsMap map[string][]*Domain
}

func NewDomainReader(conf *config.Config) *DomainReader {
	return &DomainReader{
		Conf:       conf,
		DomainsMap: map[string][]*Domain{},
	}

}

func (r *DomainReader) ReadFromSpec(dir string) error {
	domainsSpecReader := NewDomainsSpecReader()
	if err := domainsSpecReader.GetDomains(dir); err != nil {
		return err
	}
	r.DomainsMap = MergeDomainsMap(r.DomainsMap, domainsSpecReader.DomainsMap)
	return nil
}

func (r *DomainReader) ReadFromProto(dir string) error {
	domainsSpecReader := NewDomainProtoReader()
	if err := domainsSpecReader.GetDomains(dir); err != nil {
		return err
	}
	r.DomainsMap = MergeDomainsMap(r.DomainsMap, domainsSpecReader.DomainsMap)
	return nil
}

func MergeDomainsMap(map1, map2 map[string][]*Domain) map[string][]*Domain {
	rt := map[string][]*Domain{}
	for key, value := range map1 {
		rt[key] = value
	}
	for key, value := range map2 {
		if _, ok := map2[key]; ok {
			rt[key] = append(rt[key], value...)
			continue
		}
		rt[key] = value
	}
	return rt
}

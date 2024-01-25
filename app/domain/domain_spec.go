package domain

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

type DomainsSpecReader struct {
	DomainsMap map[string][]*Domain
}

func NewDomainsSpecReader() *DomainsSpecReader {
	return &DomainsSpecReader{
		DomainsMap: map[string][]*Domain{},
	}
}

func (r *DomainsSpecReader) newDomain(dir string, fileName string) (*Domain, error) {
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

func (r *DomainsSpecReader) GetDomains(dir string) error {
	if err := r.walkDir(dir); err != nil {
		return err
	}
	return nil
}

func (r *DomainsSpecReader) walkDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			childDir := path.Join(dir, fileName)
			r.walkDir(childDir)
		} else {
			dm, err := r.newDomain(dir, fileName)
			if err != nil {
				return err
			}

			r.DomainsMap[dm.ServiceName] = append(r.DomainsMap[dm.ServiceName], dm)
		}
	}
	return nil
}

package conf

import (
	"log"
	"os"

	"github.com/micro/go-micro/v2/config/encoder"
	"github.com/micro/go-micro/v2/config/encoder/json"
	"github.com/micro/go-micro/v2/config/encoder/toml"
	"github.com/micro/go-micro/v2/config/encoder/xml"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/zzsds/micro-utils/config/nacos"
)

var (
	Conf       = new(Config)
	ConfigPath string
)

// Config 配置
type Config struct {
	Db *Db
}

// Mysql 数据库
type Db struct {
	Name     string
	User     string
	Password string
	Host     string
	Charset  string
	Prefix   string
	Debug    bool
}

// InitConfig ...
func InitConfig(ops ...source.Option) *Config {
	var (
		coder  encoder.Encoder
		source source.Source
	)
	source = nacos.NewAutoSource(ops...)
	sc, err := source.Read()
	if err != nil {
		log.Fatal(err)
	}

	switch sc.Format {
	case "toml":
		coder = toml.NewEncoder()
	case "json":
		coder = json.NewEncoder()
	case "xml":
		coder = xml.NewEncoder()
	case "yaml":
		coder = yaml.NewEncoder()
	default:
		coder = json.NewEncoder()
	}
	if err := coder.Decode(sc.Data, Conf); err != nil {
		log.Fatal(err)
	}
	return Conf
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil && !os.IsNotExist(err)
}

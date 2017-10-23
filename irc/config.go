package irc

import (
	"errors"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type PassConfig struct {
	Password string
}

type TLSConfig struct {
	Key  string
	Cert string
}

func (conf *PassConfig) PasswordBytes() []byte {
	bytes, err := DecodePassword(conf.Password)
	if err != nil {
		log.Fatal("decode password error: ", err)
	}
	return bytes
}

type Config struct {
	Server struct {
		PassConfig `yaml:",inline"`
		Database   string
		Listen     []string
		TLSListen  map[string]*TLSConfig
		Log        string
		MOTD       string
		Name       string
	}

	Operator map[string]*PassConfig
}

func (conf *Config) Operators() map[Name][]byte {
	operators := make(map[Name][]byte)
	for name, opConf := range conf.Operator {
		operators[NewName(name)] = opConf.PasswordBytes()
	}
	return operators
}

func LoadConfig(filename string) (config *Config, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	if config.Server.Name == "" {
		return nil, errors.New("Server name missing")
	}
	if len(config.Server.Listen)+len(config.Server.TLSListen) == 0 {
		return nil, errors.New("Server listening addresses missing")
	}
	return config, nil
}

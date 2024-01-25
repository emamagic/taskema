package config

import (
	"taskema/pkg/richerror"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load() (Config, error) {
	op := "config.Load"
	var k = koanf.New(".")

	if err := k.Load(confmap.Provider(defaultConfig, "."), nil); err != nil {

		return Config{},
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected).
				WithMeta(map[string]interface{}{"meta": "loading config from defaultConfig"})
	}

	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {

		return Config{},
		richerror.New(op).
			WithMessage(err.Error()).
			WithCode(richerror.CodeUnexpected).
			WithMeta(map[string]interface{}{"meta": "loading config from config.yml"})
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {

		return Config{},
		richerror.New(op).
			WithMessage(err.Error()).
			WithCode(richerror.CodeUnexpected).
			WithMeta(map[string]interface{}{"meta": "unmarshaling config"})
	}

	return cfg, nil
}

package config

import (
	"bytes"
	"os"

	"github.com/dereference-xyz/trickle/model"
	"gopkg.in/yaml.v3"
)

const DecoderFilePath = "js/decoder/anchor/dist/decoder.js"

type Config struct {
	Version  int      `yaml:"version"`
	Database Database `yaml:"database"`
	Chains   []Chain  `yaml:"chains"`
}

type Database struct {
	SQLite *SQLite `yaml:"sqlite"`
}

type SQLite struct {
	File string `yaml:"file"`
}

type Chain struct {
	Solana *Solana `yaml:"solana"`
}

type Solana struct {
	Node     string    `yaml:"node"`
	Programs []Program `yaml:"programs"`
}

type Program struct {
	ProgramId string `yaml:"program_id"`
	IDL       string `yaml:"idl"`
}

func Parse(file string) (*Config, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	dec := yaml.NewDecoder(bytes.NewReader(dat))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validate(cfg *Config) error {
	if len(cfg.Chains) != 1 {
		return model.NewInputValidationError("Only exactly one chain config is supported at this time.")
	}

	if len(cfg.Chains[0].Solana.Programs) != 1 {
		return model.NewInputValidationError("Only exactly one program config is supported at this time.")
	}

	return nil
}

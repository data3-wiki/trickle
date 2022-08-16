package config

import (
	"bytes"
	"os"

	"github.com/dereference-xyz/trickle/model"
	"gopkg.in/yaml.v3"
)

// Path to anchor decoder (built as a var-type library).
const DecoderFilePath = "js/decoder/anchor/dist/decoder.js"

// User-defined config.
type Config struct {
	// Version of the yaml schema.
	Version int `yaml:"version"`
	// Config of database to use for storage.
	Database Database `yaml:"database"`
	// Chain-specific config for smart contract data to load.
	Chains []Chain `yaml:"chains"`
}

// Config of database to use for storage.
// Only one type of db should have a config specified.
type Database struct {
	// Config for SQLite.
	SQLite *SQLite `yaml:"sqlite"`
}

// SQLite-specific config.
type SQLite struct {
	// File path to db.
	File string `yaml:"file"`
}

// Chain-specific config for smart contract data to load.
// Only one chain should have a config specified.
type Chain struct {
	// Solana-specific config.
	Solana *Solana `yaml:"solana"`
}

// Solana-specific config.
type Solana struct {
	// URL of RPC node to load data from.
	Node string `yaml:"node"`
	// List of programs to load.
	Programs []Program `yaml:"programs"`
}

// Config for a specific Solana program (smart contract) to load data for.
type Program struct {
	// Public key of program.
	ProgramId string `yaml:"program_id"`
	// Path to idl.json file to use for decoding.
	IDL string `yaml:"idl"`
}

// Parse the yaml file specified by the given path into the Config struct.
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

// Run input validation after parsing.
func validate(cfg *Config) error {
	if len(cfg.Chains) != 1 {
		return model.NewInputValidationError("Only exactly one chain config is supported at this time.")
	}

	if len(cfg.Chains[0].Solana.Programs) != 1 {
		return model.NewInputValidationError("Only exactly one program config is supported at this time.")
	}

	return nil
}

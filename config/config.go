package config

import "database/sql"

type Config interface {
	Setting() Setting
	DB() *sql.DB
}

var _ Config = (*DefaultConfig)(nil)

type DefaultConfig struct {
	setting Setting
	db      *sql.DB
}

func (c *DefaultConfig) Setting() Setting {
	return c.setting
}

func (c *DefaultConfig) DB() *sql.DB {
	return c.db
}

func NewConfig(setting Setting, db *sql.DB) Config {
	return &DefaultConfig{
		setting: setting,
		db:      db,
	}
}

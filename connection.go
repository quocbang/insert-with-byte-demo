package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

func InitializeDB(cfs Postgres) (*gorm.DB, error) {
	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfs.Address,
		cfs.UserName,
		cfs.Password,
		cfs.Name,
		cfs.Port,
	)
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&YamlDataSave{}); err != nil {
		return err
	}
	return nil
}

package main

import type_ "quocbang/save-file-to-db/type"

type YamlDataSave struct {
	ID        string     `gorm:"type:text;primaryKey"`
	Type      type_.Type `gorm:"not null"`
	Content   []byte     `gorm:"type:bytea;not null"`
	CreatedBy string     `gorm:"type:text;not null"`
	CreatedAt int64      `gorm:"type:bigint;autoCreateTime"`
}

func (YamlDataSave) TableName() string {
	return "yaml_data_save"
}

type StationFileFormat struct {
	ID           string `yaml:"ID"`
	SubCompany   int    `yaml:"sub-company"`
	Factory      string `yaml:"factory"`
	DepartmentID string `yaml:"department-id"`
	Alias        string `yaml:"alias"`
	SerialNumber int    `yaml:"serial-number"`
	Description  string `yaml:"description"`
	Devices      []int  `yaml:"devices"`
}

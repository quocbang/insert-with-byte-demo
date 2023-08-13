package main

import (
	"fmt"
	"log"
	"os"
	type_ "quocbang/save-file-to-db/type"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

func main() {
	pgConnection := Postgres{
		Name:     "test",
		Address:  "localhost",
		Port:     5432,
		UserName: "test",
		Password: "test",
		Schema:   "public",
	}

	db, err := InitializeDB(pgConnection)
	if err != nil {
		log.Fatal(err)
	}

	// migrate table
	if err := MigrateTable(db); err != nil {
		log.Fatal(err)
	}

	stationData, err := os.ReadFile("station.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Create(&YamlDataSave{
		ID:        "KU-P3211-SM1518-S02-1",
		Type:      type_.Type_STATION,
		CreatedBy: "quocbang",
		Content:   stationData,
	}).Error; err != nil {
		log.Fatal(err)
	}

	station, err := getStation("KU-P3211-SM1518-S02-1", db)
	if err != nil {
		log.Fatal(err)
	}

	stationParse := StationFileFormat{}
	if err := yaml.Unmarshal(station.Content, &stationParse); err != nil {
		log.Fatal(err)
	}

	log.Println(stationParse)

	log.Println("create station successfully")
}

func getStation(id string, db *gorm.DB) (YamlDataSave, error) {
	if id == "" || db == nil {
		return YamlDataSave{}, fmt.Errorf("missing request")
	}
	yamlData := YamlDataSave{}
	if err := db.Where(`id=?`, id).Take(&yamlData).Error; err != nil {
		return YamlDataSave{}, err
	}
	return yamlData, nil
}

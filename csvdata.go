package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Volcano struct {
	Number        string `json:"Number"`
	Name     string `json:"Name"`
	Country   string `json:"Country"`
	Region string `json:"Region"`
	Type  string `json:"Type"`
	ActivityEvidence     string `json:"ActivityEvidence"`
	LastKnownEruption string `json:"LastKnownEruption"`
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
	Elevation string `json:"Elevation"`
	DominantRockType string `json:"DominantRockType"`
	TectonicSetting string `json:"TectonicSetting"`
}

var volcanoList []Volcano

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readData(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	volcanoList = []Volcano{}

	for _, record := range records {
		volcano := Volcano{
			Number:            record[0],
			Name:              record[1],
			Country:           record[2],
			Region:            record[3],
			Type:              record[4],
			ActivityEvidence:  record[5],
			LastKnownEruption: record[6],
			Latitude:          record[7],
			Longitude:         record[8],
			Elevation:         record[9],
			DominantRockType:  record[10],
			TectonicSetting:   record[11],
		}
			volcanoList = append(volcanoList, volcano)
	}
	file.Close()
}

func writeData(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, volcano := range volcanoList {
		record := []string{
			volcano.Number,
			volcano.Name,
			volcano.Country,
			volcano.Region,
			volcano.Type,
			volcano.ActivityEvidence,
			volcano.LastKnownEruption,
			volcano.Latitude,
			volcano.Longitude,
			volcano.Elevation,
			volcano.DominantRockType,
			volcano.TectonicSetting,
		}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}

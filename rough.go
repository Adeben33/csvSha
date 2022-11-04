package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//JSON struct for marshalling
type attribute1 struct {
	TraitType string `json:"TraitType"`
	Value     string `json:"Value"`
}
type attribute2 struct {
	Type  string `json:"type"`
	Value string `json:"Value"`
}

type collection struct {
	Name                string       `json:"Name"`
	Id                  string       `json:"Id"`
	Collectionattribute []attribute2 `json:"Collectionattribute"`
}
type csvJson1 struct {
	Format            string       `json:"format"`
	Name              string       `json:"Team name"`
	FileName          string       `json:"fileName"`
	Description       string       `json:"description"`
	Gender            string       `json:"gender"`
	MintingTool       string       `json:"minting_tool"`
	SensitiveContent  bool         `json:"sensitiveContent"`
	SeriesNumber      string       `json:"SeriesNumber"`
	SeriesTotal       int          `json:"series_total"`
	Attributes        []attribute1 `json:"attributes"`
	Collections       []string     `json:"collection"`
	UUID              string       `json:"UUID"`
	FilenameOutputCsv [32]byte     `json:"filename.output.csv"`
}

//Main this is the entry point for the programme
func main() {
	//var allFiles []csvJson1
	//Opening of the CSV file using
	file, err := os.Open("data.csv")
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("Open successfully")
	defer file.Close()
	//Reading the file for the purpose of harshing
	records1, _ := ioutil.ReadFile("data.csv")
	sum := sha256.Sum256(records1)

	//records returns an io.reader
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//looping to fill in the struct
	for _, line := range records {
		var item csvJson1
		item.Format = "CHIP-0007"
		item.MintingTool = line[0]
		item.SeriesNumber = line[1]
		item.FileName = line[2]
		item.Name = line[3]
		item.Description = line[4]
		item.Gender = line[5]
		item.SensitiveContent = false
		item.SeriesTotal = 420
		item.UUID = line
		item.Attributes
		item.Collections
		item.FilenameOutputCsv = sum
		//allFiles = append(allFiles, item)

		jsonData, err := json.Marshal(item)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
	}
	//Changing the slice into json
	//jsonData, err := json.Marshal(&allFiles)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Print(string(jsonData))
}

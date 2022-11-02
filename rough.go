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
type csvJson1 struct {
	Format            string   `json:"format"`
	SeriesNumber      string   `json:"seriesNumber"`
	CurrentName       string   `json:"currentName"`
	Name              string   `json:"name"`
	The               string   `json:"the"`
	Descriptor        string   `json:"descriptor"`
	NewName           string   `json:"newName"`
	Description       string   `json:"description"`
	UUID              string   `json:"UUID"`
	FilenameOutputCsv [32]byte `json:"filename.output.csv"`
}

//Main this is the entry point for the programme
func main() {
	var allFiles []csvJson1
	//Opening of the CSV file using
	file, err := os.Open("data.csv")
	if err != nil {
		log.Println(err.Error())
		//c.Status(http.StatusUnprocessableEntity)
		return
	}
	fmt.Println("Open successfully")
	defer file.Close()
	//Reading the file for the purpose of harshing
	records1, _ := ioutil.ReadFile("data.csv"
	//sum is the sha256 harsh
	sum = sha256.Sum256(records1)

	//records returns an io.reader
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err.Error())
		//c.Status(http.StatusUnprocessableEntity)
		return
	}
	//looping to fill in the struct
	for _, line := range records {
		var item csvJson1
		item.Format = "format"
		item.SeriesNumber = line[0]
		item.CurrentName = line[1]
		item.Name = line[2]
		item.The = line[3]
		item.Descriptor = line[4]
		item.NewName = line[5]
		item.Description = line[6]
		item.UUID = line[7]
		item.FilenameOutputCsv = sum
		allFiles = append(allFiles, item)
	}
	//Changing the slice into json
	jsonData, err := json.Marshal(&allFiles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(jsonData))
}

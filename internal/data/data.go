package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Leblanc-PiR/FizzBuzz/config"
)

const (
	Int1ParamStr  = "int1"
	Int2ParamStr  = "int2"
	LimitParamStr = "limit"
	Str1ParamStr  = "str1"
	Str2ParamStr  = "str2"
)

type fieldName string

const (
	fieldID               fieldName = "id"
	fieldFormattedRequest fieldName = "formattedRequest"
	fieldHits             fieldName = "hits"
)

// JSONData represents an JSON object translated into go-usable object
type JSONData struct {
	ID               int    `json:"id"`
	FormattedRequest string `json:"formattedRequest"`
	Hits             int    `json:"hits"`
}

// InitialisingDB connects (or create if doesn't exist) to JSON pseudoDB file
//
//	(Could handle migration if need be)
//	(Interested by parquet files instead of json, still short on time)
func InitialisingDB(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Printf("File %q not found\n", filename)

		// Attempting to create file if not found
		jsonFile, err := os.Create(filename)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		log.Printf("Created file: %s\n", filename)

		jsonFile.WriteString("[]")
		jsonFile.Close()
	}
}

// RecordFizzBuzzCall record fizzBuzz call values
func RecordFizzBuzzCall(int1, int2, lim int, str1, str2 string) {
	DBEntries := getDBDatas()

	DBSize := len(DBEntries)

	formattedRequest := fmt.Sprintf("%s=%d&%s=%d&%s=%d&%s=%s&%s=%s",
		Int1ParamStr,
		int1,
		Int2ParamStr,
		int2,
		LimitParamStr,
		lim,
		Str1ParamStr,
		str1,
		Str2ParamStr,
		str2,
	)

	foundRow := findDataRowByField(DBEntries, fieldFormattedRequest, formattedRequest)

	indexToIncrement := 0
	if DBSize > 0 {
		indexToIncrement = DBSize - 1
	}

	if foundRow == -1 {
		idToInsert := 0
		if DBSize > 0 {
			idToInsert = DBEntries[indexToIncrement].ID + 1
		}
		// Record new entry
		newRequest := JSONData{
			ID:               idToInsert,
			FormattedRequest: formattedRequest,
			Hits:             1,
		}

		DBEntries = append(DBEntries, newRequest)

		writeDataToDB(DBEntries, config.DBFilename)
	} else {
		// Increment value of Hits for found formattedRequest
		DBEntries[foundRow].Hits = DBEntries[foundRow].Hits + 1

		writeDataToDB(DBEntries, config.DBFilename)
	}
}

// findDataRowByField gives the id of the searched FormattedRequest if found, -1 if not
func findDataRowByField(dataSet []JSONData, field fieldName, searchedValue any) int {

	for i, row := range dataSet {
		if field == fieldFormattedRequest && row.FormattedRequest == searchedValue {
			return i
		}
		if field == fieldID && row.ID == searchedValue {
			return i
		}
		if field == fieldHits && row.Hits == searchedValue {
			return i
		}
	}

	return -1
}

// getDBDatas Reads and translates filedata
func getDBDatas() []JSONData {
	// Open "DB"
	content, err := ioutil.ReadFile(config.DBFilename)
	if err != nil {
		log.Fatal(err)
	}
	DBentries := []JSONData{}

	// Translate datas
	err = json.Unmarshal(content, &DBentries)
	if err != nil {
		log.Fatal(err)
	}

	return DBentries
}

// writeDataToDB write updated data to our JSON file
func writeDataToDB(data []JSONData, file string) {
	// Marshal data
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(fmt.Errorf("error marshaling: %w", err))
	}

	// Write json to file
	err = ioutil.WriteFile(file, dataBytes, 0644)
	if err != nil {
		log.Println(fmt.Errorf("error writing to DB: %w", err))
	}

}

// findDatasWithHighestHits search dataset and returns data(s) having the most hits
func findDatasWithHighestHits(dataSet []JSONData) []JSONData {
	res := []JSONData{{}}

	for _, row := range dataSet {
		if row.Hits > res[0].Hits {
			res[0] = row
		} else if row.Hits == res[0].Hits {
			res = append(res, row)
		}
	}

	return res
}

// GetHighestHitsRequestParams return request(s) having the most hits
func GetHighestHitsRequestParams() []JSONData {
	DBentries := getDBDatas()

	return findDatasWithHighestHits(DBentries)
}

package data

import (
	"log"
	"os"

	"github.com/Leblanc-PiR/FizzBuzz/config"
)

// InitialisingPseudoDB connects (or create if doesn't exist) to CSV pseudoDB file
// Could have migration to initialise at least column name
// Interested by parquet files instead of csv, still short on time.
func InitialisingPseudoDB(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("File %q not found\n", filename)

		// Attempting to create file if not found
		csvFile, err := os.Create(filename)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		log.Printf("Created file: %s\n", filename)

		defer csvFile.Close()
	}

	file.Close()
}

// RecordFizzBuzzCall record fizzBuzz call values
func RecordFizzBuzzCall(int1, int2, lim int, str1, str2 string) {
	file, err := os.OpenFile(config.DBFilename, os.O_RDWR, 0)
	if err != nil {
		log.Panicf("could not read frome %s: %s\n", config.DBFilename, err.Error())
		return
	}
	defer file.Close()

}

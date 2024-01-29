package helpers

import (
	"fmt"
	"log"
	"os"
)



func GetUnits() (string){
	var units string
	f, err := os.ReadFile("./units.txt")
	if err != nil {
		return ""
	}
	
	fmt.Println(f)
	units = string(f)
	return units
}

func WriteUnits(units string) string {
	err := os.WriteFile("./units.txt", []byte(units), 0644)
	if err != nil {
		log.Fatalf("Error occured while writing units.txt: %s\n", err)
		os.Exit(1)
	}
	return units
}
package core

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

var teams []string
var features []string
var cases []Case

func initData(data string) {

	r := csv.NewReader(strings.NewReader(data))

	isFirst := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if (isFirst) {
			for i, value := range record {
				if i==0 {
					continue
				}
				features = append(features, value)
			}
			isFirst = false
			continue
		}


		isFound, index := findTeam(record[0])
		if !isFound {
			teams = append(teams, record[0])
			index = len(teams)-1
		}
		icase := Case{team:index}

		for i, value := range record {
			if i==0 {
				continue
			}
			icase.features = append(icase.features, mapValue(value))
		}

		cases = append(cases, icase)
	}
}

func mapValue(value string) int{
	switch value{
	case "Y":
		return YES
	case "N":
		return NO
	default:
		return NA
	}
}

func findTeam(name string) (bool, int){
	for i,v := range teams {
		if (v==name) {
			return true, i
		}
	}
	return false, 0
}

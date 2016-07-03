package models

import (
	"encoding/csv"
	"io"
	"log"
	"github.com/astaxie/beego"
	"fmt"
	"os"
)

var teams []string
var features []string
var cases []Case

func InitCore() error {
	csv := beego.AppConfig.String("datafilename")
	fmt.Println("LOADING ", csv)
	initData(csv)

	fmt.Println(teams)
	fmt.Println(cases)
	fmt.Println(features)

	return nil
}

func initData(path string) {
	f, _ := os.Open(path)

	r := csv.NewReader(f)

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
				if i<=1 {
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
		icase := Case{team:index, label:record[1]}

		for i, value := range record {
			if i<=1 {
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

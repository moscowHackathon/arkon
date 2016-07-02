package models

import (
	"testing"

	"fmt"
)

func TestA(t *testing.T) {
	var cases = []Case{
		Case{
			team:1,
			features: []int{YES,YES,NA,NO,YES,YES,NO},
		},
		Case{
			team:1,
			features: []int{YES,YES,NA,NO,YES,YES,YES},
		},
		Case{
			team:2,
			features: []int{NA,NO,NA,NO,YES,YES,NO},
		},
		Case{
			team:2,
			features: []int{NA,NO,NA,YES,YES,YES,NA},
		},
		Case{
			team:2,
			features: []int{NA,NA,NA,YES,YES,YES,NO},
		},
		Case{
			team:2,
			features: []int{NA,NA,NA,YES,NA,YES,YES},
		},

	}

	calc:= newSessionCalc(cases)
	i:=calc.GetNextToCheck()
	Assert(t, i == 3)

	calc.ApplyAnswer(NO)
	Assert(t, len(calc.cases) == 3)
	fmt.Println(calc.cases)

	i=calc.GetNextToCheck()
	Assert(t, i == 6)

	calc.ApplyAnswer(YES)

	fmt.Println(calc.cases)

	fmt.Println(calc.CheckStatus())
	fmt.Println(cases)
	fmt.Println(calc)
}

func TestInit(t *testing.T) {
	data := `,is cart page,is checkout,is catalog,is home page down
bob_api,,Y,N,
catalog,N,N,Y,Y
content,,,,Y`
	initData(data)
	fmt.Println(teams)
	fmt.Println(features)
	fmt.Println(len(features))

	fmt.Println(cases)

}

func Assert(t *testing.T, condition bool, args ...interface{}) {
	if !condition {
		t.Error(args)
	}
}
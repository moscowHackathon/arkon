package core

import(
	"math"
	"fmt"
)

const (
	_  = iota
	YES
	NO
	NA
)


/*c := bayesian.NewClassifier("Team1", "Team2")
	c.Learn([]string{"catalog", "price", "images"}, "Team1")
	c.Learn([]string{"cart", "checkout", "shipping"}, "Team2")*/

var cases = []Case{Case{
		team:1,
		features: []int{1,1,0,-1,1,1,-1},
	},
	Case{
		team:2,
		features: []int{1,1,0,-1,1,1,-1},
	},
}
type Case struct {
	team int
	features []int
}

type SessionCalc struct {
	next int
	cases []Case
	answers []int
}

func newSessionCalc(cases []Case) SessionCalc {
	lf := len(cases[0].features)
	f := make([]int, lf)
	return  SessionCalc{
		cases:cases,
		answers: f,
	}
}

func (s *SessionCalc) GetNextToCheck() int{
	lf := len(s.cases[0].features)
	variant := make([]int, lf)
	variantTrue := make([]int, lf)
	variantFalse := make([]int, lf)
	for _, icase :=range s.cases {
		for i, val := range icase.features {
			if val==YES {
				variantTrue[i]+=1
			}
			if val==NO {
				variantFalse[i]+=1
			}
		}
	}
	l:=int(len(s.cases)/2)
	for i,_:=range variantTrue {
		variant[i] = abs(l-variantTrue[i])+abs(l-variantTrue[i])
	}

	fmt.Println(variantTrue)
	fmt.Println(variantFalse)
	fmt.Println(variant)

	res := getMinIndex(variant)
	s.next = res

	return res
}

func (s *SessionCalc) ApplyAnswer(ans int) {
	s.answers[s.next] = ans
	if ans==NA  {
		return
	}
	newCases := []Case{}
	for i,icase:=range s.cases {
		if icase.features[s.next]==NA || icase.features[s.next]==ans {
			newCases = append(newCases, s.cases[i])
		}
	}
	s.cases = newCases
}

func (s *SessionCalc) CheckStatus() (isFinish bool, result int){
	if s.isOneTeam() {
		return true, s.cases[0].team
	}
	return false, 0
}

func (s *SessionCalc) isOneTeam() bool{
	team:=s.cases[0].team
	for _,icase:=range s.cases {
		if icase.team!=team {
			return false
		}
	}
	return true
}

func getMinIndex(arr []int)int{
	min := arr[0]
	ret:=0

	for i, value := range arr {
		if value < min {
			min = value
			ret = i
		}
	}
	return ret
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

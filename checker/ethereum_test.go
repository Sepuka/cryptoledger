package checker

import "testing"

type dataSet struct {
	value          string
	expectedResult int64
}

var dataProvider = []dataSet{
	{"96523353125837500", 0},
	{"14960423015270544500", 14},
}

func TestWeiToEth(t *testing.T) {
	for _, set := range dataProvider {
		actualResult := WeiToEth(set.value)
		if actualResult != set.expectedResult {
			t.Error("For", set.value, "expected", set.expectedResult, "but got", actualResult)
		}
	}
}

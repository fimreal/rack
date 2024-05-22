package money_test

import (
	"fmt"
	"testing"

	"github.com/fimreal/rack/module/money"
)

func TestCountInterest(t *testing.T) {
	f := money.CountInterest
	本金1 := 10000.0
	利率1 := f(本金1, 22.00*12, 12)
	fmt.Printf("%.3f%%\n", 利率1*100)

	本金2 := 3000.0
	利率2 := f(本金2, 6.6*12-30, 12)
	fmt.Printf("%.3f%%\n", 利率2*100)
}

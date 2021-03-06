package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/2018/q22"
)

func main() {
	m := q22.NewMap(q22.C(9, 796), 6969)
	riskLevel := m.RiskLevel()
	fmt.Printf("riskLevel = %+v\n", riskLevel)
}

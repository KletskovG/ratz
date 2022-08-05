package main

import (
	"fmt"
	"github.com/kletskovg/ratz/pkg/util"
)

func main() {
	fmt.Println("First run of ratz app")
	util.Rates("BTC", "USD")
}

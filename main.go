package main

import (
	"fmt"
	co "github.com/MehrbanooEbrahimzade/FootballTeams/co"
	nco "github.com/MehrbanooEbrahimzade/FootballTeams/nco"
	"time"
)

func main() {
	var Concurrency = time.Now()
	co.Concurrency()
	fmt.Println(time.Since(Concurrency))

	var nonConcurrency = time.Now()
	nco.NonConcurrency()
	fmt.Println(time.Since(nonConcurrency))

}

package main

import (
	"flag"

	"github.com/gidoBOSSftw5731/log"
)

var (
	interestRate = flag.Float64("rate", 4,
		`Interest rate in percent, defaults to 4%`)

	dollarsInvested = flag.Float64("invested", 5000,
		`Amount invested during each cycle, defaults to $5000`)

	limit = flag.Float64("limit", 0,
		`At what point of investment do we stop calculating and return (default 0 means never)`)

	tBtwnCmpnd = flag.Float64("tbc", 4,
		`Multiplier of time between payments against compounding. Defaults to 1 investment per year
and 4 compounds throughout the year`)

	initInvest = flag.Float64("init", 0,
		`Initial Investment, defaults to 0`)

	loglevel = flag.Int("loglevel", 2,
		`1 for error
	2 for info
	3 for debug
	4 for trace
	each level includes all lower levels`)
)

func main() {
	flag.Parse()

	//1 for error
	//2 for info
	//3 for debug
	//4 for trace
	// each level includes all lower levels
	log.SetCallDepth(*loglevel)

	var total, cycles float64
	var ctr int8

	total = *initInvest

	// Runs once per deposit cycle
	for {
		cycles += *tBtwnCmpnd

		var i int
		// if you, say compound annaully and invest quarterly, this would only run at the
		// end of each year, by incrementing the cycles var until it is >= 1
		for cycles >= 1 {
			i++
			// divide by 100 because percentages
			total += total * (*interestRate / 100)

			//check if limit has been reached
			if total >= *limit && *limit != 0 {
				log.Fatalf("At limit, exiting with interest of %v after %v iterations and %v extra compoundings",
					total, ctr, i)
			}
			cycles--
		}
		log.Infof("Finished compounding interest, balance is now %v after %v iterations",
			total, ctr)

		log.Infof("Adding %v to total", *dollarsInvested)
		total += *dollarsInvested

		// now one unit of time is passing
		ctr++
	}

}

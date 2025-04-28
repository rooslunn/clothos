package main

import (
	"fmt"
	"log/slog"
	"math/big"
	"math/rand"
	"os"
	"time"
)

const (
	WitnessesCount = 23
	DateFormat     = "2006-01-02 15:04:05"
)

func main() {

	log := setupLogger()


	var now time.Time
	var unixTimestamp int64
	var numberToCheck *big.Int
	var isPrime bool
	var sleepSeconds int

	now = time.Now()
	year, month, day := now.Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Add(24 * time.Hour)

	log.Info("Game started", "startDate", startDate.String())

	for {

		unixTimestamp = startDate.Unix()
		numberToCheck = big.NewInt(unixTimestamp)

		isPrime = MillerRabin(numberToCheck, WitnessesCount)

		if isPrime {
			msg := fmt.Sprintf("%s is likely the PRIME (%s)", startDate.Format(DateFormat), numberToCheck.String())
			log.Info(msg)
		}

		sleepSeconds = randInt(4, 240)
		// log.Info(fmt.Sprintf("sleeping for %d secs...", sleepSeconds))

		startDate = startDate.Add(time.Duration(sleepSeconds) * time.Second)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

func randInt(min, max int) int {
	if min > max {
		panic("min cannot be greater than max")
	}
	if min == max {
		return min
	}
	return min + rand.Intn(max-min+1)
}

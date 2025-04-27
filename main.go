package main

import (
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"time"
)

const (
	WitnessesCount = 23
	TimeDeltaMinute = 1
	SleepSeconds    = 3
	DateFormat = "2006-01-02 15:04:05"
)

func main() {

	log := setupLogger()

	now := time.Now()
	year, month, day := now.Date()
	startOfDayLocal := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	currentTime := startOfDayLocal

	log.Info(fmt.Sprintf("Game started from %s", currentTime.Format(DateFormat)))

	for {

		unixTimestamp := currentTime.Unix()
		numberToCheck := big.NewInt(unixTimestamp)

		isPrime := MillerRabin(numberToCheck, WitnessesCount)
		if isPrime {
			msg := fmt.Sprintf("%s is likely the PRIME (%s)", currentTime.Format(DateFormat), numberToCheck.String())
			log.Info(msg)
		}

		incTime := TimeDeltaMinute * time.Minute 
		currentTime = currentTime.Add(incTime)

		// log.Info(fmt.Sprintf("continue after %d secs", SleepSeconds))
		time.Sleep(SleepSeconds)
	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

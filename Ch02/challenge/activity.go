package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Activity struct {
	User        string    `json:"user"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func (act *Activity) validate(r io.Reader) error {
	if len(act.User) == 0 {
		return fmt.Errorf("missing user")
	}

	if len(act.Description) == 0 {
		return fmt.Errorf("missing Description")
	}

	startTime := act.StartTime
	endTime := act.EndTime

	if startTime.After(endTime) {
		return fmt.Errorf("Start time = %s \n should be previous to the end time = %s\n",
			startTime, endTime)
	}
	return nil
}

func processActivity(r io.Reader) error {
	var act Activity

	const (
		maxSize = 100 * 1024
	)
	dec := json.NewDecoder(io.LimitReader(r, maxSize))
	if err := dec.Decode(&act); err != nil {
		return err
	}

	if err := act.validate(r); err != nil {
		log.Printf("err: processActivity - %s", err)
		// http.Error(r, "bad data", http.StatusBadRequest)
		return err
	}

	log.Printf("activity: %#v", act)
	// TODO: Store in database

	return nil
}

func main() {
	if err := processActivity(os.Stdin); err != nil {
		log.Fatal(err)
	}
}

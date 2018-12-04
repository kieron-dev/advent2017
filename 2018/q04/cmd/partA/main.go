package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Timestamp time.Time
	Message   string
}

func (e *Entry) Load(line string) {
	halves := strings.Split(line, "] ")
	e.Message = halves[1]
	timestamp := halves[0][1:]
	year, _ := strconv.Atoi(timestamp[:4])
	month, _ := strconv.Atoi(timestamp[5:7])
	day, _ := strconv.Atoi(timestamp[8:10])
	hour, _ := strconv.Atoi(timestamp[11:13])
	minute, _ := strconv.Atoi(timestamp[14:16])
	e.Timestamp = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
}

type GuardSession struct {
	Id        int
	Sleeps    map[int]int
	AsleepFor int
}

func NewGuardSession(id int) *GuardSession {
	sess := GuardSession{Id: id}
	sess.Sleeps = map[int]int{}
	return &sess
}

func (s *GuardSession) Sleeping(from, to int) {
	for i := from; i < to; i++ {
		s.Sleeps[i]++
	}
	s.AsleepFor += to - from
}

func (s *GuardSession) MaxMinute() int {
	var maxVal, maxMinute int
	for i := 0; i < 60; i++ {
		if s.Sleeps[i] > maxVal {
			maxVal = s.Sleeps[i]
			maxMinute = i
		}
	}
	return maxMinute
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var entries []Entry

	for scanner.Scan() {
		line := scanner.Text()
		entry := Entry{}
		entry.Load(line)
		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.Unix()-entries[j].Timestamp.Unix() < 0
	})

	sessions := map[int]*GuardSession{}
	var currentId int
	lastSleepStart := 0

	for _, entry := range entries {
		if strings.HasPrefix(entry.Message, "Guard") {
			fmt.Sscanf(entry.Message, "Guard #%d", &currentId)
			var session *GuardSession
			var ok bool
			session, ok = sessions[currentId]
			if !ok {
				session = NewGuardSession(currentId)
				sessions[currentId] = session
			}
			lastSleepStart = 0
		} else if strings.Contains(entry.Message, "falls") {
			lastSleepStart = entry.Timestamp.Minute()
		} else if strings.Contains(entry.Message, "wakes") {
			sessions[currentId].Sleeping(lastSleepStart, entry.Timestamp.Minute())
		}
	}

	maxSleep := 0
	maxId := 0
	for id, s := range sessions {
		if s.AsleepFor > maxSleep {
			maxSleep = s.AsleepFor
			maxId = id
		}
	}
	bestMinute := sessions[maxId].MaxMinute()
	fmt.Printf("maxId = %+v\n", maxId)
	fmt.Printf("bestMinute = %+v\n", bestMinute)
	fmt.Printf("maxId * bestMinute = %+v\n", maxId*bestMinute)

	fmt.Println("***********")

	maxVal := 0
	maxMinute := 0
	for id, s := range sessions {
		bestMinute = s.MaxMinute()
		val := s.Sleeps[bestMinute]
		if val > maxVal {
			maxVal = val
			maxId = id
			maxMinute = bestMinute
		}
	}
	fmt.Printf("maxId = %+v\n", maxId)
	fmt.Printf("maxMinute = %+v\n", maxMinute)
	fmt.Printf("maxId * maxMinute = %+v\n", maxId*maxMinute)
}

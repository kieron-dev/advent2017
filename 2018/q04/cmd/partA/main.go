package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Year    int
	Month   int
	Day     int
	Hour    int
	Minute  int
	Message string
}

func (e *Entry) Load(line string) {
	halves := strings.Split(line, "] ")
	e.Message = halves[1]
	timestamp := halves[0][1:]
	e.Year, _ = strconv.Atoi(timestamp[:4])
	e.Month, _ = strconv.Atoi(timestamp[5:7])
	e.Day, _ = strconv.Atoi(timestamp[8:10])
	e.Hour, _ = strconv.Atoi(timestamp[11:13])
	e.Minute, _ = strconv.Atoi(timestamp[14:16])
}

type GuardSession struct {
	Id    int
	Year  int
	Month int
	Day   int
	Awake []bool
}

func NewGuardSession(id, year, month, day int) *GuardSession {
	sess := GuardSession{Id: id, Year: year, Month: month, Day: day}
	sess.Awake = []bool{}
	return &sess
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
		if entries[i].Year != entries[j].Year {
			return entries[i].Year < entries[j].Year
		}
		if entries[i].Month != entries[j].Month {
			return entries[i].Month < entries[j].Month
		}
		if entries[i].Day != entries[j].Day {
			return entries[i].Day < entries[j].Day
		}
		if entries[i].Hour != entries[j].Hour {
			return entries[i].Hour < entries[j].Hour
		}
		if entries[i].Minute != entries[j].Minute {
			return entries[i].Minute < entries[j].Minute
		}
		return false
	})

	sessions := []*GuardSession{}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Message, "Guard") {
			var id, day int
			fmt.Sscanf(entry.Message, "Guard #%d", &id)
			if entry.Hour == 23 {
				day = entry.Day + 1
			} else {
				day = entry.Day
			}
			session := NewGuardSession(id, entry.Year, entry.Month, day)
			sessions = append(sessions, session)
		}
	}
	for _, s := range sessions {
		fmt.Printf("*s= %+v\n", *s)
	}
}

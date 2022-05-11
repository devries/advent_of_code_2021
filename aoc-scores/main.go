package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StarTs int64

func (v *StarTs) UnmarshalJSON(data []byte) error {
	val := string(data)

	res, err := strconv.ParseInt(strings.Trim(val, `"'`), 10, 64)
	if err != nil {
		return fmt.Errorf("Unable to decode %s to int64: %s", val, err)
	}

	*v = StarTs(res)

	return nil
}

type Completion struct {
	GetStarTs StarTs `json:"get_star_ts"`
}

type Member struct {
	LocalScore         int                        `json:"local_score"`
	CompletionDayLevel map[int]map[int]Completion `json:"completion_day_level"`
	Stars              int                        `json:"stars"`
	GlobalScore        int                        `json:"global_score"`
	Id                 string                     `json:"id"`
	Name               string                     `json:"name"`
	LastStarTs         StarTs                     `json:"last_star_ts"`
}

type Scoreboard struct {
	Members map[int]Member `json:"members"`
	Event   string         `json:"event"`
	OwnerId string         `json:"owner_id"`
}

type UserSortable struct {
	Name   string
	Stars  int
	Finish int64
}

type ByScore []UserSortable

func (a ByScore) Len() int      { return len(a) }
func (a ByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool {
	if a[i].Stars != a[j].Stars {
		return a[i].Stars > a[j].Stars
	}
	return a[i].Finish < a[j].Finish
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <scorefile>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	check(err)

	decoder := json.NewDecoder(f)

	var s Scoreboard

	err = decoder.Decode(&s)
	check(err)

	// Print out daily results for each member
	memberNumbers := []int{}
	maxNameLength := 0
	for k := range s.Members {
		memberNumbers = append(memberNumbers, k)
		nameLength := len(s.Members[k].Name)
		if nameLength > maxNameLength {
			maxNameLength = nameLength
		}
	}

	eventYear, err := strconv.ParseInt(s.Event, 10, 64)
	check(err)

	now := time.Now()

	for i := 1; i < 26; i++ {
		dayStart := time.Date(int(eventYear), 12, i, 5, 0, 0, 0, time.UTC)
		if dayStart.After(now) {
			break
		}
		for _, j := range []int{1, 2} {
			users := []UserSortable{}
			for _, n := range memberNumbers {
				completions, ok := s.Members[n].CompletionDayLevel[i]
				if !ok {
					continue
				}
				completion, ok := completions[j]
				if !ok {
					continue
				}

				users = append(users, UserSortable{s.Members[n].Name, 1, int64(completion.GetStarTs)})
			}
			sort.Sort(ByScore(users))
			if len(users) > 0 {
				fmt.Printf("Day %2d part %d:\n", i, j)
			}
			for _, u := range users {
				doneAt := time.Unix(u.Finish, 0)
				dur := doneAt.Sub(dayStart)

				fmt.Printf("%*s: %s\n", maxNameLength+4, u.Name, fmtDuration(dur))
			}
			if len(users) > 0 {
				fmt.Printf("\n")
			}
		}
		// Write time from p1 to p2
		users := []UserSortable{}
		for _, n := range memberNumbers {
			if completions, ok := s.Members[n].CompletionDayLevel[i]; ok {
				if completion2, ok := completions[2]; ok {
					completion1 := completions[1]

					ts2 := int64(completion2.GetStarTs)
					ts1 := int64(completion1.GetStarTs)
					users = append(users, UserSortable{s.Members[n].Name, 1, ts2 - ts1})
				}
			}
		}
		if len(users) > 0 {
			fmt.Printf("Day %2d time between parts:\n", i)
		}
		sort.Sort(ByScore(users))
		for _, u := range users {
			d := time.Duration(u.Finish) * time.Second
			fmt.Printf("%*s: %s\n", maxNameLength+4, u.Name, fmtDuration(d))
		}
		if len(users) > 0 {
			fmt.Printf("\n")
			fmt.Printf("\n")
		}
	}

	users := []UserSortable{}
	for _, n := range memberNumbers {
		ts := int64(s.Members[n].LastStarTs)
		if ts == 0 {
			users = append(users, UserSortable{s.Members[n].Name, 0, 0})
		} else {
			users = append(users, UserSortable{s.Members[n].Name, s.Members[n].Stars, ts})
		}
	}
	sort.Sort(ByScore(users))
	for _, u := range users {
		if u.Stars == 0 {
			fmt.Printf("%*s did not complete any stars\n", maxNameLength, u.Name)
		} else {
			finished := time.Unix(u.Finish, 0)
			fmt.Printf("%*s finished %d stars on %s\n", maxNameLength, u.Name, u.Stars, finished.Format("January 2, 2006 at 03:04 PM"))
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	if h > 0 {
		return fmt.Sprintf("%5d h %2d m", h, m)
	} else {
		return fmt.Sprintf("        %2d m", m)
	}
}

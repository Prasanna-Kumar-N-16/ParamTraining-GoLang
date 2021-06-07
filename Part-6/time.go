package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	now := time.Now()
	p(now)
	def := time.Date(1999, 02, 12, 4, 20, 60, 651387237, time.UTC)
	p(def)
	p(def.Date())
	p(def.YearDay())
	p(def.Second())
	p(def.Nanosecond())
	p(def.Weekday())
	p(def.Before(now))
	p(def.Sub(now))
	//Unix epoch.
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	//time formatting
	p(now.Format(time.RFC3339))
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p(e)
	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
}

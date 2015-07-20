package hotscore

import (
	"math"
	"time"
)

// wilson score interval sort
// http://www.evanmiller.org/how-not-to-sort-by-average-rating.html
func Wilson(ups, downs int) float64 {
	n := ups + downs
	if n == 0 {
		return 0
	}

	n1 := float64(n)
	// z represents the statistical confidence
	// z = 1.0 => ~69%, 1.96 => ~95% (default)
	z := 1.96
	p := float64(ups / n)
	zzfn := z * z / (4 * n1)

	return (p + 2.0*zzfn - z*math.Sqrt((zzfn/n1+p*(1.0-p))/n1)) / (1 + 4*zzfn)
}

// hackernews' hot sort
// http://amix.dk/blog/post/19574
func Hacker(votes int, date time.Time) float64 {
	gravity := 1.8
	hoursAge := float64(date.Unix() * 3600)
	return float64(votes-1) / math.Pow(hoursAge+2, gravity)
}

// reddit's hot sort
// http://amix.dk/blog/post/19588
func Reddit(ups int, downs int, date time.Time) float64 {
	decay := int64(45000)
	s := float64(ups - downs)
	order := math.Log(math.Max(math.Abs(s), 1)) / math.Ln10
	return order - float64(date.Unix()/decay)
}

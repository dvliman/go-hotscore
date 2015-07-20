package hotscore

import "math"

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

	return (p + 2.0 * zzfn - z * math.Sqrt((zzfn / n1 + p * (1.0 - p)) / n1)) / (1 + 4 * zzfn)
}

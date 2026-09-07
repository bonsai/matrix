package match

import (
	"math"
	"github.com/bonsai/matrix/vector"
)

// Similarity returns cosine similarity for sparse vectors.
func Similarity(a, b vector.Sparse) float64 {
	na, nb := vector.Norm(a), vector.Norm(b)
	if na == 0 || nb == 0 { return 0 }
	return vector.Dot(a, b) / (na * nb)
}

// Complementarity measures non-overlapping positive preference mass.
func Complementarity(a, b vector.Sparse) float64 {
	var sum, count float64
	seen := make(map[string]struct{}, len(a)+len(b))
	for k := range a { seen[k] = struct{}{} }
	for k := range b { seen[k] = struct{}{} }
	for k := range seen {
		x, y := a[k], b[k]
		sum += math.Abs(x-y)
		if x != 0 || y != 0 { count++ }
	}
	if count == 0 { return 0 }
	return sum / count
}

// SharedValues measures the amount of positive preference shared by both vectors.
func SharedValues(a, b vector.Sparse) float64 {
	var sum float64
	for k, x := range a {
		if y := b[k]; x > 0 && y > 0 { sum += math.Min(x, y) }
	}
	return sum
}

// Weights controls the three match dimensions.
type Weights struct { Similarity, Complementarity, SharedValues float64 }

// Score computes the weighted match score.
func Score(s, c, v float64, w Weights) float64 {
	return w.Similarity*s + w.Complementarity*c + w.SharedValues*v
}

// Result contains the explainable components of a match.
type Result struct {
	Similarity float64 `json:"similarity"`
	Complementarity float64 `json:"complementarity"`
	SharedValues float64 `json:"shared_values"`
	Score float64 `json:"score"`
}

// Match computes all match dimensions in one pure operation.
func Match(a, b vector.Sparse, w Weights) Result {
	s, c, v := Similarity(a,b), Complementarity(a,b), SharedValues(a,b)
	return Result{s, c, v, Score(s,c,v,w)}
}

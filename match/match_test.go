package match

import (
	"math"
	"testing"
	"github.com/bonsai/matrix/vector"
)

func TestSimilarity(t *testing.T) {
	a := vector.Sparse{"jazz": 1, "vinyl": 1}
	if got := Similarity(a, a); math.Abs(got-1) > 1e-9 { t.Fatalf("got %v", got) }
}

func TestMatch(t *testing.T) {
	a := vector.Sparse{"jazz": 1, "vinyl": 1}
	b := vector.Sparse{"jazz": 1, "sake": 1}
	w := Weights{Similarity: .4, Complementarity: .3, SharedValues: .3}
	r := Match(a, b, w)
	if r.Score <= 0 { t.Fatalf("expected positive score: %+v", r) }
}

package vector

// Sparse is a sparse vector keyed by ontology/tag ID.
type Sparse map[string]float64

// Normalize returns a copy scaled so the largest absolute value is 1.
func (v Sparse) Normalize() Sparse {
	var max float64
	for _, x := range v {
		if x < 0 {
			x = -x
		}
		if x > max {
			max = x
		}
	}
	out := make(Sparse, len(v))
	if max == 0 {
		for k, x := range v { out[k] = x }
		return out
	}
	for k, x := range v { out[k] = x / max }
	return out
}

// Dot computes the sparse dot product.
func Dot(a, b Sparse) float64 {
	if len(a) > len(b) { a, b = b, a }
	var sum float64
	for k, x := range a { sum += x * b[k] }
	return sum
}

// Norm returns the Euclidean norm.
func Norm(v Sparse) float64 {
	return Dot(v, v) ** 0.5
}

package structs

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Perimeter(w, h float64) float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Area(w, h float64) float64 {
	return r.width * r.height
}

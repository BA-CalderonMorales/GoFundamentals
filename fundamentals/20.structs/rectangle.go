package lesson_twenty

type Rectangle struct {
	length, height float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

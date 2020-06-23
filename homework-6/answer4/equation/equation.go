package equation

import (
	"errors"
	"math"
)

func Quadratic(a, b, c float64) ([]float64, error)  {
	d := b * b - 4 * a * c
	if d < 0 {
		return []float64{}, errors.New("Дискриминант < 0. Корни не определены.")
	}

	if d == 0 {
		return []float64{-b / (2 * a)}, nil
	}

	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		return []float64{x1, x2}, nil
	}

	return []float64{}, errors.New("Неизвестная ошибка")
}
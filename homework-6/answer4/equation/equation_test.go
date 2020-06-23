package equation_test

import (
	"GoCourse/homework-6/answer4/equation"
	"errors"
	"testing"
)

type testQuadraticData struct {
	a, b, c float64
	expected []float64
	error
}
func TestQuadratic(t *testing.T) {
	testSet := []testQuadraticData{
		{
			a: 1,
			b: 2,
			c: 3,
			expected: []float64{},
			error:          errors.New("Дискриминант < 0. Корни не определены."),
		},
		{
			a: 1,
			b: -2,
			c: -3,
			expected: []float64{3, -1},
			error:          nil,
		},
		{
			a: -1,
			b: -2,
			c: 15,
			expected: []float64{-5, 3},
			error:          nil,
		},
		{
			a: 1,
			b: 12,
			c: 36,
			expected: []float64{-6},
			error:          nil,
		},
	}

	for i, data := range testSet {
		res, err := equation.Quadratic(data.a, data.b, data.c)

		if nil != err && data.error == nil {
			t.Error(
				"Неожиданная ошибка",
				"a: ", data.a,
				"b: ", data.b,
				"c: ", data.c,
				"error: ", err,
				"expected: ", data.expected,
				"data:", i,
			)
		}

		if nil != err && data.error != nil && data.error.Error() != err.Error() {
			t.Error(
				"Неверная ошибка",
				"a: ", data.a,
				"b: ", data.b,
				"c: ", data.c,
				"error: ", err,
				"expected: ", data.expected,
				"data:", i,
				)
		}

		if len(data.expected) != len(res) {
			t.Error(
				"Неправильный результат",
				"a: ", data.a,
				"b: ", data.b,
				"c: ", data.c,
				"error: ", err,
				"expected: ", data.expected,
				"data:", i,
			)
		}

		for exceptedIndex, expectedValue := range data.expected {
			if res[exceptedIndex] != expectedValue {
				t.Error(
					"Корни не совпадают",
					"a: ", data.a,
					"b: ", data.b,
					"c: ", data.c,
					"error: ", err,
					"expected: ", data.expected,
					"excepted value: ", expectedValue,
					"res value: ", res[exceptedIndex],
					"data:", i,
				)
			}
		}
	}
}
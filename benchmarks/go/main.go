// SPDX-License-Identifier: BSD-3-Clause
package main

import "github.com/go-ruby-matrix/matrix"

func build(n int, f func(i, j int) float64) *matrix.Matrix {
	rows := make([][]any, n)
	for i := range rows {
		rows[i] = make([]any, n)
		for j := range rows[i] {
			rows[i][j] = f(i, j)
		}
	}
	m, _ := matrix.New(rows)
	return m
}

func main() {
	n := 24
	a := build(n, func(i, j int) float64 { return float64((i*7+j*3)%13)*0.5 + 0.25 })
	b := build(n, func(i, j int) float64 { return float64((i*5+j*2)%11)*0.25 + 0.5 })
	bench("mul-24x24", 200, func() { v, _ := a.Mul(b); sink = v })
	bench("transpose-24x24", 2000, func() { sink = a.Transpose() })
}

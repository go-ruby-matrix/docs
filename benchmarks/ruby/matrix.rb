# frozen_string_literal: true
# SPDX-License-Identifier: BSD-3-Clause
require "matrix"
require_relative "_harness"
n = 24
rows_a = Array.new(n) { |i| Array.new(n) { |j| (((i * 7 + j * 3) % 13) * 0.5 + 0.25) } }
rows_b = Array.new(n) { |i| Array.new(n) { |j| (((i * 5 + j * 2) % 11) * 0.25 + 0.5) } }
a = Matrix[*rows_a]
b = Matrix[*rows_b]
bench("mul-24x24", 200) { a * b }
bench("transpose-24x24", 2000) { a.transpose }

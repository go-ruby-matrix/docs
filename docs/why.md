# Why pure Go

`go-ruby-matrix/matrix` reimplements Ruby's `matrix` standard library in **pure
Go, with cgo disabled**. `Matrix`/`Vector` arithmetic is **deterministic and
interpreter-independent**: given its entries, the result is a pure function of
those entries — no live binding, no evaluation of arbitrary Ruby. That is exactly
the part that can — and should — live as a standalone Go library, separate from
the interpreter.

## A reusable library, bound by rbgo

This library is the deterministic core that backs Ruby's `matrix` in
[go-embedded-ruby](https://github.com/go-embedded-ruby/ruby). It is **standalone
and reusable** so that:

- any Go program can import `github.com/go-ruby-matrix/matrix` directly, with no
  Ruby runtime;
- the dependency runs the *other* way — `rbgo` binds this module as a native
  module (the same pattern as [go-ruby-regexp](https://github.com/go-ruby-regexp)
  and [go-ruby-erb](https://github.com/go-ruby-erb)), rather than this module
  depending on the interpreter;
- the behaviour is pinned by a **differential oracle** against the system
  `ruby -rmatrix`, independent of any one consumer.

## Exact where MRI is exact

Ruby's `matrix` carries Ruby's numeric tower: Integer and Rational entries stay
exact, and a division only escapes to Float when the input is Float. This package
mirrors that with a `Num` tower built on `math/big` — `*big.Int` for Integers,
`*big.Rat` for Rationals, `float64` for Floats — and reproduces MRI's
`inverse_from` Gauss-Jordan elimination **step for step** (partial pivoting on
the largest absolute pivot), so even a Float matrix's inverse reproduces MRI's
exact rounding (e.g. `-1.9999999999999998`). An Integer matrix's `determinant`
stays an Integer; `inverse` of an Integer matrix prints exact Rationals
(`(1/1)`, not `1`). The arithmetic is a pure function of the entries — no
interpreter required.

## Why pure Go matters here

Because the library is CGO-free and dependency-free, it:

- cross-compiles to every Go target with no C toolchain, and links into a single
  static binary;
- has **no dependency on the Ruby runtime** — the dependency runs the other way;
- can be differentially tested against `ruby -rmatrix` wherever a `ruby` is on
  `PATH`, while the cross-arch lanes (where `ruby` is absent) still validate the
  library itself.

See [Usage & API](api.md) for the surface and [Roadmap](roadmap.md) for what is
in scope.

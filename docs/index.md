# go-ruby-matrix documentation

**Ruby's `matrix` stdlib in pure Go — `Matrix` & `Vector` with exact arithmetic, MRI-compatible, no cgo.**

`go-ruby-matrix/matrix` is a faithful, pure-Go (zero cgo) reimplementation of
Ruby's `matrix` standard library — `Matrix` and `Vector` with MRI 4.0.5-faithful
behaviour — matching reference Ruby byte-for-byte. The module path is
`github.com/go-ruby-matrix/matrix`.

It matches MRI's `to_s` / `inspect` formatting, its error classes, and —
crucially — its **exact arithmetic**: an Integer matrix's `determinant` stays an
Integer, and `inverse` produces exact Rationals
(`Matrix[[1,2],[3,4]].inverse` → `Matrix[[(-2/1), (1/1)], [(3/2), (-1/2)]]`) —
**without any Ruby runtime**. It is the `matrix` backend bound into
[go-embedded-ruby](https://github.com/go-embedded-ruby/ruby) by `rbgo` as a
native module, a standalone sibling of
[go-ruby-regexp](https://github.com/go-ruby-regexp),
[go-ruby-erb](https://github.com/go-ruby-erb) and
[go-ruby-yaml](https://github.com/go-ruby-yaml). The dependency runs the other
way: this library has **no dependency on the Ruby runtime**.

!!! success "Status: complete — MRI byte-exact"
    A faithful port of the `Matrix` and `Vector` API — constructors, exact operations (determinant / inverse / rank over the numeric tower), predicates, equality & formatting, the full `Vector` surface, and MRI's error classes. Validated by a **differential oracle** against the system `ruby -rmatrix`, at 100% coverage, `gofmt` + `go vet` clean, CI green across the six 64-bit Go targets and three OSes.

## Quick taste

```go
m, _ := matrix.New([][]any{{1, 2}, {3, 4}})

d, _ := m.Determinant()
fmt.Println(d)         // -2          (Integer — exact)

inv, _ := m.Inverse()
fmt.Println(inv.ToS()) // Matrix[[(-2/1), (1/1)], [(3/2), (-1/2)]]  (exact Rationals)
```

## Repositories

| Repo | What it is |
| --- | --- |
| [`matrix`](https://github.com/go-ruby-matrix/matrix) | the library — Ruby's `matrix` stdlib in pure Go |
| [`docs`](https://github.com/go-ruby-matrix/docs) | this documentation site (MkDocs Material, versioned with mike) |
| [`go-ruby-matrix.github.io`](https://github.com/go-ruby-matrix/go-ruby-matrix.github.io) | the organization landing page (Hugo) |
| [`brand`](https://github.com/go-ruby-matrix/brand) | logo and brand assets |

## Principles

- **Pure Go, `CGO_ENABLED=0`** — trivial cross-compilation, a single static
  binary, no C toolchain.
- **Exact where MRI is exact.** A `Num` tower on `math/big` keeps Integer and
  Rational entries exact, with Ruby's promotion rules.
- **MRI byte-exact.** Every operation's `ToS` matches reference Ruby's `inspect`,
  validated by a differential oracle against `ruby -rmatrix`.
- **Standalone & reusable.** Extracted from rbgo's internals; no dependency on
  the Ruby runtime — the dependency runs the other way.
- **100% test coverage** is the target, enforced as a CI gate, across 6 arches
  and 3 OSes.

## Where to go next

- [Why pure Go](why.md) — why `Matrix`/`Vector` arithmetic is deterministic
  enough to live as a standalone, interpreter-independent Go library.
- [Usage & API](api.md) — the numeric tower, the full `Matrix`/`Vector` surface,
  and worked examples.
- [Roadmap](roadmap.md) — what is done and what is downstream by design.

Source lives at [github.com/go-ruby-matrix/matrix](https://github.com/go-ruby-matrix/matrix).

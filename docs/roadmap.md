# Roadmap

`go-ruby-matrix/matrix` is grown **test-first**, each capability
differential-tested against MRI rather than built in isolation. Ruby's `matrix`
standard library — the deterministic, interpreter-independent slice — is
**complete**.

| Stage | What | Status |
| --- | --- | --- |
| The `Num` numeric tower | Entries flow through `Num` (`*big.Int` Integers, `*big.Rat` Rationals, `float64` Floats) with Ruby's promotion: Integer op Integer stays Integer, any Rational stays Rational even when whole, any Float dominates. | **Done** |
| Constructors & accessors | `New` / `Build` / `Identity` / `Zero` / `Diagonal` / `Scalar` / `RowVector` / `ColumnVector` / `Rows` / `Columns` / `HStack` / `VStack`; `RowCount` / `ColumnCount` / `At` / `Row` / `Column` / `Each` / `Minor` / `FirstMinor` / `ToA`. | **Done** |
| Exact operations | `Add` / `Sub` / `Neg` / `Mul` (matrix·matrix, ·scalar, ·vector) / `Div` / `Pow` / `Transpose` / `Trace` / `Determinant` / `Inverse` / `Rank` — exact over the tower, reproducing MRI's `inverse_from` Gauss-Jordan step for step. | **Done** |
| Predicates & formatting | `Square` / `IsDiagonal` / `Symmetric` / `Orthogonal` / `Singular` / `Regular` / `LowerTriangular` / `UpperTriangular` / `IsZero`; `Eql` (`==`), `Hash`; `ToS` / `Inspect` (`Matrix[[…], […]]`, `Matrix.empty(r, c)`). | **Done** |
| Vector | `Elements` / `At` / `Size` / `Add` / `Sub` / `Mul` / `InnerProduct` (dot) / `CrossProduct` / `Magnitude` (norm) / `Normalize` / `Each` / `Map` / `Angle` / `Eql` / `Independent`. | **Done** |
| Differential oracle & coverage | Deterministic ruby-free tests hold coverage at 100%; a differential oracle compares every operation's `ToS` against `ruby -rmatrix … .inspect`, including the exact-arithmetic results and the Ruby-faithful float formatter; gofmt + go vet clean, green across all six 64-bit Go arches and three OSes. | **Done** |

## Documented out-of-scope boundaries

These are **deliberate**, recorded so the module's surface is unambiguous:

- **No interpreter.** The library implements the deterministic arithmetic; it
  never runs arbitrary Ruby. Mapping a host's live numerics to and from the `Num`
  tower is the consumer's job — that is why `rbgo` binds this module rather than
  the reverse.
- **Reference is reference Ruby (MRI).** Byte-for-byte conformance targets MRI
  4.0.5's `matrix`; the differential oracle gates on `RUBY_VERSION >= "4.0"`.
- **Standalone & reusable.** The module has no dependency on the Ruby runtime;
  the dependency runs the other way.

See [Usage & API](api.md) for the surface and [Why pure Go](why.md) for the
deterministic/interpreter split.

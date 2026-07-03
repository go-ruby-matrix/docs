# Performance

`go-ruby-matrix/matrix` is the pure-Go library that
[`rbgo`](https://github.com/go-embedded-ruby/ruby) binds for Ruby's `matrix`.
This page records the **methodology** of the ecosystem-wide per-module parity
benchmark; it does not quote numbers that have not been measured on this module.

## What is measured

The **same** Ruby script — a `Matrix`/`Vector` workload (construction,
multiplication, `determinant`, `inverse`, `rank` over the numeric tower) on a
representative set of matrices — is run under every runtime. `rbgo`'s number
reflects **this pure-Go library doing the work**; every other column is that
interpreter's own `matrix` stdlib. So the comparison is the **Ruby-visible
operation**, apples-to-apples across interpreters. The script prints a
deterministic checksum and its output is checked **byte-identical to MRI** before
timing.

## Method

- **Host:** a single fixed machine; **best-of-N wall time** (best, not mean, to
  suppress scheduler noise); single-shot processes, no warm-up beyond the
  script's own loop.
- **Runtimes:** `ruby` (MRI, the oracle) and `ruby --yjit`; `jruby`;
  `truffleruby` — each running its own `matrix`, against `rbgo` running this
  library.
- The benchmark script and harness live in rbgo's repo under
  [`bench/modules/`](https://github.com/go-embedded-ruby/ruby/tree/main/bench/modules).
  Reproduce with the per-module runner there.

## Result (best of 5, ms)

| Runtime | time | vs MRI |
| --- | ---: | ---: |
| **rbgo** (go-ruby-matrix) | 20 | 0.40× |
| MRI (ruby 4.0.5) | 50 | 1.00× |
| MRI + YJIT | 50 | 1.00× |
| JRuby 10.1.0.0 | 1230 | 24.60× |
| TruffleRuby 34.0.1 | 210 | 4.20× |

rbgo runs on **go-ruby-matrix** and is **faster than MRI** here (0.40x): MRI's `Matrix` multiply/determinant is Ruby-coded, so the compiled pure-Go library wins. At ~20 ms the row is near the noise floor.

!!! note "Honest framing"
    JRuby and TruffleRuby are timed **cold, single-shot**, so they carry JVM /
    Graal startup on every run — read them as one-shot `ruby file.rb` costs, the
    same way `rbgo` and MRI are measured, not as steady-state JIT numbers. Rows
    that complete in well under ~200 ms carry the most relative noise; treat
    their ratios as order-of-magnitude. These are **real measured numbers** from
    the 2026-06-30 run (Apple M-series; `ruby 4.0.5 +PRISM`, `jruby 10.1.0.0`,
    `truffleruby 34.0.1`) — nothing is fabricated or cherry-picked.

## Library-level benchmark (Go API vs runtimes) — 2026-07-03

This section measures the **pure-Go library directly, through its Go API** — not
the `rbgo` interpreter path recorded above. It isolates the library primitive
from Ruby-interpreter dispatch, answering the parity question head-on: *is the
pure-Go implementation as fast as the reference runtime's own `matrix`?* The
**same workload, same inputs, same iteration counts** run through the Go library
and through each reference runtime's stdlib; outputs were checked identical to
MRI before any timing.

- **Host:** Apple M4 Max (`Mac16,5`, arm64), macOS 26.5.1 — **date 2026-07-03**.
- **Runtimes:** Go 1.26.4 · MRI `ruby 4.0.5 +PRISM` · MRI + YJIT · JRuby 10.1.0.0
  (OpenJDK 25) · TruffleRuby 34.0.1 (GraalVM CE Native).
- **Method:** each process runs 3 untimed warm-up passes, then 25 timed passes of
  a fixed inner loop, timed with a monotonic clock; the **best** pass is reported
  as **ns/op** (lower is better). `vs MRI` < 1.00× means *faster than MRI*.
  Interpreter start-up is outside the timed region, so these are operation costs,
  not `ruby file.rb` process costs.

#### mul-24x24

| Runtime | ns/op | vs MRI |
| --- | ---: | ---: |
| **go-ruby (pure Go)** | 116574.6 | 0.20× |
| MRI | 580445.0 | 1.00× |
| MRI + YJIT | 175905.0 | 0.30× |
| JRuby | 310597.5 | 0.54× |
| TruffleRuby | 29052.3 | 0.05× |

#### transpose-24x24

| Runtime | ns/op | vs MRI |
| --- | ---: | ---: |
| **go-ruby (pure Go)** | 1312.2 | 0.37× |
| MRI | 3542.0 | 1.00× |
| MRI + YJIT | 3456.0 | 0.98× |
| JRuby | 1779.1 | 0.50× |
| TruffleRuby | 2033.2 | 0.57× |

24×24 float multiply is **~5× faster than MRI** (0.20×) and transpose ~2.7× (0.37×): MRI's `Matrix` is pure Ruby. TruffleRuby's Graal JIT compiles the multiply hot loop to native code and leads this row — a fair steady-state datapoint since the multiply loop is long enough to trigger compilation.

!!! note "Reproduce"
    The harness is committed under
    [`benchmarks/`](https://github.com/go-ruby-matrix/docs/tree/main/benchmarks):
    a self-contained Go driver (`go/`, pins the published library via
    `go.mod`), the equivalent `ruby/matrix.rb` workload, and `run.sh`. Run
    `bash benchmarks/run.sh`; env `OUTER`/`WARM` tune the pass budget and
    `RUBY`/`JRUBY`/`TRUFFLERUBY` select the runtime binaries.

!!! warning "Warm-up budget & noise — honest framing"
    Numbers reflect a **fixed warm-process budget** (3 warm-up + 25 timed passes
    in one process). The JVM/GraalVM JITs (JRuby, TruffleRuby) may need a larger
    warm-up to reach steady state, so their columns can **understate** peak
    throughput — most visibly TruffleRuby on the shortest loops (a few cold-JIT
    outliers are noted in the text). Sub-microsecond rows carry the most relative
    noise; treat those ratios as order-of-magnitude. Every number here is a
    **real measured value** from the dated run above — nothing is fabricated,
    estimated, or cherry-picked. The go-ruby column is the pure-Go library; every
    other column is that interpreter's own stdlib doing the equivalent work.

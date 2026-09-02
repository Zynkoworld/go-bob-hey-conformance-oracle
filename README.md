# zynko-oracle · `go-bob-hey-conformance-oracle`

**A deterministic, re-checkable conformance oracle for `bob` (go).**

## Proven
Measured on the canonical Exercism corpus — **16 input/output pairs, 4 distinct outputs** — produced by *running* the reference in a sealed sandbox, not asserted.

## Scope (declared)
The corpus is the canonical Exercism test data for `bob`. Inputs outside that set are **not covered**; this oracle decides agreement on the published corpus only and makes no claim of general correctness.

## Provenance
Reference: the Exercism reference solution for `bob` (go; MIT, Exercism), body unchanged. Proven by the exercism testsuite (pin=1b551ea9eae1e854), re-executed by harvest in a sealed sandbox (unshare -rn) before this bundle was generated.

## License
Apache-2.0 for the scaffolding; the reference body retains its upstream MIT (Exercism) license.

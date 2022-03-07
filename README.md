# anybase

[![CI](https://github.com/spenserblack/anybase/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/anybase/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/anybase/branch/main/graph/badge.svg?token=05dxpMxxTg)](https://codecov.io/gh/spenserblack/anybase)

Create an arbitrary base-n encoder and decoder with any set digits, with the limit that `n` is in `[2, 256)` and that each digit
can be represented as a single byte.

This does not truly allow *any* base-n, for the sake of API simplicity and familiarity based on the `encoding/hex` standard
library. Base-1 would be absurd and implemented as tallies (`1` = `1`, `2` = `11`, etc.), and Base-257+ would not be able
to be stored in a `byte`.

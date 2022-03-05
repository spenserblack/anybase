# anybase

[![CI](https://github.com/spenserblack/anybase/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/anybase/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/anybase/branch/main/graph/badge.svg?token=05dxpMxxTg)](https://codecov.io/gh/spenserblack/anybase)

Create an arbitrary base-n encoder and decoder with any set digits, with the limit that `n` is in `[2, 256)` and that the digits
can be represented as a single byte.

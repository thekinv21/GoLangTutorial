**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# Basic Types in Go

- **Integers:** int, int8, int16, int32, int64, and for unsigned uint, uint8 (byte), etc.
- **Floating point:** float32, float64
- **Text:** string, and rune (alias for int32) for a single character
- **Boolean:** bool
- **Complex numbers:** complex64, complex128

## Integers (Signed)

| Type             | Size        | Minimum                    | Maximum                    |
| ---------------- | ----------- | -------------------------- | -------------------------- |
| `int8`           | 8 bit       | -128                       | +127                       |
| `int16`          | 16 bit      | -32,768                    | +32,767                    |
| `int32` (`rune`) | 32 bit      | -2,147,483,648             | +2,147,483,647             |
| `int64`          | 64 bit      | -9,223,372,036,854,775,808 | +9,223,372,036,854,775,807 |
| `int`            | 32 / 64 bit | Platform dependent         | Platform dependent         |

## Integers (Unsigned)

| Type             | Size        | Minimum | Maximum                        |
| ---------------- | ----------- | ------- | ------------------------------ |
| `uint8` (`byte`) | 8 bit       | 0       | 255                            |
| `uint16`         | 16 bit      | 0       | 65,535                         |
| `uint32`         | 32 bit      | 0       | 4,294,967,295                  |
| `uint64`         | 64 bit      | 0       | 18,446,744,073,709,551,615     |
| `uint`           | 32 / 64 bit | 0       | Platform dependent             |
| `uintptr`        | 32 / 64 bit | 0       | Large enough to hold a pointer |

## Floating-Point Numbers

| Type      | Size   | Range (approx.)       | Precision     |
| --------- | ------ | --------------------- | ------------- |
| `float32` | 32 bit | ±1.18e-38 … ±3.4e38   | ~6-7 digits   |
| `float64` | 64 bit | ±2.23e-308 … ±1.8e308 | ~15-16 digits |

## Complex Numbers

| Type         | Size    | Description                            |
| ------------ | ------- | -------------------------------------- |
| `complex64`  | 64 bit  | Real and imaginary parts are `float32` |
| `complex128` | 128 bit | Real and imaginary parts are `float64` |

## Other Types

| Type     | Values                                 | Zero Value |
| -------- | -------------------------------------- | ---------- |
| `bool`   | `true` / `false`                       | `false`    |
| `string` | UTF-8 text (immutable)                 | `""`       |
| `byte`   | Alias for `uint8`                      | `0`        |
| `rune`   | Alias for `int32` (Unicode code point) | `0`        |

> Note: The zero value of all numeric types is `0`. Limits are available in the `math` package: `math.MaxInt8`, `math.MinInt64`, `math.MaxUint32`, `math.MaxFloat64`.

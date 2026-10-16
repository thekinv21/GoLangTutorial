[English](./README_EN.md) | **Turkish** | [Russian](./README_RU.md)

# Go’da Temel Tipler

- **Tam sayılar:** int, int8, int16, int32, int64, işaretsizler için uint, uint8 (byte) vb.
- **Ondalıklı:** float32, float64
- **Metin:** string, tek karakter için rune (int32 takma adı)
- **Mantıksal:** bool
- **Karmaşık sayı:** complex64, complex128

## Tam Sayılar (İşaretli)

| Tip              | Boyut       | Minimum                    | Maksimum                   |
| ---------------- | ----------- | -------------------------- | -------------------------- |
| `int8`           | 8 bit       | -128                       | +127                       |
| `int16`          | 16 bit      | -32.768                    | +32.767                    |
| `int32` (`rune`) | 32 bit      | -2.147.483.648             | +2.147.483.647             |
| `int64`          | 64 bit      | -9.223.372.036.854.775.808 | +9.223.372.036.854.775.807 |
| `int`            | 32 / 64 bit | Platforma göre             | Platforma göre             |

## Tam Sayılar (İşaretsiz)

| Tip              | Boyut       | Minimum | Maksimum                   |
| ---------------- | ----------- | ------- | -------------------------- |
| `uint8` (`byte`) | 8 bit       | 0       | 255                        |
| `uint16`         | 16 bit      | 0       | 65.535                     |
| `uint32`         | 32 bit      | 0       | 4.294.967.295              |
| `uint64`         | 64 bit      | 0       | 18.446.744.073.709.551.615 |
| `uint`           | 32 / 64 bit | 0       | Platforma göre             |
| `uintptr`        | 32 / 64 bit | 0       | Pointer tutacak kadar      |

## Ondalıklı Sayılar

| Tip       | Boyut  | Aralık (yaklaşık)     | Hassasiyet     |
| --------- | ------ | --------------------- | -------------- |
| `float32` | 32 bit | ±1.18e-38 … ±3.4e38   | ~6-7 basamak   |
| `float64` | 64 bit | ±2.23e-308 … ±1.8e308 | ~15-16 basamak |

## Karmaşık Sayılar

| Tip          | Boyut   | Açıklama                        |
| ------------ | ------- | ------------------------------- |
| `complex64`  | 64 bit  | Gerçek ve sanal kısım `float32` |
| `complex128` | 128 bit | Gerçek ve sanal kısım `float64` |

## Diğer Tipler

| Tip      | Değerler                             | Sıfır Değeri |
| -------- | ------------------------------------ | ------------ |
| `bool`   | `true` / `false`                     | `false`      |
| `string` | UTF-8 metin (değiştirilemez)         | `""`         |
| `byte`   | `uint8` takma adı                    | `0`          |
| `rune`   | `int32` takma adı (Unicode karakter) | `0`          |

> Not: Sayısal tiplerin sıfır değeri `0`'dır. Sınırlar için `math` paketi kullanılabilir: `math.MaxInt8`, `math.MinInt64`, `math.MaxUint32`, `math.MaxFloat64`.

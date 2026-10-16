[English](./README_EN.md) | [Turkish](./README_TR.md) | **Russian**

# Константы в Go: const

```bash
const Pi = 3.14159

const (
    Ponedelnik = iota  // 0
    Vtornik            // 1
    Sreda              // 2
)
```

**iota** — это счётчик, который увеличивается на единицу в каждой строке; идеально подходит для конструкций, похожих на enum.

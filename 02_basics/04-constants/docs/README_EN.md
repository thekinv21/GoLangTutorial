**English** | [Turkish](./README_TR.md) | [Russian](./README_RU.md)

# Constants in Go: const

```bash
const Pi = 3.14159

const (
    Monday = iota     // 0
    Tuesday           // 1
    Wednesday         // 2
)
```

**iota** is a counter that increases by one on each line; it is ideal for enum-like constructs.

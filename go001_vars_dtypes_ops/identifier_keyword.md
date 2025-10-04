# `Identifier`

**Predeclared Identifiers**

- For Constants: `true`, `false`, `iota`, `nil`
- For Types:
  - `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
  - `float32`, `float64`
  - `complex128`, `complex64`
  - `bool`
  - `byte`
  - `rune`
  - `string`
  - `error`
- For Functions: `make`, `len`, `cap`, `new`, `append`, `copy`, `close`,`delete`, `complex`, `real`, `imag`, `panic`, `recover`

**Blank identifier**  
`_` (underscore) used as an anonymous placeholder

**Exported variable**

- the variable that can be `accessed from other package`
- conditions:
  - first character should be Unicode upper case letter (lower case identifier using within same package).
  - the identifier should be declared in the `package block` or be a `variable`, `function`, `type`, or `method` name within that package.

#

# `Keywords`

|            |               |          |             |          |
| :--------: | :-----------: | :------: | :---------: | :------: |
|  `break`   |   `default`   |  `func`  | `interface` | `select` |
|   `case`   |    `defer`    |   `go`   |    `map`    | `struct` |
|   `chan`   |    `else`     |  `goto`  |  `package`  | `switch` |
|  `const`   | `fallthrough` |   `if`   |   `range`   |  `type`  |
| `continue` |     `for`     | `import` |  `return`   |  `var`   |

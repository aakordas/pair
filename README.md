# pair
Inspired by C++'s `std:pair`,  one can create pairs with elements of the same kind (i.e. [1, 2]) or of different kind (i.e. ["a", 1]). Custom/Composite types are also supported. Besides different construction functions, an equality function is also provided.

## Installation

To install, type
```
go get github.com/aakordas/pair
```

## Usage

To use it in a program, do
```
import "github.com/aakordas/pair"
```
## Examples

```go
import "github.com/aakordas/pair"

p := pair.Create()
p.Set(1, 2)
p1 := NewPair(1, 2)

p.Equals(p1) // => true

p1.Set(2, 1)
p.Equals(p1) // => false

p2 := NewPair(p, p1) // Valid pair

// While a pair can contain another pair, it cannot check for equality, because
// interface{} is not indexable. I'm considering changing the code to tackle
// this issue.
p2.Equals(p1) // PANIC!!
```

Check the test file for more examples.

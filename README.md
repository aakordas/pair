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
**Breaking changes in last update!**

```go
import "github.com/aakordas/pair"

var p Pair
p.Set(1, 2)

p1 := pair.Make(1, 2)

p.Equals(p1) // => true

p1.Set(2, 1)
p.Equals(p1) // => false

p2 := pair.Make(p, p1) // Valid pair

p2.Equals(p1) // => false
```

Check the test file for (a few) more examples.

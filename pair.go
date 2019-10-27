/* Inspired by C++ pairs. Currently it has rather basic functionality. One can
 * create pairs with elements of different kinds just fine. (a, b) != (b, a)
 */

package pair

type Pair []interface{}

// Create creates and returns a new pair.
func Create() (p Pair) {
	p = make([]interface{}, 2, 2)

	return p
}

// Set makes a pair with the provided arguments as values
func (p Pair) Set(e1, e2 interface{}) {
	p[0] = e1
	p[1] = e2
}

// NewPair creates and initializes a new pair and returns it.
func NewPair(e1, e2 interface{}) (p Pair) {
	p = make([]interface{}, 2, 2)

	p[0] = e1
	p[1] = e2

	return p
}

// Equals checks if two pairs are equal
func (p Pair) Equals(pa Pair) (bool) {
	if p[0] == pa[0] && p[1] == pa[1] {
		return true
	}

	return false
}

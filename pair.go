/* Inspired by C++ pairs. Currently it has rather basic functionality. One can
 * create pairs with elements of different kinds just fine. (a, b) != (b, a)
 */

package pair

type Pair [2]interface{}

// Make makes and returns a new pair with the provided values.
func Make(e1, e2 interface{}) (p Pair) {
	p[0] = e1
	p[1] = e2

	return p
}

// Set sets the pair's values as the provided ones.
func (p *Pair) Set(e1, e2 interface{}) {
	p[0] = e1
	p[1] = e2
}

// Equals checks if two pairs are equal.
func (p *Pair) Equals(pa Pair) bool {
	if p[0] == pa[0] && p[1] == pa[1] {
		return true
	}

	return false
}

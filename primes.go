package main

import (
	"math/big"
	"math/rand"
	"time"
)

// power calculates (base^exp) mod n efficiently using modular exponentiation.
func power(base, exp, n *big.Int) *big.Int {
	res := big.NewInt(1)
	base.Mod(base, n)
	for exp.Cmp(big.NewInt(0)) > 0 {
		if new(big.Int).And(exp, big.NewInt(1)).Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, base).Mod(res, n)
		}
		exp.Rsh(exp, 1)
		base.Mul(base, base).Mod(base, n)
	}
	return res
}

// isComposite checks if n is composite using the given base.
// It returns true if n is composite, and false if it's a strong probable prime.
func isComposite(n, d, s, base *big.Int) bool {
	x := power(base, d, n)
	if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
		return false
	}
	for i := big.NewInt(0); i.Cmp(s) < 0; i.Add(i, big.NewInt(1)) {
		x.Mul(x, x).Mod(x, n)
		if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
			return false
		}
	}
	return true
}

// MillerRabin performs the Miller-Rabin primality test on n with k iterations (witnesses).
// It returns true if n is likely prime, and false if n is definitely composite.
func MillerRabin(n *big.Int, k int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}
	if new(big.Int).And(n, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 { // Check if n is even and greater than 2
		return false
	}

	s := big.NewInt(0)
	d := new(big.Int).Sub(n, big.NewInt(1))
	for new(big.Int).And(d, big.NewInt(1)).Cmp(big.NewInt(0)) == 0 {
		d.Rsh(d, 1)
		s.Add(s, big.NewInt(1))
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for range k {
		base := new(big.Int).Rand(rng, new(big.Int).Sub(n, big.NewInt(3)))
		base.Add(base, big.NewInt(2)) // Ensure base is in the range [2, n-2]
		if isComposite(n, d, s, base) {
			return false // n is definitely composite
		}
	}

	return true // n is likely prime
}

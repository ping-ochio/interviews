package main

// This is how to manage big numbers in golang.

import (
	"fmt"
	"math/big"
)

// ---------------- MAIN --------------------------
func main() {
	k := big.NewInt(0)
	k.SetInt64(0)
	k = puzzle(big.NewInt(10))
	fmt.Printf("D:%v\n", k)

	k = puzzle(big.NewInt(100))
	fmt.Printf("D:%v\n", k)

	k.Exp(big.NewInt(2022), big.NewInt(100), nil)
	//k = puzzle(k)
	fmt.Printf("D:%v\n", k)
}

// ------------- function puzzle --------------------
func puzzle(N *big.Int) *big.Int {

	// Variables Inicialization
	A := big.NewInt(0)
	B := big.NewInt(1)
	C := big.NewInt(0)
	D := big.NewInt(0)
	X := big.NewInt(0)

	// ---- Variables temp
	a1 := big.NewInt(0)
	am := big.NewInt(0)

	b1 := big.NewInt(0)
	bm := big.NewInt(0)

	c1 := big.NewInt(0)
	cm := big.NewInt(0)
	val := big.NewInt(0)

	//----------------- Partial sum ------------------
	sAB := big.NewInt(0)  //--- (3 * B + 4 * A )
	sABC := big.NewInt(0) //---  (2 * C + 3 * B + 4 * A)

	A.SetInt64(1)
	B.SetInt64(1)
	C.SetInt64(1)
	D.SetInt64(1)
	X.SetInt64(1)
	a1.SetInt64(0)
	b1.SetInt64(0)
	c1.SetInt64(0)
	am.SetInt64(4)
	bm.SetInt64(3)
	cm.SetInt64(2)
	sAB.SetInt64(0)
	sABC.SetInt64(0)
	val.SetInt64(10000000000)

	i := big.NewInt(0)
	b := big.NewInt(1)
	for i.Cmp(N) < 0 {
		i.Add(i, b)
		a1.Mul(am, A)     // "4*A"
		b1.Mul(bm, B)     // "3*B"
		c1.Mul(cm, C)     // "2*C"
		sAB.Add(b1, a1)   //<-- 3*B + 4*A
		sABC.Add(c1, sAB) //<-- 2 * C + 3 * B + 4 * A

		X.Add(D, sABC)
		A.Set(B)
		B.Set(C)
		C.Set(D)
		D.Set(X)

	}

	return D.Rem(D, val)
}

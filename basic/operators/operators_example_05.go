package main

import "fmt"

func main() {

	var a int = 20
	var c int

	c = a
	fmt.Printf("c = a，c 值为 %d\n", c)

	c += a
	fmt.Printf("c += a，c 值为 %d\n", c)

	c -= a
	fmt.Printf("c -= a，c 值为 %d\n", c)

	c *= a
	fmt.Printf("c *= a，c 值为 %d\n", c)

	c /= a
	fmt.Printf("c /= a，c 值为 %d\n", c)

	fmt.Printf("\nc = %d，二进制为 %08b\n", c, c)

	c <<= 2
	fmt.Printf("c <<= 2，c 值为 %d (%08b)\n", c, c)

	c >>= 2
	fmt.Printf("c >>= 2，c 值为 %d (%08b)\n", c, c)

	c &= 2
	fmt.Printf("c &= 2，c 值为 %d (%08b)\n", c, c)

	c ^= 2
	fmt.Printf("c ^= 2，c 值为 %d (%08b)\n", c, c)

	c |= 2
	fmt.Printf("c |= 2，c 值为 %d (%08b)\n", c, c)
}

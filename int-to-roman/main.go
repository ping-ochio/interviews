package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var unip int
	// var decp int
	// var cenp int
	// var milp int

	var num string
	fmt.Println("Enter an integer between 0 and 9999 ")
	fmt.Scanf("%s", &num)
	fmt.Println(Toroman(num))
	fmt.Println(Suma(3, 5))

}

func Toroman(num string) string {

	uni := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	dec := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	cen := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	mil := []string{"", "M", "MM", "MMM", "⊺⊽", "⊽", "⊽Ī", "⊽ĪĪ", "⊽ĪĪĪ", "ĪẊ"}

	val := strings.Split(num, "")
	lv := len(val) - 1

	// switch lv {
	// case 3:
	// 	milp, _ = strconv.Atoi(val[lv-3])
	// 	fallthrough
	// case 2:
	// 	cenp, _ = strconv.Atoi(val[lv-2])
	// 	fallthrough
	// case 1:
	// 	decp, _ = strconv.Atoi(val[lv-1])
	// 	fallthrough
	// case 0:
	// 	unip, _ = strconv.Atoi(val[lv])
	// }

	// fmt.Printf("%s%s%s%s\n", mil[milp], cen[cenp], dec[decp], uni[unip])
	rom := make([]string, 0)
	p := [][]string{uni, dec, cen, mil}
	for i := lv; i >= 0; i-- {
		mp, _ := strconv.Atoi(val[lv-i])
		rom = append(rom, fmt.Sprintf("%v", p[i][mp]))
	}
	println()
	return strings.Join(rom, "")
}
func Suma(a, b int) int {
	return a + b

}

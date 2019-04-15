package numbers2string

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

var (
	num int
	err error
	)

func main() {
	fmt.Println("Enter number")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		num, err = strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Conversion failed")
		}
		starts(num)
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//num = input
	//var res string = ""
	//numlen := num.len()
}

func starts(n int) {
	var res string
	hundreds := n / 100
	tensUnits := intMod(int(n), 100)
	var smallnumbers = []string {
		"zero", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen",
	}

	var tensnumbers = []string{
		"", "", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety",
	}
	//var scaleNumbers = []string{
	//	"", "thousand", "million", "billion",
	//}
	if hundreds != 0 {
		res += smallnumbers[hundreds] + " hundred"
		if tensUnits != 0 {
			res += " and "
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		res += tensnumbers[tens]
		if units != 0 {
			res += "-"+smallnumbers[units]
		}
	} else if tensUnits != 0 {
		res += smallnumbers[tensUnits]
	}

	if n == 0 {
		res = smallnumbers[n]
	}

	fmt.Println(res)

}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

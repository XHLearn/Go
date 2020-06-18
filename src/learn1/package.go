package learn1

import "fmt"
import "regexp"
import "strconv"

func MyRegexp() {
	searchIn := "John:2578.55 William: 48335.55 Steve: 47.192"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

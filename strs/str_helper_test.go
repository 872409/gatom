package strs

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

const (
	regular = `^(13[0-9]|14[57]|15[0-35-9]|18[06-9])\d{8}$`
)

func IsMobileNum(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
func TestStrToInt(t *testing.T) {

	s := GStr("1").To(false)
	s2 := GStr("true").To(false)
	s3 := GStr("a").To(10)
	fmt.Println(s, s2, s3)

	val := StrToInt("a", 2)
	fmt.Println(val)
}

type A struct {
	Name string
}

func TestMatch(t *testing.T) {
	fmt.Println(strings.Contains("TigerwolfC", "wolf"))
	// var re = regexp.MustCompile(`(?:0(1?)|1(2?)|2(3?)|3(4?)|4(5?)|5(6?)|6(7?)|7(8?)|8(9?)|9(0?)){6}\d`)
	// var str = `12345678956789`
	//
	// for i, match := range re.FindAllString(str, 1) {
	// 	fmt.Println(match, "found at index", i)
	// }

	pat := `1{4,}|2{4,}|3{4,}|4{4,}|5{4,}|6{4,}|7{4,}|8{4,}|9{4,}`
	src := `111121444455556666338888339999451443333333334423111133557782222`
	// reg := regexp.MustCompile(pat)
	// fmt.Printf("%q\n", reg.FindAllString(src, -1))
	// fmt.Println(IsMobileNum("18606061024"),IsMobileNum("18006061024"))
	fmt.Println(regexp.MatchString(pat, src))
}

func TestStrToVal(t *testing.T) {

	// patterns := []string{
	// 	`([\d])\\1{1,}([\d])\\2{1,}`,                                                                         // AABB
	// 	`(?:0(? = 9)|9(? = 8)|8(? = 7)|7(? = 6)|6(? = 5)|5(? = 4)|4(? = 3)|3(? = 2)|2(? = 1)|1(? = 0)){3}\d`, // 匹配4位顺降 0987
	// 	`(?:0(? = 1)|1(? = 2)|2(? = 3)|3(? = 4)|4(? = 5)|5(? = 6)|6(? = 7)|7(? = 8)|8(? = 9)|9(? = 0)){3}\d`, // 匹配4位顺增
	// 	`([\d])\1{2,}`,                                                                                       // 匹配3位以上的重复数字
	// 	`(([\d]){1,}([\d]){1,})\1{1,}`,                                                                       // 匹配ABAB ABCABC ABCDABCD
	// };
	ms := map[string]A{"aa": {Name: "N"}}

	fmt.Println("aa", ms, ms["aaaa"].Name)

	//
	// val := StrTo("10a", 1)
	// val2 := StrTo("1", true)
	// val3 := StrTo("false", true)
	// val4 := StrTo("xx", true)
	//
	// fmt.Println(val, val2, val3, val4)
}

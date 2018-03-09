package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	one := re.Find([]byte(a))
	fmt.Println("Find: ", string(one))

	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll: ", all)

	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex: ", index)

	allIndex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex: ", allIndex)
	// 	Find:  am
	// FindAll:  [[97 109] [108 101 97 114] [110 105 110 103] [108 97 110 103] [117 97 103 101]]
	// FindIndex:  [2 4]
	// FindAllIndex:  [[2 4] [5 9] [9 13] [17 21] [21 25]]

	re2, _ := regexp.Compile("am(.*)lang(.*)")
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch:", submatch)

	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchIndex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchIndex)

	submatchAll := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchAll)

	submatchAllIndex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchAllIndex)
	/*
		am learning Go language
		 learning Go
		uage
		[2 25 4 17 21 25]
		[[[97 109 32 108 101 97 114 110 105 110 103 32 71 111 32 108 97 110 103 117 97 103 101] [32 108 101 97 114 110 105 110 103 32 71 111 32] [117 97 103 101]]]
		[[2 25 4 17 21 25]]
	*/

	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))

}

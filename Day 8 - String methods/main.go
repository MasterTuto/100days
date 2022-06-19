package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Menor(s1, s2 string) (string, string) {
	if strings.Compare(s1, s2) < 0 {
		return s1, s2
	}

	return s2, s1
}

func testCut() {
	show := func(c, s string) {
		b, a, f := strings.Cut(c, s)

		fmt.Printf("Cut(%v, %v) = (%v, %v, %v)\n", c, s, b, a, f)
	}

	show("Days of Code", "Days")
	show("Days of Code", "of")
	show("Days of Code", "Code")
	show("Days of Code", "Years")
	fmt.Printf("\n\n")
}

func testFields() {
	show := func(s string) {
		f := strings.Fields(s)
		fc := strings.FieldsFunc(s, func(r rune) bool {
			return r == ';' || r == ','
		})

		fmt.Printf("Fields(%v) = %v\n", s, f)
		fmt.Printf("FieldsFunc(%v, c == ';' or c == ',') = %v\n", s, fc)
	}

	show(" breno carvalho    da silva   ")
	show(" breno;carvalho,da silva;   ")
	fmt.Printf("\n\n")
}

func testHasPrefix() {
	show := func(s, p string) {
		fmt.Printf("HasPrefix(%v, %v) = %v\n", s, p, strings.HasPrefix(s, p))
	}

	show("Ofidioglota", "Ofidio")
	show("School", "School")
	show("Sons", "Parents")
	fmt.Printf("\n\n")
}

func testHasSuffix() {
	show := func(s, p string) {
		fmt.Printf("HasSuffix(%v, %v) = %v\n", s, p, strings.HasSuffix(s, p))
	}

	show("Ofidioglota", "glota")
	show("School", "School")
	show("Parents", "Sons")
	fmt.Printf("\n\n")
}

func testIndex() {
	show := func(s, ss string) {
		fmt.Printf("Index(%v, %v) = %v\n", s, ss, strings.Index(s, ss))
		fmt.Printf("IndexAny(%v, %v) = %v\n", s, ss, strings.IndexAny(s, ss))
		fmt.Printf("IndexByte(%v, %v) = %v\n", s, 'a', strings.IndexByte(s, 'a'))
		fmt.Printf("IndexRune(%v, %v) = %v\n", s, 'b', strings.IndexRune(s, 'b'))
		fmt.Printf("IndexFunc(%v, %v) = %v\n", s, ss, strings.IndexFunc(s, func(c rune) bool {
			return c == 'a' && rune(s[len(s)/2]) == c
		}))
	}

	show("Breno Carvalho da Silva", "Carvalho")
	show("Breno Carvalho da Silva", "Breno")
	show("Breno Carvalho da Silva", "a")
	show("Breno Carvalho da Silva", "Olavo de Carvalho")
	fmt.Printf("\n\n")
}

func testJoin() {
	show := func(parts []string, joiner string) {
		fmt.Printf("Join(%v, %v) = %v\n", parts, joiner, strings.Join(parts, joiner))
	}

	show([]string{"Breno", "Carvalho", "da", "Silva"}, " ")
	show([]string{"Breno", "Duoglas", "John", "Carlos"}, ";")
	fmt.Printf("\n\n")
}

func testMap() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'a':
			return 'a' + (r-'a'+13)%26
		default:
			return r
		}
	}

	show := func(s string) {
		fmt.Printf("Map(rot13, %v) = %v\n", s, strings.Map(rot13, s))
	}

	show("Golang")
	show("Foo bar")
	show("I am not a poet")
	fmt.Printf("\n\n")
}

func testLastIndex() {
	show := func(s, ss string) {
		fmt.Printf("LastIndex(%v, %v) = %v\n", s, ss, strings.LastIndex(s, ss))
		fmt.Printf("LastIndexAny(%v, %v) = %v\n", s, ss, strings.LastIndexAny(s, ss))
		fmt.Printf("LastIndexByte(%v, %v) = %v\n", s, 'a', strings.LastIndexByte(s, 'a'))
		fmt.Printf("LastIndexFunc(%v, %v) = %v\n", s, ss, strings.LastIndexFunc(s, func(c rune) bool {
			return c == 'a' && rune(s[len(s)/2]) == c
		}))
	}

	show("Breno Carvalho da Silva", "Carvalho")
	show("Breno Carvalho da Silva", "Breno")
	show("Breno Carvalho da Silva", "a")
	show("Breno Carvalho da Silva", "Olavo de Carvalho")
	fmt.Printf("\n\n")
}

func testRepeat() {
	show := func(text string, times int) {
		fmt.Printf("Repeat(%v, %v) = %v\n", text, times, strings.Repeat(text, times))
	}

	show("Breno ", 5)
	show("Golang", 0)
	fmt.Printf("\n\n")
}

func testReplace() {
	show := func(i, s1, s2 string, t int) {
		fmt.Printf("Replace(%v, %v, %v, %v) = %v\n", i, s1, s2, t, strings.Replace(i, s1, s2, t))
		fmt.Printf("ReplaceAll(%v, %v, %v) = %v\n", i, s1, s2, strings.ReplaceAll(i, s1, s2))
	}

	show("Golang is the worst language", "worst", "best", 1)
	show("Golang is the best language", "e", "X", 2)
	show("Golang is the best language", "e", "X", -1)
	fmt.Printf("\n\n")
}

func testSplit() {
	show := func(s, sep string) {

		fmt.Printf("Split(%v, %v) =  %v\n", s, sep, strings.Split(s, sep))
	}

	show("Breno Carvalho da Silva", ",")
	show("example,of,csv,file", ",")
	fmt.Printf("\n\n")
}

func testSplitAfter() {
	show := func(s, sep string) {

		fmt.Printf("SplitAfter(%v, %v) =  %v\n", s, sep, strings.SplitAfter(s, sep))
	}

	show("Testing this example", " ")
	show("We. are. legion.", ". ")
	fmt.Printf("\n\n")
}

func testSplitN() {
	show := func(s, sep string, n int) {

		fmt.Printf("SplitN(%v, %v) =  %v\n", s, sep, strings.SplitN(s, sep, n))
	}

	show("Breno Carvalho da Silva", ",", 1)
	show("example,of,csv,file", ",", 2)
	fmt.Printf("\n\n")
}

func testSplitAfterN() {
	show := func(s, sep string, n int) {

		fmt.Printf("SplitAfterN(%v, %v, %v) =  %v\n", s, sep, n, strings.SplitAfterN(s, sep, n))
	}

	show("Testing this example", " ", 1)
	show("We. are. legion.", ". ", 2)
	fmt.Printf("\n\n")
}

func testToTitleSpecial() {
	show := func(c unicode.SpecialCase, n, s string) {

		fmt.Printf("ToTitleSpecial(%v, %v) =  %v\n", n, s, strings.ToTitleSpecial(c, s))
	}

	show(unicode.AzeriCase, "unicode.AzeriCase", "Testing wtf this is")
	show(unicode.CaseRanges, "unicode.CaseRanges", "opaa")
	fmt.Printf("\n\n")
}

func testTrim() {
	show := func(s, cutset string) {

		fmt.Printf("Trim(%v, %v) =  %v\n", s, cutset, strings.Trim(s, cutset))
	}

	show("'''intelligent person'''", "'''")
	show("I am happy!!!", "!")
	fmt.Printf("\n\n")
}

func testTrimLeft() {
	show := func(s, cutset string) {

		fmt.Printf("TrimLeft(%v, %v) =  %v\n", s, cutset, strings.TrimLeft(s, cutset))
	}

	show("'''intelligent person'''", "'''")
	show("I am happy!!!", "!")
	fmt.Printf("\n\n")
}

func testTrimRight() {
	show := func(s, cutset string) {

		fmt.Printf("TrimRight(%v, %v) =  %v\n", s, cutset, strings.TrimRight(s, cutset))
	}

	show("'''intelligent person'''", "'''")
	show("I am happy!!!", "!")
	fmt.Printf("\n\n")
}

func testTrimSuffix() {
	show := func(s, cutset string) {

		fmt.Printf("TrimSuffix(%v, %v) =  %v\n", s, cutset, strings.TrimSuffix(s, cutset))
	}

	show("'''intelligent person'''", "'''")
	show("I am happy!!!", "!")
	fmt.Printf("\n\n")
}

func testTrimPrefix() {
	show := func(s, cutset string) {

		fmt.Printf("TrimPrefix(%v, %v) =  %v\n", s, cutset, strings.TrimPrefix(s, cutset))
	}

	show("'''intelligent person'''", "'''")
	show("I am happy!!!", "!")
	fmt.Printf("\n\n")
}

func main() {
	Nome := "breno Carvalho da silva"
	ClonedNome := strings.Clone(Nome)
	fmt.Println("Nome == ClonedNome: ", Nome == ClonedNome)

	OutroNome := "Joao Calisto de Araujo"

	s1, s2 := Menor(Nome, OutroNome)
	fmt.Printf("%v < %v\n", s1, s2)

	if contains := strings.Contains(Nome, "silva"); contains {
		fmt.Printf("\"%v\" contains \"silva\"\n", Nome)
	}

	oCount := strings.Count(OutroNome, "o")
	fmt.Printf("%v repeats %v %v times", OutroNome, "o", oCount)

	testCut()
	testFields()
	testHasPrefix()
	testHasSuffix()
	testIndex()
	testJoin()
	testLastIndex()
	testMap()
	testRepeat()
	testReplace()
	testSplit()
	testSplitN()
	testSplitAfterN()
	testToTitleSpecial()
	testTrim()
	testTrimLeft()
	testTrimRight()
	testTrimPrefix()
	testTrimSuffix()
}

package Exercises

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func GrowUpNumber(n int) []int {
	var result = make([]int, 0)

	for i := n; i >= 0; i-- {
		result = append(result, i)
	}

	return result
}

func TestGroupUpNumber(t *testing.T) {
	input := 5
	expected := []int{5, 4, 3, 2, 1, 0}

	result := GrowUpNumber(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GrowUpNumber(%d) = %v; want %v", input, result, expected)
	}
}

func isMatchIdentity(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			return false
		}
	}
	return true
}

func TestIsMatchIdentity(t *testing.T) {
	input := "aaaa" // true // bbbbb // true // cccaaa false
	expected := true

	result := isMatchIdentity(input)

	if result != expected {
		t.Errorf("isMatchIdentity(%q) = %v; want %v", input, result, expected)
	}
}

func alphabetCharacter(s string) string {
	charCount := make(map[rune]int)

	for _, v := range s {
		charCount[v]++
	}

	var chars []rune

	for k := range charCount {
		chars = append(chars, k)
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	result := ""
	for _, char := range chars {
		for i := 0; i < charCount[char]; i++ {
			result += string(char)
		}
	}

	return result
}

func TestAlphabetCharacter(t *testing.T) {
	input := "lamborgini"
	expected := "abgiilmnor"

	result := alphabetCharacter(input)

	if result != expected {
		t.Errorf("alphabetCharacter(%q) = %v; want %v", input, result, expected)
	}
}

func existHigherNumber(n []int, f int) bool {
	if len(n) == 0 {
		return false
	}

	for _, v := range n {
		if v > f {
			return true
		}
	}
	return false
}

func TestExistHigherNumber(t *testing.T) {
	input := []int{8, 4, 20, 32, 1}
	f := 10
	expected := true

	result := existHigherNumber(input, f)

	if result != expected {
		t.Errorf("existHigherNumber(%v, %v) = %v; want %v", input, f, result, expected)
	}
}

func howManyAlphabet(n int) string {
	l := "Lamborgini"
	a := strings.Repeat("a", n)
	return l[0:1] + a + l[1:]
}

func TestHowManyAlphabet(t *testing.T) {
	input := 4
	expected := "Laaaaamborgini"

	alphabet := howManyAlphabet(input)

	if alphabet != expected {
		t.Errorf("howManyAlphabet(%d) = %v; want %v", input, alphabet, expected)
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		right := len(s) - i - 1
		if string(s[i]) != string(s[right]) {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	input := "sos"
	expected := true

	result := isPalindrome(input)

	if result != expected {
		t.Errorf("isPalindrome(%q) = %v; want %v", input, result, expected)
	}
}

func removeCharacter(s []string) []string {
	const char = "C"
	var result = make([]string, 0)
	for _, v := range s {
		firstChar := string(v[0])
		if firstChar != char {
			result = append(result, v)
		}
	}
	return result
}

func TestRemoveCharacter(t *testing.T) {
	input := []string{"Cacing", "Bebek", "Cicak", "Beruang"}
	expected := []string{"Bebek", "Beruang"}

	result := removeCharacter(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeCharacter(%v) = %v; want %v", input, result, expected)
	}
}

func matchDictionary(s string, d []string) []string {
	match := make([]string, 0)

	for _, v := range d {
		if strings.Contains(v, s) {
			match = append(match, v)
		}
	}
	return match
}

func TestMatchDictionary(t *testing.T) {
	input := []string{"kuda", "kurus", "kece", "kucel"}
	s := "ku"

	expected := []string{"kuda", "kurus", "kucel"}

	result := matchDictionary(s, input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("matchDictionary(%v, %q) = %v; want %v", input, s, result, expected)
	}
}

func additiveInverse(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = -num
	}
	return result
}

func TestAdditiveInverse(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{-1, -2, -3, -4, -5}

	result := additiveInverse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("additiveInverse(%v) = %v; want %v", input, result, expected)
	}
}

func reverseBool(b interface{}) (bool, string) {
	bb := reflect.ValueOf(b)

	if bb.Kind() == reflect.Bool {
		return !bb.Bool(), "valid type"
	} else {
		return false, "bukan boolean"
	}
}

func TestReverseBool(t *testing.T) {
	input := 10
	expected := false
	expectedMessage := "bukan boolean"

	b, s := reverseBool(input)

	if b != expected || s != expectedMessage {
		t.Errorf("reverseBool(%v) = (%v, %q); want (%v, %q)", input, b, s, expected, expectedMessage)
	}
}

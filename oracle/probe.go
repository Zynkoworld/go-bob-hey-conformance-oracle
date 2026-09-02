package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



import "strings"

// Hey returns Bob's responses to a given remark.
func Hey(remark string) string {
	switch remark = strings.TrimSpace(remark); {
	case silent(remark):
		return "Fine. Be that way!"
	case yelling(remark):
		if asking(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case asking(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

func yelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != strings.ToUpper(remark)
}

func asking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func silent(remark string) bool {
	return remark == ""
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []string{"Does this cryogenic chamber make me look fat?", "WATCH OUT!", "", "Tom-ay-to, tom-aaaah-to.", "fffbbcbeab?", ":) ?", "Wait! Hang on. Are you going to be OK?", "\nDoes this cryogenic chamber make\n me look fat?", "DO LIONS EAT PEOPLE? AHHHHH.", "1, 2, 3 GO!", "I HATE THE DENTIST", "          ", "\n\r \t", "It's OK if you don't want to go work for NASA.", "1, 2, 3", "         hmmmmmmm..."}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, Hey(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}

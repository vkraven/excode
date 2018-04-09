package main

import (
	"fmt"
//	"errors"
	"unicode"
//	"unicode/utf8"
	"bufio"
	"os"
	"strconv"
//	"math/rand"
	"strings"
)

const VERSION = "0.03a"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func translate(me rune, n int) byte {
	board := "abcdefghijklmnopqrstuvwxyz"
	var out byte
	ischar := false
	isupper := unicode.IsUpper(me)
	act := unicode.ToLower(me)
	for x, y := range board {
		if act == y {
			newi := ((x + n) % 25)
			if newi < 0 {
				newi += 25
			}
			out = board[newi]
			ischar = true
		}
	}

	if ischar == false {
		switch me {
		case ' ':
			out = byte(me)
		default:
			out = byte(me)
		}
	}

	if isupper {
		x := rune(out)
		x = unicode.ToUpper(x)
		out = byte(x)
	}
	return out
}

func strshift(me string, n int) string {
	var out []byte
	for _, y := range me{
		this := translate(y, n)
		fmt.Printf(string(y) + " -> " + string(this) + "\n")
		out = append(out, this)
	}

	return string(out)

}

func caser(in string, n int) string {
	out := in

	if n < 3 {
		if n == 1 {
			out = strings.ToUpper(in)
		} else if n == 2 {
			out = strings.ToLower(in)
		}
	} /* else {
		if n == 4 {
			out = in
		} else if n == 3 {
			for x, y := range out {
				rand.Seed(time.Now().UnixNano())
				seed := rand.Intn(100)
				y = y
				if seed < 50 {
					y = unicode.ToUpper(rune(in[x]))
				} else { y = unicode.ToLower(rune(in[x])) }
			}*/
	return out

}

func punctuatindexor(me string, n int) string {
	var out string
	i := 0
	state := false
	for _, y := range me {
		if  (i % 4) == 0 && i != 0 && n == 5 && state == false {
			out = out + " "
			state = true
		}

		if unicode.IsSpace(y) {
			if n == 2 || n == 4 {
				out = out + string(y)
				i += 1
				state = false
			}
		} else if unicode.IsPunct(y) {
			if n == 1 || n == 4 {
				out = out + string(y)
				i += 1
				state = false
			}
		} else {
			out = out + string(y)
			i += 1
			state = false
		}
	}

	if n == 5 {
		out = strings.ToUpper(out)
	}
	return out
}

func main() {
/*	test := "Hello There!"	as the variable names suggest, this part was used for testing and is no longer needed.
	fmt.Println(test)
	var final string
	final = strshift(test, 5)
	fmt.Println(final)
	var final2 string
	final2 = strshift(final, -5)
	fmt.Println(final2)*/

	reader := bufio.NewReader(os.Stdin)
	var nexttest string
	fmt.Printf("ExCode\nVersion " + VERSION + "\nNB: This is a toy. \n\nPlease provide the message you wish encrypted. As this is an alphabetical substitution cypher, it will not operate over digits. ExCode advises the user to spell out their numbers in alphabets instead.\n\nPlease enter your message:\n")
	nexttest, err := reader.ReadString('\n')
	check(err)
	nexttest = nexttest[:len(nexttest)-1]

	var info string
	var shby int
	fmt.Println("Now provide the degree of shift by which you wish your message encyphered:\n(Input a number between 1 and 25. 0 is supported, but trivial, since an alphabetical shift by 0 is identical to the original alphabet as it is. A shift by 26 is equally frivolous.)\n\nYour desired degree of shift:")
	info, err = reader.ReadString('\n')
	check(err)
	shby, err = strconv.Atoi(info[:len(info)-1])
	check(err)

	var andcase string
	var cs int
	fmt.Println("Indicate if the message's upper and lower casings should be altered. The code may be in: \n1 - ALL CAPS;\n2 - all lowercase; or\n3 - Left as is.\n\n(Observe that leaving the cases as is will render your coded message far easier to break (as per the capitalisation rules of your language.)\n\nEnter a number between 1 and 3:")
	andcase, err = reader.ReadString('\n')
	check(err)
	cs, err = strconv.Atoi(andcase[:len(andcase)-1])
	check(err)

	var puncind string
	var pi int
	fmt.Println("Now, indicate if and how the punctuation and indentation of the message should be scrambled. The encoder may:\n1 - Removeallspaces;\n2  Remove all punctuation\n3RemoveAllSpacesAndPunctuationOr\n4 - Leave them as is; or\n5 - Enigma mode.\n\n(Note once again that leaving the punctuation and spaces in the message renders the original message far easier to derive from the code alone, as it clearly delineates words, sentences and prefixes/suffixes in the code.)\n\nEnter a number between 1 and 5:")
	puncind, err = reader.ReadString('\n')
	check(err)
	pi, err = strconv.Atoi(puncind[:len(puncind)-1])
	check(err)

	var nextfinal string
	nextfinal = strshift(nexttest, shby)
	fmt.Println(nextfinal)
	nextfinal = caser(nextfinal, cs)
	fmt.Println(nextfinal)
	nextfinal = punctuatindexor(nextfinal, pi)
	fmt.Println(nextfinal)
}

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

const (
	VERSION = "0.07a"
	FREQUENCY = "etaoinsrhldcumfpgwybvkxjqz"
	MOSTCOMMON = "etaoinshrld"
	LEASTCOMMON = "zqjxkvb"
	board = "abcdefghijklmnopqrstuvwxyz"

)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func translate(me rune, n int) byte {
	var out byte
	ischar := false
	isupper := unicode.IsUpper(me)
	act := unicode.ToLower(me)
	for x, y := range board {
		if act == y {
			newi := ((x + n) % 26)
			if newi < 0 {
				newi += 26
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
//		fmt.Printf(string(y) + " -> " + string(this) + "\n")
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

func buildkeyarray(key string) []int {
	var outbase []int

	work := key

	work = strings.ToLower(work)

	for _, j := range work {
		x := strings.Index(board, string(j))
		outbase = append(outbase, x)
	}

	return outbase
}

func codevigcypher(reader *bufio.Reader) {
	fmt.Println("\nVigenère Cypher Encryption Mechanism Enabled - \n")
	var plaintext string
	var err error
//	var ptlength int
	var keylength int

	fmt.Printf("\nPlease provide the message you wish encrypted.\nAs this is an alphabetical substitution cypher, it will NOT operate over digits, punctuation, or whitespaces. ExCode advises the user to spell out their numbers in alphabetical letters instead, and to consider scrambling character casing and sentence formatting in the options that follow.\n\nPlease enter your message:\n")
	plaintext, err = reader.ReadString('\n')
	check(err)
	plaintext = plaintext[:len(plaintext)-1]
//	ptlength = len(plaintext)

	var key string
	unacceptablekey := true
	fmt.Printf("\nNow, provide your secret passphrase.\nThis passphrase will be used to encypher and decypher your message. Therefore, ExCode expects the user NOT to share this passphrase with anyone the user wishes to keep this message from. Only share this secret passphrase with the person with whom you wish to communicate - be it another person or your future self. As the Vigenère cypher is also alphabetical, your passphrase MUST be alphabetical and cannot contain digits, whitespace, or punctuation. (Observe, that by the operation of the Vigenère cypher, a single character passphrase is logically equivalent to a Shift (Caesar) cypher encryption. Observe also that a passphrase consisting only of 'a's and/or 'A's or both will result in a vacuous encryption).\n\nEnter your secret passphrase:\n")
	for unacceptablekey {
		key, err = reader.ReadString('\n')
		check(err)
		key = key[:len(key)-1]
		unacceptablekey = false
		for _, j := range key {
			if unicode.IsPunct(j) || unicode.IsSpace(j) || unicode.IsDigit(j) {
				unacceptablekey = true
			}
		}

		if unacceptablekey {
			fmt.Printf("\nThe secret passphrase contains unacceptable characters, and cannot encypher the above message.\nPlease enter your secret passphrase (WITHOUT spaces, digits, or punctuation): ")
		} else { keylength = len(key) }
	}

	var andcase string
	var cs int
	fmt.Println("Indicate if and how the message's casing should be altered.\nNote that although leaving the cases as they will likely make the coded message easier to break, this may be avoided with a short enough message and a long enough passphrase. ExCode still recommends, however, that the cases be operated on to ensure the greatest possible secrecy. Choose between:\n 1 -\tALL CAPS;\n2 -\tall lowercase; or\n3 -\tkeep casing as is.\n\nIndicate your desired casing operation:")
	andcase, err = reader.ReadString('\n')
	check(err)
	andcase = andcase[:len(andcase)-1]
	cs, err = strconv.Atoi(andcase)
	check(err)

	var puncind string
	var pi int
	fmt.Println("\nNow, indicate if and how the punctuation and indentation of the message should be scrambled. The encoder may:\n\n 1-Removeallspaces;\n 2\t Remove all punctuation\n 3RemoveAllSpacesAndPunctuationOr\n 4 -\tLeave them as is; or\n 5 -\tEnigma mode.\n\n(Note once again that leaving the punctuation and spaces in the message renders the original message somewhat easier to derive from the code alone, but that this can be overcome with a strong enough message/passphrase combination.)\n\nEnter a number between 1 and 5:")
	puncind, err = reader.ReadString('\n')
	check(err)
	pi, err = strconv.Atoi(puncind[:len(puncind)-1])
	check(err)

	keybase := buildkeyarray(key)

	semifinal := envig(plaintext, keybase, keylength)

	semifinal = caser(semifinal, cs)
	semifinal = punctuatindexor(semifinal, pi)
	fmt.Println(semifinal)

}

func envig(me string, keybase []int, keylength int) string {
	count := 0
	var semifinal string
	semifinal = ""
	for _, y := range me {
		if unicode.IsLetter(y) {
			ourn := (count % keylength)
			semifinal += string(translate(y, keybase[ourn]))
			count += 1
		} else {
			semifinal += string(y)
		}
	}
	return semifinal
}

func devig(me string, keybase []int, keylength int) string {
	workbase := keybase
	for i, j := range keybase {
		workbase[i] = (0 - j)
	}
	semifinal := envig(me, workbase, keylength)

	return semifinal

}

func codeunvigcypher(reader *bufio.Reader) {
	var err error
	var cyphertext string
	fmt.Println("Enter the Vigenere Cyphertext:")
	cyphertext, err = reader.ReadString('\n')
	check(err)
	cyphertext = cyphertext[:len(cyphertext)-1]

	var key string
	fmt.Println("Enter your key:")
	key, err = reader.ReadString('\n')
	check(err)
	key = key[:len(key)-1]
	keylength := len(key)

	keybase := buildkeyarray(key)

	final := devig(cyphertext, keybase, keylength)
	fmt.Println(final)
}

func codeshiftcypher(reader *bufio.Reader, message string, degofshift string, casing int, punctex int) {
	fmt.Println("\nSubstitution Cypher Encryption Mechanism Enabled - \n")
	var nexttest string
	var err error
	if message == "" {
	fmt.Printf("\nPlease provide the message you wish encrypted.\nAs this is an alphabetical substitution cypher, it will NOT operate over digits, punctuation, or whitespaces. ExCode advises the user to spell out their numbers in alphabetical letters instead.\n\nPlease enter your message:\n")
	nexttest, err = reader.ReadString('\n')
	check(err)
	nexttest = nexttest[:len(nexttest)-1]
	} else { nexttest = message }

	var info string
	var shby int
	if degofshift == "" {
	fmt.Println("\nNow provide the degree of shift by which you wish your message encyphered:\n(Input a number between 1 and 25. 0 is supported, but trivial, since an alphabetical shift by 0 is identical to the original alphabet as it is. A shift by 26 is equally frivolous.)\nNB: Enter 13 here for ROT13 fuctionality.\n\nYour desired degree of shift:")
	info, err = reader.ReadString('\n')
	check(err)
	shby, err = strconv.Atoi(info[:len(info)-1])
	check(err)
	} else {
		info = degofshift
		shby, err = strconv.Atoi(info)
		check(err)
	}

	var andcase string
	var cs int
	if casing == 0 {
	fmt.Println("\nIndicate if the message's upper and lower casings should be altered. The code may be in: \n1 - ALL CAPS;\n2 - all lowercase; or\n3 - Left as is.\n\n(Observe that leaving the cases as is will render your coded message far easier to break - as per the capitalisation rules of your language.)\n\nEnter a number between 1 and 3:")
	andcase, err = reader.ReadString('\n')
	check(err)
	cs, err = strconv.Atoi(andcase[:len(andcase)-1])
	check(err)
	} else { cs = casing }

	var puncind string
	var pi int
	if punctex == 0 {
	fmt.Println("\nNow, indicate if and how the punctuation and indentation of the message should be scrambled. The encoder may:\n\n 1-Removeallspaces;\n 2\t Remove all punctuation\n 3RemoveAllSpacesAndPunctuationOr\n 4 -\tLeave them as is; or\n 5 -\tEnigma mode.\n\n(Note once again that leaving the punctuation and spaces in the message renders the original message far easier to derive from the code alone, as it clearly delineates words, sentences and prefixes/suffixes in the code.)\n\nEnter a number between 1 and 5:")
	puncind, err = reader.ReadString('\n')
	check(err)
	pi, err = strconv.Atoi(puncind[:len(puncind)-1])
	check(err)
	} else { pi = punctex }

	var nextfinal string
	nextfinal = strshift(nexttest, shby)
//	fmt.Println(nextfinal)
	nextfinal = caser(nextfinal, cs)
//	fmt.Println(nextfinal)
	nextfinal = punctuatindexor(nextfinal, pi)
	fmt.Println("ENCYPHERED MESSAGE:\n" + nextfinal)

}

func maxintarray(x [26]int) ([]int, int) {
	max := 0
	var index []int
	for _, j := range x {
		if j > max {
			max = j
		}
	}
	for i, j := range x {
		if j == max {
			index = append(index, i)
		}
	}
	return index, max
}

func generatebase(input string) [26]int {
	var base [26]int
	for index, _ := range base {
		base[index] = 0
	}

	for _, y := range input {
		place := strings.Index(board, string(y))
		if place >= 0 {
			base[place] += 1
		}
	}

	return base

}
func tryfreq(ind int, base [26]int) [2][11]int64 {
	var n int
	var letter rune
	var scoreboard [2][11]int64
	for s, t := range MOSTCOMMON {
		letter = t

		for x, y := range board {
			if y == letter {
				n = (ind - x)
				break
			}
		}
		n = (0 - n)

		var score int64
		score = 0

		for i, j := range base {
			state := ((i+n) % 26)
			for state < 0 {
				state += 26
			}
			points := strings.Index(FREQUENCY, string(board[state]))
			score += int64(points*j)
		}
		scoreboard[0][s] = int64(n)
		scoreboard[1][s] = score
	}
	return scoreboard

}
func breakshiftcypher(reader *bufio.Reader) {


	var ct string
	fmt.Println("Please provide the coded message you desire ExCode to break:")
	ct, err := reader.ReadString('\n')
	check(err)
	ct = ct[:len(ct)-1]

	spacecount, punctcount, digitcount, lettercount := 0, 0, 0, 0

	for _, y := range ct {
		if unicode.IsSpace(y) {
			spacecount += 1
		} else if unicode.IsPunct(y) {
			punctcount += 1
		} else if unicode.IsDigit(y) {
			digitcount += 1
		} else if unicode.IsLetter(y) {
			lettercount += 1
		}
	}

	fmt.Printf("\n\nAnalysis:\n\n\tSpaces:\t%d\n\tPunctuation:\t%d\n\tDigits:\t%d\n\tLetters:\t%d\n\n", spacecount, punctcount, digitcount, lettercount)

	var base [26]int
	for index, _ := range base {
		base[index] = 0
	}

	work := ct
	work = strings.ToLower(work)
	base = generatebase(work)
	stark := "["
	for _, y := range board {
		stark = stark + string(y) + " "
	}
	stark += "]"
//	fmt.Println(stark)
//	fmt.Println(base)
	list, _ := maxintarray(base)
//	fmt.Println(list)
//	fmt.Println(max)

	testscore := tryfreq(list[0], base)
//	fmt.Println(testscore)

	var tind int
	var t2nd int
//	var t3rd int
	min := testscore[1][0]
	for i, j := range testscore[1] {
		if j < min {
			min = j
			tind = i
		}
	}
	min = testscore[1][0]
	for i, j := range testscore[1] {
		if j < min && i != tind {
			min = j
			t2nd = i
		}
	}
//	min = testscore[1][0]
//	for i, j := range testscore[1] {
//		if j < min && i != tind && i != t2nd {
//			min = j
//			t3rd = i
//		}
//	}
	fmt.Println("\n\nAnalysis complete. Presenting the two most likely possibilities:\n\n")
	fmt.Println(strshift(ct, int(testscore[0][tind])) + "\n")
	fmt.Println(strshift(ct, int(testscore[0][t2nd])) + "\n")
//	fmt.Println(strshift(ct, int(testscore[0][t3rd])) + "\n")
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
	fmt.Printf("ExCode\nVersion " + VERSION + "\nNB: This is a toy. \n\n")

	optionchosen := true
	for optionchosen {
	fmt.Printf("Awaiting your command: ")
	encorbreak, err := reader.ReadString('\n')
	check(err)
	encorbreak = encorbreak[:len(encorbreak)-1]

	switch encorbreak {
	case "code":
	//	optionchosen = false
		codeshiftcypher(reader, "", "", 0, 0)
	case "break":
	//	optionchosen = false
		breakshiftcypher(reader)
	case "rot13":
	//	optionchosen = false
		codeshiftcypher(reader, "", "13", 3, 4)
	case "vig":
		codevigcypher(reader)
	case "devig":
		codeunvigcypher(reader)
	case "exit":
		optionchosen=false
		break
	case "quit":
		optionchosen=false
		break
	default:
		fmt.Println("\nInvalid command. (As of version " + VERSION + ", ExCode accepts only three commands - 'code', 'break' and 'rot13'. The first two are to encyper and decypher a given message respectively. The rot13 option allows easy access to ROT-13 translation capabilities.)")
	}
	}
}

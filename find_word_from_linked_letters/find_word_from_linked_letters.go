package findwordfromlinkedletters

import "strings"

// fmt.Println(findWord([]string{"P>E", "E>R", "R>U"}))                                                  // PERU
// 	fmt.Println(findWord([]string{"I>N", "A>I", "P>A", "S>P"}))                                           // SPAIN
// 	fmt.Println(findWord([]string{"U>N", "G>A", "R>Y", "H>U", "N>G", "A>R"}))                             // HUNGARY
// 	fmt.Println(findWord([]string{"I>F", "W>I", "S>W", "F>T"}))                                           // SWIFT
// 	fmt.Println(findWord([]string{"R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"}))                      // PORTUGAL
// 	fmt.Println(findWord([]string{"W>I", "R>L", "T>Z", "Z>E", "S>W", "E>R", "L>A", "A>N", "N>D", "I>T"})) // SWITZERLAND

func FindWord(linkedLetters []string) string {
	linkedLettersMap := make(map[string]string)

	lettersCounters := make(map[string]int)

	// count occurrences of each letter
	for _, rule := range linkedLetters {
		kv := strings.Split(rule, ">")
		linkedLettersMap[kv[0]] = kv[1]
		lettersCounters[kv[0]]++
		lettersCounters[kv[1]]++
	}

	// letters with only one occurrence are either
	// - the firstOrNext letter (if it occurs as key)
	// - or the last one otherwise
	var firstOrNext, last string
	for letter, counter := range lettersCounters {
		if counter == 2 {
			continue
		}

		if _, occursAsKey := linkedLettersMap[letter]; occursAsKey {
			firstOrNext = letter
			continue
		}
		last = letter

		if len(firstOrNext) > 0 && len(last) > 0 {
			break
		}
	}

	word := firstOrNext
	for {
		firstOrNext = linkedLettersMap[firstOrNext]
		word += firstOrNext
		if firstOrNext == last {
			break
		}
	}

	return word
}

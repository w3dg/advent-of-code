package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"

	aoc "github.com/w3dg/aocutils"
	"github.com/zyedidia/generic/set"
)

type Card struct {
	ID      int
	Winning []int
	Have    []int
}

type Win struct {
	sourcecard int
	wins       int
}

var Wins []Win = []Win{}
var cardsmap map[int]int = make(map[int]int)

func p1(cards []Card) int {
	r := 0
	for _, card := range cards {
		cardsmap[card.ID] = 1
		wst, hst := set.NewMapset(card.Winning...), set.NewMapset(card.Have...)
		commonset := wst.Intersection(hst)
		winnings := commonset.Keys()
		l := len(winnings)
		Wins = append(Wins, Win{sourcecard: card.ID, wins: l})

		r += int(math.Pow(2, float64(l-1)))
	}
	return r
}

func p2(cards []Card) int {
	for _, v := range Wins {
		for i := range v.wins {
			newcardid := v.sourcecard + (i + 1)           // index to card number offset
			cardsmap[newcardid] += cardsmap[v.sourcecard] // add by current amount of cards to next totals, to accumulate them
		}
	}
	s := 0
	for _, v := range cardsmap {
		s += v
	}
	return s
}

func main() {
	infile := "./sample.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}

	lines, err := aoc.ReadFileLines(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	getNums := func(numline string) []int {
		l := len(numline)
		nums := []int{}

		for i := 0; i < l; i += 3 {
			sn := numline[i : i+2]
			// fmt.Println(sn)
			sn = strings.TrimSpace(sn)
			nums = append(nums, aoc.ParseNumOrPanic(sn))
		}

		return nums
	}

	cards := []Card{}
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		cardid, rest := parts[0], parts[1]

		idstr := strings.TrimPrefix(cardid, "Card ")
		re := regexp.MustCompile(`\d+`)
		num := re.FindString(idstr)
		id := aoc.ParseNumOrPanic(num)

		numbers := strings.Split(rest, " | ")

		w := []int{}
		h := []int{}

		for i := range numbers {
			switch i {
			case 0:
				// winning set
				w = append(w, getNums(numbers[i])...)
			case 1:
				// have set
				h = append(h, getNums(numbers[i])...)
			default:
				panic("More than two splits in card numbers")
			}
		}

		c := Card{ID: id, Winning: w, Have: h}
		cards = append(cards, c)
	}

	// for _, c := range cards {
	// 	fmt.Println(c)
	// }
	r1, r2 := p1(cards), p2(cards)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}

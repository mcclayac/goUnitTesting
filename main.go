package main

import (
	"fmt"
	"poetry"
)

func main() {

	p1 := poetry.NewPoem()
	p2 := poetry.Poem{{"The mortal fruit upon the bough",
		"Hands above the nuptial bed.",
		"The cat-bird in the tree returns",
		"The forfeit of his mutual vow.",
	}, {"The hard, untimely apple of",
		"The branch that feeds on watered rain,",
		"Takes the place upon her lips",
		"Of her late lamented love."},
		{"Many hands together press,",
			"Shaped within a static prayer",
			"Recall to one the chorister",
			"Docile in his sexless dress.",
		}}

	p1 = p1

	v, c, q := p2.Stats()
	//v, c := poetry.
	fmt.Printf("Valowels: %d, Constanants: %d,  Punctuations: %d\n\n", v, c, q)
	fmt.Printf("Stanzaz: %d,  Lines: %d ", p2.NumStanzas(), p2.NumLines())

}

package main

import(
	"fmt"
	"log"
	"../datafile"
)

func main(){
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil{
		log.Fatal(lines)
	}
	counts := make(map[string]int) // map
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}

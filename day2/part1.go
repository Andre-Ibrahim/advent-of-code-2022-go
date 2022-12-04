package main
import ("fmt"
		"io/ioutil" 
		"log"
		"strings"
	)

func main(){
	stream, err:= ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	arr := strings.Split(readString, "\n")
	// A: Rock, B: Paper, C: scisor
	// X: Rock, Y: Paper, Z: scisor

	var roundScore int = 0;

	for _, s := range arr {
		v := strings.Split(s, " ")
		
		if v[1] == "X" {
			roundScore += 1
		}else if v[1] == "Y" {
			roundScore += 2
		}else if v[1] == "Z" {
			roundScore += 3
		}

		if (([]rune(v[0])[0]) + 1 + 1) % 3 == ([]rune(v[1])[0] - 1) % 3 {
			roundScore += 6
		} else if (([]rune(v[0])[0]) + 1 - 1) % 3 == ([]rune(v[1])[0] - 1) % 3 {
			roundScore += 0
		}else {
			roundScore += 3
		}
		
	}

	fmt.Println(roundScore)
}

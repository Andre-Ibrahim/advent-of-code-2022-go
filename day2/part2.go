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
	// X: Lose, Y: Draw, Z: win

	strArr := []string{"A", "B", "C"}

	var roundScore int = 0;

	for _, s := range arr {
		v := strings.Split(s, " ")
		
		if v[1] == "X" {
			roundScore += 0 + getPoints(strArr[(([]rune(v[0])[0]) + 1 - 1) % 3])
		}else if v[1] == "Y" {
			roundScore += 3 + getPoints(strArr[(([]rune(v[0])[0]) + 1) % 3])
		}else if v[1] == "Z" {
			roundScore += 6 + getPoints(strArr[(([]rune(v[0])[0]) + 1 + 1) % 3])
		}
		
	}

	fmt.Println(roundScore)
}


func getPoints(s string) int{
	if(s == "A"){
		return 1
	}

	if(s == "B"){
		return 2
	}

	if(s == "C"){
		return 3
	}

	return 0
}

package main
import ("fmt"
		"io/ioutil" 
		"log"
		"strings"
		"strconv"
	)

func main(){
	stream, err:= ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	arr := strings.Split(readString, "\n")

	var total int = 0
	totals := []int{}
	for _, s := range arr {
		if len(s) > 0 {
			intVar, err := strconv.Atoi(s)
    		if err != nil {
        		// ... handle error
        	panic(err)
    		}

			total = total + intVar
		}
		
		if len(s) == 0 {
			totals = append(totals, total)
			total = 0
		}
		
	}

	max := []int{0, 0, 0}

	for _, m := range totals {
		if m > max[0] {
			max[2] = max[1]
			
			max[1] = max[0]
			
			max[0] = m

		}else if m > max[1] {

			max[2] = max[1]

			max[1] = m
		}else if m > max[2] {
			max[2] = m
		}
	}

	var t int = 0

	for _, v := range max {
		t += v
	}

	fmt.Println(t)
}

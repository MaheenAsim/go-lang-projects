// Task ---> 4

// prob_3 prompts the user for numbers, sums them, and returns the total.
func prob_3() float64 {
	var input string
	var total float64

	for {
		fmt.Print("Enter a number to continue: ")
		fmt.Scanln(&input)
		number, y := strconv.ParseFloat(input, 64)
		if y != nil {
			break // Exit loop if the input is not a valid float
		}
		total += number
	}
	return total // ask: for this function do we consider ints as floats so not problematic?
	// TA said it was ok
}

func main(){
  // Task ---> 4

	result := prob_3()
	fmt.Printf("The running total was %v\n", result)

}


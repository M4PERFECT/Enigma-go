package enigmago

import (
	"bufio"
	"fmt"
	"os"
)

interface component{
 encyrptForword()
}

func main() {
	for {
		fmt.Println("Welcome to the Enigma machine. Please select an option:-")
		fmt.Println("1. encryption")
		fmt.Println("2. decryption")
		fmt.Println("3. exit")

		fmt.Println("option(1,2 or 3):")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		switch input {
		case "1":
			encryption()
		case "2":
			decryption()
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Please enter one of the three options.")
			continue
		}
	}
}

// Example configuration where each letter maps to the next letter in the alphabet
// A -> B, B -> C, C -> D, ..., Y -> Z, Z -> A
var exampleWiring = [26]rune{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	0,
}

func encryption() {
	fmt.Println("Starting encryption")
	r := newRotor(1, 0, 1, 1, exampleWiring)
	r.encryption()
	return
}

func decryption() {
	panic("unimplemented")
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2 {
		fmt.Println("numgame new <username>")
		fmt.Println("numgame list")
		return
	} else if len(os.Args) == 3 {
		// username := os.Args[2] şimdilik
		fmt.Println("Zorluk seçin lütfen\n1) Kolay (1-25 | 1 Katsayı)\n2) Orta (1-50 | 3 Katsayı)\n3) Zor (1-75 | 5 Katsayı)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if len(scanner.Text()) != 1 {
			fmt.Println("Hatalı giriş!")
			return
		} else {
			if len(scanner.Text()) == 1 {
				zorluk, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Sayı metine dönüştürülemedi")
					return
				}
				// generatedNum := numGenerator(zorluk) şimdilik
			} else {
				fmt.Println("Hatalı giriş")
			}
		}
	} else {
		leaderboard()
	}
}

func numGenerator(zorluk int) (generatedNum int) {
	uzunluk := (zorluk * 25)
	generatedNum = rand.Intn(uzunluk) + 1
	return generatedNum
}

func leaderboard() {
	fmt.Println("blabla") // şimdilik
}

func game(num int, zorluk int, userName string) {

}

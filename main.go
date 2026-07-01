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
		username := os.Args[2]
		fmt.Println("Zorluk seçin lütfen\n1) Kolay (1-25 | 2 Katsayı)\n2) Orta (1-50 | 4 Katsayı)\n3) Zor (1-75 | 6 Katsayı)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if len(scanner.Text()) != 1 {
			fmt.Println("Hatalı giriş!")
			return
		} else {
			zorluk, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Sayı metine dönüştürülemedi")
				return
			}
			generatedNum, boyut, hak, katSayi := numGenerator(zorluk)
			game(generatedNum, boyut, hak, katSayi, username)
		}
	} else {
		leaderboard()
	}
}

func numGenerator(zorluk int) (generatedNum int, boyut int, hak int, carpan int) {
	uzunluk := (zorluk * 25)
	generatedNum = rand.Intn(uzunluk) + 1
	tahminHakki := rand.Intn(zorluk*2) + 3
	katSayi := (zorluk * 2)
	return generatedNum, uzunluk, tahminHakki, katSayi
}

func leaderboard() {
	fmt.Println("blabla") // şimdilik
}

func game(num int, zorluk int, tahminHakki int, katSayi int, userName string) {
	fmt.Printf("1-%d arasında bir sayı tuttum. %d Adet hakkın var. Tahmin et!\n", zorluk, tahminHakki)
	for i := 1; i <= tahminHakki; i++ {
		var tahmin int
		_, err := fmt.Scan(&tahmin)
		if err != nil {
			fmt.Println("Geçersiz giriş yaptınız.")
			return
		}
		if tahmin == num {
			var score int
			if i != tahminHakki {
				score = katSayi * (tahminHakki - i)
			} else {
				score = katSayi
			}
			fmt.Printf("Tebrikler! %d. tahmininizde doğru bildiniz. Puanınız -> %d\n", i, score)
			return
		} else if tahmin > num {
			fmt.Printf("Girdiğiniz sayı çok büyük. %d Adet hakkınız kaldı.", (tahminHakki - i))
		} else {
			fmt.Printf("Girdiğiniz sayı çok küçük. %d Adet hakkınız kaldı.", (tahminHakki - i))
		}
	}
	fmt.Printf("Maalesef sayıyı bilemedin. Tuttuğum sayı %d idi", num)
}

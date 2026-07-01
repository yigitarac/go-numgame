package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type Player struct {
	Nickname string
	Score    int
}

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
		path, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Home dizini bulunamadı")
			return
		}
		leaderboard(path)
	}
}

func numGenerator(zorluk int) (generatedNum int, boyut int, hak int, carpan int) {
	uzunluk := (zorluk * 25)
	generatedNum = rand.Intn(uzunluk) + 1
	tahminHakki := rand.Intn(zorluk*2) + 3
	katSayi := (zorluk * 2)
	return generatedNum, uzunluk, tahminHakki, katSayi
}

func leaderboard(yol string) {
	dosya := filepath.Join(yol, ".numgame_scoreboard.json")
	okunanOyuncular, err := os.ReadFile(dosya)
	if err != nil {
		fmt.Println("Dosya okunamadı.")
		return
	}
	var gamer []Player
	err = json.Unmarshal(okunanOyuncular, &gamer)
	sort.Slice(gamer, func(i, j int) bool {
		if gamer[i].Score > gamer[j].Score {
			return true
		} else {
			return false
		}
	})
	for i := range gamer {
		fmt.Printf("%d. %s - Skor: %d", (i + 1), gamer[i].Nickname, gamer[i].Score)
	}
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
			var gamer Player
			gamer.Nickname = userName
			gamer.Score = score
			path, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Home dizini bulunamadı.")
				return
			}
			oyunuKaydet(gamer, path)
			return
		} else if tahmin > num {
			fmt.Printf("Girdiğiniz sayı çok büyük. %d Adet hakkınız kaldı.", (tahminHakki - i))
		} else {
			fmt.Printf("Girdiğiniz sayı çok küçük. %d Adet hakkınız kaldı.", (tahminHakki - i))
		}
	}
	fmt.Printf("Maalesef sayıyı bilemedin. Tuttuğum sayı %d idi", num)
}

func oyunuKaydet(oyuncu Player, yol string) {
	dosya := filepath.Join(yol, ".numgame_scoreboard.json")
	dosyaIcerigi, err := os.ReadFile(dosya)
	if err != nil {
		if os.IsNotExist(err) {

		} else {
			fmt.Println("Dosya içeriği okunamadı")
			return
		}
	}
	var gamer []Player
	err = json.Unmarshal(dosyaIcerigi, &gamer)
	gamer = append(gamer, oyuncu)
	eklenmisDosya, err := json.Marshal(gamer)
	os.WriteFile(dosya, eklenmisDosya, 0644)
}

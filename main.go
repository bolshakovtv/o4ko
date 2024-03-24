package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// rnd - генератор псевдослучайных чисел
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// randCard получает псевдо случайное число от 2 до 11 включительно, кроме 5
func randCard() int {
	for {
		card := rnd.Intn(10) + 2
		if card != 5 {
			return card
		}
	}
}

// input запрашивает и возвращает ввод пользователя в консоли
func input(title string) string {
	fmt.Print(title)
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println(err)
	}
	return s
}

// pullingCard достает псевдо случайную карту из колоды
func pullingCard(cards map[int]string, deck map[string]int) int {
	for {
		cardValue := randCard()
		deck[cards[cardValue]]--
		fmt.Println(deck)
		if deck[cards[cardValue]] >= 0 {
			return cardValue
		}
	}
}

// total очищает терминал и выводит итоговую сумму очков игрока
func total(sumUser int) {

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Printf("Total: %d\n", sumUser)
}

func main() {

	// Объявление мапы значений карт в колоде
	cards := map[int]string{2: "Валет", 3: "Дама", 4: "Король", 6: "Шестерка", 7: "Семерка", 8: "Восьмерка",
		9: "Девятка", 10: "Десятка", 11: "Туз"}

	// Объявление мапы количества карт в колоде
	deck := map[string]int{"Валет": 4, "Дама": 4, "Король": 4, "Шестерка": 4, "Семерка": 4, "Восьмерка": 4, "Девятка": 4,
		"Десятка": 4, "Туз": 4}

	var cardValue int
	var sumBot int
	var sumUser int
	var answer string
	botPlaying := true
	userPlaying := true
	gameFinish := false
	duration, _ := time.ParseDuration("2s")

	total(sumUser)
	for gameFinish == false {

		if botPlaying {
			switch {
			case sumBot >= 16 && sumBot <= 21 && (!userPlaying || sumBot > sumUser):
				fmt.Println("Компьютер решил остановиться")
				botPlaying = false
			default:
				cardValue = pullingCard(cards, deck)
				sumBot += cardValue
				fmt.Println("Ход компьютера...")
				time.Sleep(duration)
				//fmt.Printf("Компютер вытащил: %s\n", cards[cardValue])
				time.Sleep(duration)
				total(sumUser)
				if sumBot > 21 {
					fmt.Printf("У компьютера перебор: %d. Вы победили. Поздравляю!\n", sumBot)
					return
				}

			}
		}

		if userPlaying {
			fmt.Println("Ваш ход...")
			for answer != "n" && answer != "y" {
				answer = input("Тянете карту? Введи букву [y] для согласия или [n] для отказа: ")
			}
			switch answer {
			case "y":
				cardValue = pullingCard(cards, deck)
				sumUser += cardValue
				fmt.Printf("Вы вытащили: %s\n", cards[cardValue])
				time.Sleep(duration)
				total(sumUser)
				if sumUser > 21 {
					fmt.Printf("У вас перебор: %d. Компьютер победил. Потрачено...\n", sumUser)
					return
				}
			case "n":
				userPlaying = false
			}
			answer = ""
		}

		if !botPlaying && !userPlaying {
			gameFinish = true
		}
	}
	if sumBot > sumUser {
		fmt.Printf("Копьютер набрал: %d, Вы набрали: %d. Компьютер победил.\n", sumBot, sumUser)
		return
	} else if sumBot == sumUser {
		fmt.Printf("Копьютер набрал: %d, Вы набрали: %d. Ничья!\n", sumBot, sumUser)
		return
	}
	fmt.Printf("Копьютер набрал: %d, Вы набрали: %d. Вы победили, поздравляю!\n", sumBot, sumUser)
	return
}

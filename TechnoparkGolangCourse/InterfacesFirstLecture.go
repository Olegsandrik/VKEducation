package main

import (
	"fmt"
	"strconv"
)

// Для проверки того, что тесты покрывают весь код, требуется запустить все тесты покета! (Важно) go test ./...
// Выношу в отдельном файле, так как слишком много всего в TestingGoFirstLecture, а тема важная
// Утиная типизация

type Duck interface { // Все что умеет плавать, летать и крякать удовлетворяет этому интерфейсу
	Fly()
	Swim()
	Cruck()
}

type Payer interface { // Им можно платить
	Pay(int) error
}

type Wallet struct { // Структура кошелька
	Cash int
}

func (w *Wallet) Pay(amount int) error { // Кошелек умеет Pay!
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) { // Работает у всех, кто умеет платить
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными?")
	case *Phone:
		fmt.Println("Крутая мобилка!")
	}

	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	myPhone := &Phone{Money: 1000}
	Buy(myPhone)

	PayForMeWithPhone(myPhone)

	fmt.Printf("Raw payment : %#v\n", myWallet)
	fmt.Printf("Способ оплаты: %s\n", myWallet)
}

type Phone struct {
	Money   int
	AppleID string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("Not enough money on account")
	}
	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("PLease, enter phone")
	}
	return nil
}

// --------------

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	Payer
	Ringer
}

func PayForMeWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Ошибка при оплате %v\n\n", err)
		return
	}
	fmt.Printf("Турникет открыт через %T\n\n", phone)
}

func (w *Wallet) String() string { // Без него будет бред при выводе с %s
	return "Кошелёк в котором " + strconv.Itoa(w.Cash) + " денег"
}

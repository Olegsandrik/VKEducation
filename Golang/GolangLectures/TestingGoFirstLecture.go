package main

import (
	"fmt"
	"unicode/utf8"
)

type UserId int
type AllUsers map[UserId]User
type User struct {
	UserName    string // Public поля, если с маленькой, то это private
	UserAddress string
	UserAge     int
}

const ( // Кроме Гошки нигде такой красоты нет
	ziro = iota
	_    // Underscore - отсутствие, которое можно использовать при дебаге
	two
	tree
)

const ( // Не только инкриминирование. Можно задать любую удобную итерацию
	start = iota
	first = iota * 10
	second
	third
)

func SuperFunc(shrek []string) {
	shrek[0] = "You change this bro"
	fmt.Printf("MyStr = %v внутри SuperFunc\n", shrek)

}

func NotSuperFunc(shrek []string) {
	shrek = append(shrek, "Is it!?")
	fmt.Printf("MyStr = %v внутри NotSuperFunc\n", shrek)
}

func GigaSuperFunc(shrek *[]string) {
	*shrek = append(*shrek, "It's work honey")
	fmt.Printf("MyStr = %v внутри GigaSuperFunc\n", *shrek)
}

func main() {
	fmt.Print("Пошла работа с массивами!\n")

	FastPeremen := "1234" // Быстрое объявление переменной

	var array [10]int // Объявление массива, все значения стоковые => нулевые в случае int
	_ = array[0]      // Обращение к 0му элементу массива
	// _ = array[11]     // Panic

	// Слайсы = динамические массивы
	slice := make([]int, 10, 20)
	// len = 10, capacity = 20. Первые 10 эл-тов 0, можно добавить еще 10 без аллоцирования памяти
	slice = append(slice, 12, 19, 228)                                                     // Просто добавляем элементы
	fmt.Printf("Элементы слайса %d %d %d %d\n", slice[1], slice[10], slice[11], slice[12]) // 0, 12, 19, 228

	fmt.Printf("Массивы правят миром %v\n", array)
	// Рисует весь массив [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("Массивы правят миром %#v\n", array)
	// Рисует весь массив и данные об этом массиве [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Print("Пошла работа со сканом!\n")

	var ScamInt int
	fmt.Println("Введите ваше число")
	_, err := fmt.Scanf("%d", &ScamInt)
	if err != nil {
		fmt.Print("Error")
		return
	}

	fmt.Println("Введите вашу строчку")
	var ScamStr string
	_, err = fmt.Scanf("%s", &ScamStr)
	if err != nil {
		fmt.Print("Error")
		return
	}

	fmt.Printf("Ваше число: %d, ваша строчка: %s, быстрая переменная: %s\n", int64(ScamInt), ScamStr, FastPeremen)
	// ScamInt приводим руками к типу int64

	fmt.Println("Работа с iota №1: ", ziro, two, tree)             // 0 2 3
	fmt.Println("Работа с iota №2: ", start, first, second, third) // 0 10 20 30

	fmt.Print("Пошла работа с функциями!\n")
	MyStr := []string{"1234"}
	fmt.Printf("MyStr = %v изначально\n", MyStr)
	SuperFunc(MyStr)
	fmt.Printf("MyStr = %v после работы с SuperFunc\n", MyStr)
	NotSuperFunc(MyStr)
	fmt.Printf("MyStr = %v после работы с NotSuperFunc\n", MyStr)
	GigaSuperFunc(&MyStr)
	fmt.Printf("MyStr = %v после работы с GigaSuperFunc\n", MyStr)

	fmt.Print("Пошла работа со строчками!\n")
	var MyRune rune = '茶' // Она uint32 для UTF-8 символов
	fmt.Printf("Да здравствуют руны из коробки %#U\n", MyRune)
	// Строки неизменяемые, каждый символ это руна.
	SymbolsInFastPerem := utf8.RuneCountInString(FastPeremen) // Кол-во символов в строчке FastPerem - 4
	fmt.Println("Количество символов в FastPerem ", SymbolsInFastPerem)
	fmt.Print("Пошла работа с HashMap!\n")
	Users := map[int]string{} // Создаем Хеш-табличку с парой ключ-int значение-string
	Users[1] = "Alex"         // Добавили значение с ключом 1
	fmt.Printf("Значение хеш-таблицы с ключом 1: %s и количество элементов в ней: %d\n", Users[1], len(Users))
	// проверка на существование ключа
	_, mNameExist := Users[2]
	fmt.Println("mNameExist:", mNameExist)
	// удаление ключа (clear удалит вообще все)
	delete(Users, 1)
	fmt.Printf("Количество элементов в хеш-таблице: %d\n", len(Users))

	// Указатели
	a := 2
	b := &a // Содержит адрес а, то есть является указателем
	*b = 3  // Сделает a = 3
	c := &a // Новый указатель на переменную a

	// получение указателя на переменнут типа int
	// инициализировано значением по-умолчанию
	d := new(int) // на 0 указывает
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c и a не изменились

	c = d   // теперь с указывает туда же, куда d
	*c = 14 // с = 14 -> d = 14, a = 12

	// простое условие
	boolVal := true
	if boolVal {
		fmt.Println("boolVal is true")
	}

	// Мапки без порядка!
	mapVal := map[string]string{"name": "rvasily"}
	// получаем только признак сущестования ключа
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	// switch по 1 переменной
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough // Проваливает в следующий case. Если это отсутствует, то после case выход из switch
	case "test", "lastName":
		// some work
	default:
		// some work
	}

	str := "Привет, Мир! 💩"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}

	fmt.Println("Пошла работа с defer!")
	defer fmt.Println("After work") // <- Стек выполнения функций
	defer func() {
		fmt.Println(getSomeVars())
	}()
	fmt.Println("Some userful work\n")

	// recover помогает спастись от panic

	MyUser := User{UserName: "Alex", UserAddress: "Moscow", UserAge: 19}
	fmt.Printf("Данные о пользователе: %v\n", MyUser)
	MyUser.SetName("Nikita")
	fmt.Printf("Данные о пользователе полсе SetName: %v\n", MyUser)
	MyUser.UpdateName("Danila")
	fmt.Printf("Данные о пользователе полсе UpdateName: %v\n", MyUser)

}

// Несколько результатов, last это всегда ошибка по договоренности
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// Много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// не фиксированное количество параметров
func sum(in ...int) (result int) {
	for _, val := range in {
		result += val
	}
	return
}

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"
}

func SumAges(p1, p2 User) int { // Метод, возвращающий сумму возрастов двух пользователей
	return p1.UserAge + p2.UserAge
}

// не изменит оригинальной структуры, для который вызван
func (p User) UpdateName(name string) {
	p.UserName = name
}

// изменяет оригинальную структуру
func (p *User) SetName(name string) {
	p.UserName = name // Компилятор сам все понимает
}

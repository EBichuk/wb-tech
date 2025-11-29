package main

/*
Что выведет программа?
Объяснить вывод программы.
*/

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return &customError{}
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*
Выведет:
error

Почему?
Переменная err объявлена под типом интерфейса error, ей присваивается результат функции test.
Функция test() возвращает указатель на тип customError, который удовлетворяет интерфейсу error (реализован метод Error() string).
err не будет равно nil, так как в структуре itab будет храниться значение конкретного типа в поле type
*/

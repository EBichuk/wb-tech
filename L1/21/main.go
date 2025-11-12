package main

import "fmt"

/*
Реализовать паттерн проектирования «Адаптер» на любом примере.

Применимость:
Паттерн адаптер позволяет без изменения кода уже существующего класса и его интерфейса, использовать их в
нашей системе

+
1. Позволяет использовать сторонние компоненты, не изменяя существующий код
2. Позволяет использовать структуры с разными интерфейсами совместно

-
1. Усложнение кода, так как появляются новые структуры и интерфесы
2. Если структуры слишком разные, то адаптер будет слишком большим

Примеры использования:
Для разных хранилищ данных

*/

// Существующая структура
type MySQL struct {
}

func (a *MySQL) Find(query string) []string {
	return []string{"item1", "item2"}
}

// Потребитель
type DB interface {
	Query(sql string) []string
}

// Адаптер, реализующий несовместимый интерфейс
type MySQLAdapter struct {
	*MySQL
}

func NewMySQLAdapter(m *MySQL) DB {
	return &MySQLAdapter{m}
}

func (a *MySQLAdapter) Query(sql string) []string {
	return a.Find(sql)
}

func main() {
	db := NewMySQLAdapter(&MySQL{})

	fmt.Println(db.Query("SELECT * FROM items"))
}

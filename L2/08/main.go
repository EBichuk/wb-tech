package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
Создать программу, печатающую точное текущее время с использованием NTP-сервера.

Реализовать проект как модуль Go.
Использовать библиотеку ntp для получения времени.
Программа должна выводить текущее время, полученное через NTP (Network Time Protocol).
Необходимо обрабатывать ошибки библиотеки: в случае ошибки вывести её текст в STDERR и вернуть ненулевой код выхода.
Код должен проходить проверки (vet и golint), т.е. быть написан идиоматически корректно.
*/

func GetNtpTime(address string) (time.Time, error) {
	return ntp.Time(address)
}

func main() {
	ntpServer := flag.String("ntp", "0.ru.pool.ntp.org", "address ntp server")
	flag.Parse()

	time, err := GetNtpTime(*ntpServer)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(time)
}

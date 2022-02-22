package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

const ticketPrice = float32(10.50)

type reserveInfo struct {
	qtd   int
	total float32
}

func printTable(reserve map[string]reserveInfo) {
	fmt.Println("Tickets reservados:")
	fmt.Println("Nome \t \t Quantidade \t \tTotal ")
	for key, element := range reserve {
		fmt.Printf("%v \t \t %v \t \t R$ %v\n", key, element.qtd, element.total)
	}
}

func main() {
	var reserve = make(map[string]reserveInfo)

	fmt.Println("Bem vindo ao sistema de reserva de tickets")
	for {
		fmt.Print("Qual o seu nome? ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()

		fmt.Print("Quantos tickets você quer reservar? ")
		var tickets int
		fmt.Scanln(&tickets)

		if reflect.TypeOf(tickets).Kind() != reflect.Int {
			fmt.Println("type fo tickets: ", reflect.TypeOf(tickets).Kind())
		}
		total := float32(tickets) * ticketPrice

		fmt.Print(name, ", você confirma a reserva de ", tickets, " tickets, com o total de ", total, "? (s/n): ")

		var response string
		fmt.Scanln(&response)
		if response != "s" && response != "S" {
			printTable(reserve)
			break
		}

		reserve[name] = reserveInfo{qtd: reserve[name].qtd + tickets, total: reserve[name].total + total}

		fmt.Println("Wow... seus tickets foram reservados!!!")
		fmt.Print("Quer reservar mais tickets? (s/n): ")
		fmt.Scanln(&response)
		if response != "s" && response != "S" {
			printTable(reserve)
			break
		}
	}
}

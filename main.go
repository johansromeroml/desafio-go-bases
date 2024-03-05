package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	ticketsCSV, err := os.Open("tickets.csv")
	defer func(f *os.File) {
		f.Close()
		fmt.Println("End of execution")

	}(ticketsCSV)

	if err != nil {
		panic(err)
	}
	//fmt.Println("ticketsCSV: ", ticketsCSV)

	ticketsData, err := io.ReadAll(ticketsCSV)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("ticketsData: ", string(ticketsData))

	ticketsPreSlice := strings.Split(string(ticketsData), "\n")

	//fmt.Println("tickets Pre slice: ", ticketsPreSlice)
	ticketList := tickets.Tickets{}

	for _, ticketPre := range ticketsPreSlice[:len(ticketsPreSlice)-1] {
		tick := tickets.TicketFromString(ticketPre)
		//fmt.Println(tick)
		ticketList.Tickets = append(ticketList.Tickets, tick)
	}

	//fmt.Println(ticketList.Tickets[1])

	totalTickets, err := tickets.GetTotalTickets("Brazil", &ticketList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Numero total de tickets : ", totalTickets)

	earlyMorning, morning, afternoon, evening, err := tickets.GetTicketsByDayframe(&ticketList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Tickets de madrugada: ", earlyMorning)
	fmt.Println("Tickets de mañana: ", morning)
	fmt.Println("Tickets de tarde: ", afternoon)
	fmt.Println("Tickets de noche: ", evening)
	//fmt.Println("tut", earlyMorning+morning+afternoon+evening)

	avgToBrazil, err := tickets.AverageDestination("Brazil", &ticketList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Average tickets to Brazil: ", avgToBrazil)
}

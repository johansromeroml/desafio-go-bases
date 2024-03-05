package tickets

import (
	"strconv"
	"strings"
)

type Ticket struct {
	ID      int
	Name    string
	Email   string
	Country string
	Hour    string
	Price   int
}

type Tickets struct {
	Tickets []Ticket
}

func TicketFromString(in string) Ticket {
	//println("--", in)
	attributes := strings.Split(in, ",")
	//println(attributes)
	id, _ := strconv.Atoi(attributes[0])
	price, _ := strconv.Atoi(attributes[5])
	return Ticket{ID: id, Name: attributes[1], Email: attributes[2], Country: attributes[3], Hour: attributes[4], Price: price}
}

// ejemplo 1
func GetTotalTickets(destination string, t *Tickets) (total int, err error) {
	total = 0
	for _, ticket := range t.Tickets {
		if ticket.Country == destination {
			total++
		}
	}
	return
}

// ejemplo 2
func GetTicketsByDayframe(t *Tickets) (a, b, c, d int, err error) {
	a = 0
	b = 0
	c = 0
	d = 0
	for _, ticket := range t.Tickets {
		tim := strings.Split(ticket.Hour, ":")
		hours, _ := strconv.Atoi(tim[0])
		//mins, _ := strconv.Atoi(tim[1])
		switch {
		case hours <= 6:
			a++
		case hours >= 7 && hours <= 12:
			b++
		case hours >= 13 && hours <= 19:
			c++
		case hours >= 20:
			d++
		}
	}
	return
}

// ejemplo 3
func AverageDestination(destination string, t *Tickets) (avg float64, err error) {
	flightsToDestination, err := GetTotalTickets(destination, t)
	totalTickets := len(t.Tickets)
	//println("f2d: ", flightsToDestination, ", tut: ", totalTickets)
	avg = float64(flightsToDestination) * 100 / float64(totalTickets)
	return
}

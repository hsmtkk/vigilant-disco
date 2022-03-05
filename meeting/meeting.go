package meeting

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func Arrange(name string) string {
	guests := NewGuests(name)
	sort.Slice(guests, func(i, j int) bool {
		if guests[i].LastName == guests[j].LastName {
			return guests[i].FirstName < guests[j].FirstName
		} else {
			return guests[i].LastName < guests[j].LastName
		}
	})
	return FormatGuests(guests)
}

type Guest struct {
	FirstName string
	LastName  string
}

func (g Guest) String() string {
	return fmt.Sprintf("(%s, %s)", g.LastName, g.FirstName)
}

func NewGuest(name string) (Guest, error) {
	elems := strings.Split(name, ":")
	if len(elems) != 2 {
		return Guest{}, fmt.Errorf("invalid format; %s", name)
	}
	firstName := strings.ToUpper(elems[0])
	lastName := strings.ToUpper(elems[1])
	return Guest{FirstName: firstName, LastName: lastName}, nil
}

func NewGuests(names string) []Guest {
	guests := []Guest{}
	ns := strings.Split(names, ";")
	for _, name := range ns {
		guest, err := NewGuest(name)
		if err != nil {
			log.Print(err)
		}
		guests = append(guests, guest)
	}
	return guests
}

func FormatGuests(guests []Guest) string {
	ss := []string{}
	for _, guest := range guests {
		ss = append(ss, guest.String())
	}
	return strings.Join(ss, "")
}

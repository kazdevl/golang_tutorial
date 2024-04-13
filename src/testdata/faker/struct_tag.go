package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

type TargetWithTags struct {
	Latitude           float32           `faker:"lat"`
	Longitude          float32           `faker:"long"`
	RealAddress        faker.RealAddress `faker:"real_address"`
	CreditCardType     string            `faker:"cc_type"`
	Email              string            `faker:"email"`
	IPV4               string            `faker:"ipv4"`
	IPV6               string            `faker:"ipv6"`
	Password           string            `faker:"password"`
	Jwt                string            `faker:"jwt"`
	PhoneNumber        string            `faker:"phone_number"`
	MacAddress         string            `faker:"mac_address"`
	URL                string            `faker:"url`
	UserName           string            `faker:"username"`
	Name               string            `faker:"name"`
	UnixTime           int64             `faker:"unix_time"`
	Date               string            `faker:"date"`
	Time               string            `faker:"time"`
	MonthName          string            `faker:"month_name"`
	Year               string            `faker:"year"`
	DayOfWeek          string            `faker:"day_of_week"`
	DayOfMonth         string            `faker:"day_of_month"`
	Currency           string            `faker:"currency"`
	AmountWithCurrency string            `faker:"amount_with_currency"`
	UUID               string            `faker:"uuid_digit"`
	Skip               string            `faker:"-"`
	AccountId          int               `faker:"oneof: 15, 20, 25"`
}

func createDummys() []*TargetWithTags {
	as := make([]*TargetWithTags, 0, 3)
	for range 3 {
		a := TargetWithTags{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}
		as = append(as, &a)
	}
	return as
}

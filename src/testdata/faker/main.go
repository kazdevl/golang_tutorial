package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

type Target struct {
	Name string
	Age  int
}

type TargetWithTags struct {
	Latitude           float32           `faker:"lat"`
	Longitude          float32           `faker:"long"`
	RealAddress        faker.RealAddress `faker:"real_address"`
	CreditCardNumber   string            `faker:"cc_number"`
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
	DayOfMonth         string            `faker:"day_of_week"`
	Currency           string            `faker:"currency"`
	AmountWithCurrency string            `faker:"amount_with_currency"`
	UUID               string            `faker:"uuid"`
	Skip               string            `faker:"-"`
	AccountId          int               `faker: "oneof: 15, 20, 25"`
}

func main() {
	var t Target
	_ = faker.FakeData(&t)
	fmt.Printf("Target: %+v\n", t)

	fmt.Println("*****")

	for i := 0; i < 10; i++ {
		var t2 TargetWithTags
		_ = faker.FakeData(&t2)
		fmt.Printf("TargetWithTags: %+v\n", t2)
	}
}

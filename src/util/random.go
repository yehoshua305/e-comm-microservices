package util

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
)

var rng *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"
var firstNames = []string{"John", "Jane", "Alex", "Emily", "Chris", "Katie", "Michael", "Sarah", "David", "Laura"}
var lastNames = []string{"Smith", "Doe", "Johnson", "Brown", "Wilson", "Moore", "Taylor", "Anderson", "Thomas", "Jackson"}

var streets = []string{"Main St", "High St", "Park Ave", "Maple St", "Oak St", "Pine St", "Cedar St", "Elm St", "Washington St", "Lake St"}
var cities = []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose"}
var states = []string{"NY", "CA", "IL", "TX", "AZ", "PA", "TX", "CA", "TX", "CA"}

// init functions are called automatically when the package is initialized,
// which happens after all package-level variables have been initialized and
// before the main function starts.
func init() {
	// Create a new random number generator and seed it
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner
func RandomOwner() string {
	return RandomString(6)
}

// RandomPhone
func RandomPhone() string {
	var sb strings.Builder

	k := len(numbers)

	for i := 0; i < 10; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomMoney
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GHC"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail
func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomFullName() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return firstName + " " + lastName
}

func RandomAddress() string {
    streetNumber := rand.Intn(9999) + 1
    street := streets[rand.Intn(len(streets))]
    city := cities[rand.Intn(len(cities))]
    state := states[rand.Intn(len(states))]
    zip := rand.Intn(89999) + 10000
    return fmt.Sprintf("%d %s, %s, %s %05d", streetNumber, street, city, state, zip)
}

func RandomStatus() string {
	status := []string{"SHIPPED", "DELIVERED", "PROCESSING"}
	n := len(status)
	return status[rand.Intn(n)]
}
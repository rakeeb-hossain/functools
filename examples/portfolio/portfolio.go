// This showcases several applications of functools helpers to a store of user objects and their associated
// portfolio holdings.
package main

import (
	"github.com/rakeeb-hossain/functools"
	"log"
)

type user struct {
	username 	 string
	age 		 int
	hasPortfolio bool
}

type holding struct {
	ticker 		string
	boughtTime 	int
	quantity	float64
	price		float64
}

type portfolio struct {
	holdings []holding
}

var (
	users = []user{
		{"gopher", 21, true},
		{"rakeeb", 20, false},
		{"jack", 22, true}}

	usersPortfolioMap = map[string]portfolio{
		"gopher": {[]holding{{"TSLA", 1639768692, 4.5, 1000}, {"ABNB", 1639163892, 2.5, 200}}},
		"jack":	  {[]holding{ {"BTC", 1512933492, 5, 1000}, {"ETH", 1639768692, 10, 100}} }}
)

func main() {
	// Count users with linked portfolios
	log.Printf("num users with linked portfolios: %d", functools.Count(users, func(u user) bool { return u.hasPortfolio }))

	// Print usernames of users with linked portfolios
	functools.ForEach(
		functools.Filter(users, func(u user) bool { return u.hasPortfolio }),
		func(u user) { log.Printf("%s has a linked portfolio\n", u.username) },
	)

	// For users with connected portfolios, get portfolio values
	usersWithPortfolio := functools.Filter(users, func(u user) bool { return u.hasPortfolio })
	userPortfolioValues := functools.Map(usersWithPortfolio, func(u user) float64 {
		return functools.Reduce(usersPortfolioMap[u.username].holdings, 0, func(accum float64, h holding) float64 {
			return accum + h.quantity*h.price
		})
	})

	for i, _ := range usersWithPortfolio {
		log.Printf("user %s has portfolio value %f\n", usersWithPortfolio[i].username, userPortfolioValues[i])
	}

	// Get total price of assets in all connected portfolios
	totalVal := functools.Sum(userPortfolioValues)
	log.Printf("total asset value: %f", totalVal)
}
package main

import (
	"fmt"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main() {
	// Set your secret key. Remember to switch to your live secret key in production.
	// See your keys here: https://dashboard.stripe.com/apikeys
	stripe.Key = "sk_test_51M15lZF8QHNbm1rwoEyiGyQ848a3nqxKg1JaLWAanzSQjXOyd7PIKu5S3x1lhzzdJ4dxEA6aB1S8CR2MECBKwV0L00B43k8cdk"

	c, _ := customer.Get("cus_MkbL824M2P30XC", nil)
	fmt.Println(c)

	fmt.Println("Hello, Authentication")
}

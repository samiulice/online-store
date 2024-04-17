package cards

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transactions struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

// Charge is the alias for CreatePaymentIntent. It's a more meaningful name
func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	//create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	//Add metadata to that transaction info
	// params.AddMetadata("key", "value ")

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessaeg(stripeErr.Code)
		}
		return nil, msg, err
	}
	return pi, "", nil
}

// cardErrorMessaeg returns a string msg that corresponds to a specific error code
func cardErrorMessaeg(code stripe.ErrorCode) string {
	var msg = ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"

	case stripe.ErrorCodeExpiredCard:
		msg = "Your card is expired"

	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect zip code"

	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount is to large to charge to your card"

	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is to small to charge to your card"

	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient banlance"

	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Your postal code is invalid"

	default:
		msg = "Your card was declined"
	}

	return msg
}

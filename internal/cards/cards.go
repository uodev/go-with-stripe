package cards

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	// params.AddMetadata("key", "value")

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "kart reddedildi"
	case stripe.ErrorCodeExpiredCard:
		msg = "kart tarihi gecmis"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "cvc kodu hatali"
	case stripe.ErrorCodeIncorrectZip:
		msg = "posta kodu hatali"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "cok buyuk sayi girdiniz"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "cok kucuk sayi girdiniz"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "yeteriz bakiye"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "posta kodu hatali"
	default:
		msg = "kart reddedildi"
	}

	return msg
}

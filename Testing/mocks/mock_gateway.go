package mocks


type PaymentGateway interface {
    DoSomething(num int) (int, error)
}


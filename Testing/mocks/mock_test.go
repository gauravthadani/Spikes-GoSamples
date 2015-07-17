package mocks
import (
  "testing"
  "github.com/stretchr/testify/mock"
)


type MyMockedObject struct{

  mock.Mock

}



func (m *MyMockedObject) DoSomething(number int) (int, error) {

  args := m.Called(number)
  return args.Int(0), args.Error(1)

}

func targetFuncThatDoesSomethingWithObj(p PaymentGateway){

    p.DoSomething(123)

}

func TestSomething(t *testing.T) {

  testObj := new(MyMockedObject)
  testObj.On("DoSomething", 123).Return(123, nil)
  targetFuncThatDoesSomethingWithObj(testObj)
  testObj.AssertExpectations(t)

}
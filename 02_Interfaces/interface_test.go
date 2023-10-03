package Interfaces

import (
	"fmt"
	"testing"
)

type Phone interface {
	call() string
}

type NokiaPhone struct {
}

func (nokiaPhone *NokiaPhone) call() string {
	return ("Here's Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone *IPhone) call() string {
	return ("Here's iPhone, I can call you!")
}

func PhoneCall(phone Phone) {
	fmt.Printf("%T:\t%v\n", phone, phone.call())
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = new(NokiaPhone)
	PhoneCall(phone)
	phone = &IPhone{}
	PhoneCall(phone)
}

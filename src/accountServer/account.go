package accountServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type Account struct {
}

func New() *Account {
	return &Account{}
}

func (this *Account) Init() {

}

func (this *Account) Type() akmessage.ServerType {
	return akmessage.ServerType_Account
}

func (this *Account) Run(d *in.Input) {

}

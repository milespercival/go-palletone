package ptn

import (
	"github.com/palletone/go-palletone/consensus/jury"
	"github.com/palletone/go-palletone/common/event"
	"github.com/palletone/go-palletone/dag/modules"
	"github.com/palletone/go-palletone/core/accounts/keystore"
	"github.com/palletone/go-palletone/common"
)

type contractInf interface {
	SubscribeContractEvent(ch chan<- jury.ContractExeEvent) event.Subscription
	ProcessContractEvent(event *jury.ContractExeEvent) error

	SubscribeContractSigEvent(ch chan<- jury.ContractSigEvent) event.Subscription
	ProcessContractSigEvent(event *jury.ContractSigEvent) error

	RunContractLoop(addr common.Address, ks *keystore.KeyStore) error
	CheckContractTxValid(tx *modules.Transaction) bool
}

func (self *ProtocolManager) contractExecRecvLoop() {
	for {
		select {
		case event := <-self.contractExecCh:
			go self.contractProc.ProcessContractEvent(&event)

		case <-self.contractExecSub.Err():
			return
		}
	}
}

func (self *ProtocolManager) contractSigRecvLoop() {
	for {
		select {
		case event := <-self.contractSigCh:
			go self.contractProc.ProcessContractSigEvent(&event)

		case <-self.contractSigSub.Err():
			return
		}
	}
}

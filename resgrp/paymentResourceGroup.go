package resgrp

import (
	iface "github.com/insanidade/ppsdk/interfaces"
	model "github.com/insanidade/ppsdk/model"
	txn "github.com/insanidade/ppsdk/transactions"
)

type PaymentResourceGroup struct {
}

func NewPaymentResourceGroup() *PaymentResourceGroup {
	return &PaymentResourceGroup{}
}

func (prg *PaymentResourceGroup) BuildPaymentTransaction() *txn.PaymentTransaction {
	return txn.NewPaymentTransaction(model.NewDefaultPaymentRequestContainer(), model.NewPaymentResponseContainer())
}
func (prg *PaymentResourceGroup) BuildGetPaymentDetailsTransaction(payid string) *txn.PaymentTransaction {
	return txn.NewPaymentTransaction(model.NewGetPaymentRequestContainer(payid), model.NewPaymentResponseContainer())
}

func (prg *PaymentResourceGroup) BuildCreateBATokenTransaction() *txn.CreateBATokenTransaction {
	return txn.NewCreateBATokenTransaction()
}

func (prg *PaymentResourceGroup) BuildCreateBAgreementTransaction() *txn.CreateBAgreementTransaction {
	return txn.NewCreateBAgreementTransaction()
}

func (prg *PaymentResourceGroup) BuildCalculateFinancingOptionsTransaction() *txn.PaymentTransaction {
	return txn.NewPaymentTransaction(model.NewDefaultCalculateFinancingOptionsRequestContainer(), model.NewCalculateFinancingOptionsResponseContainer())
}

func (prg *PaymentResourceGroup) BuildCapturePaymentRefWithInstallmentsTransaction() *txn.PaymentTransaction {
	return txn.NewPaymentTransaction(model.NewDefaultPaymentRequestContainer(), model.NewPaymentResponseContainer())
}

//#################################################################
// TODO: implement order transaction
func (prg *PaymentResourceGroup) BuildOrderTransaction() iface.Transaction {
	return nil
}

// TODO: implement authorize transaction
func (prg *PaymentResourceGroup) BuildAuthorizeTransaction() iface.Transaction {
	return nil
}

// ##################################################################
// ######################INTERFACE IMPLEMENTATIONS###################
// ##################################################################

// func (pb *PaymentBuilder) GetRequestContainer() iface.RequestContainer {
// 	pb.bodyFactory()
// 	return model.NewPaymentRequestContainer(pb.header,
// 		pb.body,
// 		paypalSandboxAPIURL+createPaymentURL)
// }
// func (pb *PaymentBuilder) BuildResponseContainer() iface.ResponseContainer {
// 	return nil
// }

// ##################################################################
// #####################END OF INTERFACE IMPLEMENTATIONS#############
// ##################################################################

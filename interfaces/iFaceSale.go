package interfaces

type IFaceSale interface {
	GetID() string
	GetAmount() *IFaceAmount
	GetPaymentMode() string
	GetState() string
	GetReasonCode() string
	GetProtectionEligibility() string
	GetProtectionEligibilityType() string
	GetClearingTime() string
	GetPaymentHoldStatus() string
	GetPaymentHoldReasons() []IFacePaymentHoldReason
	GetTransactionFee() *IFaceValueInfo
	GetReceivableAmount() *IFaceValueInfo
	GetExchangeRate() string
	GetFMFDetails() *IFaceFMFDetails
	GetReceiptID() string
	GetParentPayment() string
	GetProcessorResponse() *IFaceProcessorResponse
	GetBillingAgreementID() string
	GetCreateTime() string
	GetUpdateTime() string
	GetLinks() []IFaceLink

	SetID(ID string)
	SetAmount(a *IFaceAmount)
	SetPaymentMode(pm string)
	SetState(s string)
	SetReasonCode(rc string)
	SetProtectionEligibility(pe string)
	SetProtectionEligibilityType(pet string)
	SetClearingTime(ct string)
	SetPaymentHoldStatus(phs string)
	SetPaymentHoldReasons(phr []IFacePaymentHoldReason)
	SetTransactionFee(tf *IFaceValueInfo)
	SetReceivableAmount(ra *IFaceValueInfo)
	SetExchangeRate(er string)
	SetFMFDetails(fd *IFaceFMFDetails)
	SetReceiptID(rID string)
	SetParentPayment(pp string)
	SetProcessorResponse(pr *IFaceProcessorResponse)
	SetBillingAgreementID(baID string)
	SetCreateTime(ct string)
	SetUpdateTime(ut string)
	SetLinks(l []IFaceLink)
}

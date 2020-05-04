package interfaces

type IFaceAuthorization interface {
	GetID() string
	GetAmount() *IFaceAmount
	GetPaymentMode() string
	GetState() string
	GetReasonCode() string
	GetProtectionEligibility() string
	GetProtectionEligibilityType() string
	GetFMFDetails() *IFaceFMFDetails
	GetReceiptID() string
	GetParentPayment() string
	GetProcessorResponse() *IFaceProcessorResponse
	GetValidUntil() string
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
	SetFMFDetails(fmfd *IFaceFMFDetails)
	SetReceiptID(rID string)
	SetParentPayment(pp string)
	SetProcessorResponse(pr *IFaceProcessorResponse)
	SetValidUntil(vu string)
	SetCreateTime(ct string)
	SetUpdateTime(ut string)
	SetLinks(l []IFaceLink)
}

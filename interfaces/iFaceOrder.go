package interfaces

type IFaceOrder interface {
	GetID() string
	GetAmount() *IFaceAmount
	GetPaymentMode() string
	GetState() string
	GetReasonCode() string
	GetProtectionEligibility() string
	GetProtectionEligibilityType() string
	GetFMFDetails() *IFaceFMFDetails
	GetParentPayment() string
	GetCreateTime() string
	GetUpdateTime() string
	GetLinks() []IFaceLink
}

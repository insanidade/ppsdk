package interfaces

type IFaceProcessorResponse interface {
	GetResponseCode() string
	GetAVSCode() string
	GetCVVCode() string
	GetAdviceCode() string
	GetECISubmitted() string
	GetVPAS() string

	SetResponseCode(rc string)
	SetAVSCode(avs string)
	SetCVVCode(cvv string)
	SetAdviceCode(ac string)
	SetECISubmitted(eci string)
	SetVPAS(vpas string)
}

package model

//RelatedResource defines a related_resource json object
type RelatedResource struct {
	Sale          *Sale          `json:"sale,omitempty"`
	Authorization *Authorization `json:"authorization,omitempty"`
	Order         *Order         `json:"order,omitempty"`
	Capture       *Capture       `json:"capture,omitempty"`
	Refund        *Refund        `json:"refund,omitempty"`
}

//NewRelatedResource returns an instance of a RelatedResource object
func NewRelatedResource(sale *Sale,
	auth *Authorization,
	order *Order,
	capt *Capture,
	refund *Refund) *RelatedResource {
	return &RelatedResource{
		Sale:          sale,
		Authorization: auth,
		Order:         order,
		Capture:       capt,
		Refund:        refund,
	}
}

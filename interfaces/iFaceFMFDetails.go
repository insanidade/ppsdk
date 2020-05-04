package interfaces

type IFaceFMFDetails interface {
	GetFilterType() string
	GetFilterID() string
	GetName() string
	GetDescription() string

	SetFilterType(ft string)
	SetFilterID(fID string)
	SetName(n string)
	SetDescription(d string)
}

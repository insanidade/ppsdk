package model

type FMFDetails struct {
	FilterType  string `json:"filter_type,omitempty"`
	FilterID    string `json:"filter_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewDefaultFMFDetails() *FMFDetails {
	return &FMFDetails{}
}

func (fd *FMFDetails) GetDescription() string {
	return fd.Description
}

func (fd *FMFDetails) SetDescription(value string) {
	fd.Description = value
}
func (fd *FMFDetails) GetName() string {
	return fd.Name
}
func (fd *FMFDetails) SetName(value string) {
	fd.Name = value
}

func (fd *FMFDetails) GetFilterID() string {
	return fd.FilterID
}
func (fd *FMFDetails) SetFilterID(value string) {
	fd.FilterID = value
}

func (fd *FMFDetails) GetFilterType() string {
	return fd.FilterType
}
func (fd *FMFDetails) SetFilterType(value string) {
	fd.FilterType = value
}

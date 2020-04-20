package interfaces

//Invoice is inteface that must be implemented by anyone
//who wants to create his/her own invoice type.
type Invoice interface {
	String() string
}

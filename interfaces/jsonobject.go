package interfaces

//JSONObject represents a json object.
//It may contain other JsonObject's in it
type JSONObject interface {
	// SimpleElements() map[string]string
	AddObject(object JSONObject)
}

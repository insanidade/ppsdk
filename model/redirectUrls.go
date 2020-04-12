package model

type RedirectURLS struct {
	ReturnURL string `json:"return_url,omitempty"`
	CancelURL string `json:"cancel_url,omitempty"`
}

//NewRedirectURLS returns a new instance of RedirectURLS.
//r : the return url
//c: the cancel url
func NewRedirectURLS(r string, c string) *RedirectURLS {
	return &RedirectURLS{
		ReturnURL: r,
		CancelURL: c,
	}
}

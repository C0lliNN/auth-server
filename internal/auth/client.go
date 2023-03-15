// go:generate stringer -type=ClientType -linecomment

package auth

type ClientType int

const (
	Public ClientType = iota + 1
	Confidential
)

type Client struct {
	ID          string
	Type        ClientType
	Secret      *string
	RedirectURI *string
}

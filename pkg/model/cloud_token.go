package model

type CloudToken string

func (t CloudToken) String() string {
	return string(t)
}

package actions

type ActionErrType int

const (
	ErrConflict ActionErrType = 0
	ErrNotFound ActionErrType = 1
)

type ActionErr struct {
	Message string
	Type    ActionErrType
}

func (a ActionErr) Error() string {
	return a.Message
}

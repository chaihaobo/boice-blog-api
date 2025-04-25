package constant

const (
	ContextKeyTrx ContextKey = iota + 1
	ContextKeyUserID
	ContextKeyPassword
)

type ContextKey int

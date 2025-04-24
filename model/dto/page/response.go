package page

type (
	Response[T any] struct {
		Data  []T   `json:"data"`
		Total int64 `json:"total"`
	}
)

func NewResponse[T any](data []T, total int64) *Response[T] {
	return &Response[T]{
		Data:  data,
		Total: total,
	}
}

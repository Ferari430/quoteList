package payload

type Quote struct {
	Author string `json:"author" validate:"required"`
	Quote  string `json:"quote" validate:"required"`
	Id     uint64 `json:"-"`
}

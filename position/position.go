package position

type Position string

const (
	PositionTop    Position = "top"
	PositionBottom Position = "bottom"
	PositionLeft   Position = "left"
	PositionRight  Position = "right"
)

func (p Position) String() string {
	return string(p)
}

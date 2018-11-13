package drawer

type Drawer struct {
	Body   string
	Offset int
}

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func New(body string, offset int) *Drawer {
	return &Drawer{
		Body:   body,
		Offset: offset,
	}
}

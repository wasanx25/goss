package drawer

type Drawer struct {
	Body   string
	Offset int
}

func New(body string, offset int) *Drawer {
	return &Drawer{
		Body:   body,
		Offset: offset,
	}
}

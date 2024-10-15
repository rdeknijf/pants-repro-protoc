package repro_protoc

type Mother struct {
	Name string
}

func NewMother(name string) *Mother {
	return &Mother{
		Name: name,
	}
}

package checker

type BodyTree struct {
	Name   string
	Childs []BodyTree
}

func (mt BodyTree) GetAll() []string {
	return mt.getAllInternal(make([]string, 0), "")
}

func (mt BodyTree) getAllInternal(res []string, prefix string) []string {
	name := mt.Name

	if prefix != "" {
		name = prefix + " " + name
	}

	res = append(res, name)

	for _, ch := range mt.Childs {
		res = ch.getAllInternal(res, name)
	}

	return res
}

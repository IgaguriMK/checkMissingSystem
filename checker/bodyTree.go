package checker

type BodyTree struct {
	Name   string
	Childs []BodyTree
}

func (mt BodyTree) GetAll(prefix string) []string {
	return mt.getAllInternal(make([]string, 0), prefix)
}

func (mt BodyTree) getAllInternal(res []string, prefix string) []string {
	fullname := prefix

	if mt.Name != "" {
		fullname = fullname + " " + mt.Name
	}

	res = append(res, fullname)

	for _, ch := range mt.Childs {
		res = ch.getAllInternal(res, fullname)
	}

	return res
}

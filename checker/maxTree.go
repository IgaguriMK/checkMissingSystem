package checker

type MaxTree struct {
	Name   string
	Childs []MaxTree
}

func (mt MaxTree) GetAll(prefix string) []string {
	return mt.getAllInternal(make([]string, 0), prefix)
}

func (mt MaxTree) getAllInternal(res []string, prefix string) []string {
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

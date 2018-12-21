package cache

type Mem struct {
	values map[string]string
}

func NewMem() *Mem {
	return &Mem{}
}

package prototype

//Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Register(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

func (p *PrototypeManager) Create(name string) Cloneable {
	obj := p.prototypes[name]
	return obj.Clone()
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) Clone() *Student {
	t := *s
	return &t
}

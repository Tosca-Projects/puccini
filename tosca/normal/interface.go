package normal

//
// Interface
//

type Interface struct {
	NodeTemplate *NodeTemplate `json:"-" yaml:"-"`
	Group        *Group        `json:"-" yaml:"-"`
	Relationship *Relationship `json:"-" yaml:"-"`
	Name         string        `json:"-" yaml:"-"`

	Description   string         `json:"description" yaml:"description"`
	Types         Types          `json:"types" yaml:"types"`
	Inputs        Constrainables `json:"inputs" yaml:"inputs"`
	Operations    Operations     `json:"operations" yaml:"operations"`
	Notifications Notifications  `json:"notifications" yaml:"notifications"`
}

func (self *NodeTemplate) NewInterface(name string) *Interface {
	intr := &Interface{
		NodeTemplate:  self,
		Name:          name,
		Types:         make(Types),
		Inputs:        make(Constrainables),
		Operations:    make(Operations),
		Notifications: make(Notifications),
	}
	self.Interfaces[name] = intr
	return intr
}

func (self *Group) NewInterface(name string) *Interface {
	intr := &Interface{
		Group:         self,
		Name:          name,
		Types:         make(Types),
		Inputs:        make(Constrainables),
		Operations:    make(Operations),
		Notifications: make(Notifications),
	}
	self.Interfaces[name] = intr
	return intr
}

func (self *Relationship) NewInterface(name string) *Interface {
	intr := &Interface{
		Relationship:  self,
		Name:          name,
		Types:         make(Types),
		Inputs:        make(Constrainables),
		Operations:    make(Operations),
		Notifications: make(Notifications),
	}
	self.Interfaces[name] = intr
	return intr
}

//
// Interfaces
//

type Interfaces map[string]*Interface

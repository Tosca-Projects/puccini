package parser

import (
	"fmt"
	"sync"

	"github.com/tliron/puccini/format"
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/problems"
)

type Context struct {
	Root    *Unit
	Quirks  []string
	Units   Units
	Parsing sync.Map
	WG      sync.WaitGroup
	Locker  sync.Mutex
}

func NewContext(quirks []string) Context {
	return Context{
		Quirks: quirks,
	}
}

func (self *Context) GetProblems() *problems.Problems {
	return self.Root.GetContext().Problems
}

func (self *Context) AddUnit(entityPtr interface{}, container *Unit, nameTransformer tosca.NameTransformer) *Unit {
	unit := NewUnit(entityPtr, container, nameTransformer)

	self.Locker.Lock()
	self.Units = append(self.Units, unit)
	self.Locker.Unlock()

	if container != nil {
		// Merge problems into container
		// Note: This happens every time the same unit is imported,
		// so it could be that that problems are merged multiple times,
		// but because problems are de-duped, everything is OK :)
		container.GetContext().Problems.Merge(unit.GetContext().Problems)
	}

	self.goReadImports(unit)

	return unit
}

// Print

func (self *Context) PrintImports(indent int) {
	format.PrintIndent(indent)
	fmt.Fprintf(format.Stdout, "%s\n", format.ColorValue(self.Root.GetContext().URL.String()))
	self.Root.PrintImports(indent, format.TreePrefix{})
}

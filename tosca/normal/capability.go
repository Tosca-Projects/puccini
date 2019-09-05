package normal

import (
	"encoding/json"
	"math"
)

//
// Capability
//

type Capability struct {
	NodeTemplate *NodeTemplate
	Name         string

	Description          string
	Types                Types
	Properties           Constrainables
	Attributes           Constrainables
	MinRelationshipCount uint64
	MaxRelationshipCount uint64
}

func (self *NodeTemplate) NewCapability(name string) *Capability {
	capability := &Capability{
		NodeTemplate:         self,
		Name:                 name,
		Types:                make(Types),
		Properties:           make(Constrainables),
		Attributes:           make(Constrainables),
		MaxRelationshipCount: math.MaxUint64,
	}
	self.Capabilities[name] = capability
	return capability
}

func (self *Capability) Marshalable() interface{} {
	var minRelationshipCount int64
	var maxRelationshipCount int64
	minRelationshipCount = int64(self.MinRelationshipCount)
	if self.MaxRelationshipCount == math.MaxUint64 {
		maxRelationshipCount = -1
	} else {
		maxRelationshipCount = int64(self.MaxRelationshipCount)
	}

	return &struct {
		Description          string         `json:"description" yaml:"description"`
		Types                Types          `json:"types" yaml:"types"`
		Properties           Constrainables `json:"properties" yaml:"properties"`
		Attributes           Constrainables `json:"attributes" yaml:"attributes"`
		MinRelationshipCount int64          `json:"minRelationshipCount" yaml:"minRelationshipCount"`
		MaxRelationshipCount int64          `json:"maxRelationshipCount" yaml:"maxRelationshipCount"`
	}{
		Description:          self.Description,
		Types:                self.Types,
		Properties:           self.Properties,
		Attributes:           self.Attributes,
		MinRelationshipCount: minRelationshipCount,
		MaxRelationshipCount: maxRelationshipCount,
	}
}

// json.Marshaler interface
func (self *Capability) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Marshalable())
}

// yaml.Marshaler interface
func (self *Capability) MarshalYAML() (interface{}, error) {
	return self.Marshalable(), nil
}

//
// Capabilities
//

type Capabilities map[string]*Capability

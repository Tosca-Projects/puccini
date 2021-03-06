package tosca_v1_3

import (
	"strings"

	"github.com/tliron/puccini/tosca"
)

//
// Metadata
//

type Metadata map[string]string

// tosca.Reader signature
func ReadMetadata(context *tosca.Context) tosca.EntityPtr {
	var self Metadata

	if context.Is("!!map") {
		metadata := context.ReadStringStringMap()
		if metadata != nil {
			self = *metadata
		}
	}

	if self != nil {
		for key, value := range self {
			if strings.HasPrefix(key, "puccini.scriptlet.import:") {
				context.ImportScriptlet(key[25:], value)
				delete(self, key)
			} else if strings.HasPrefix(key, "puccini.scriptlet:") {
				context.EmbedScriptlet(key[18:], value)
				delete(self, key)
			}
		}
	}

	return self
}

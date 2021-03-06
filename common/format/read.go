package format

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/tliron/puccini/ard"
	"github.com/tliron/yamlkeys"
)

func Read(reader io.Reader, format string) (ard.Value, error) {
	switch format {
	case "yaml", "":
		return ReadYAML(reader)
	case "json":
		return ReadJSON(reader)
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

func ReadYAML(reader io.Reader) (ard.Value, error) {
	return yamlkeys.Decode(reader)
}

func ReadAllYAML(reader io.Reader) (ard.List, error) {
	return yamlkeys.DecodeAll(reader)
}

func ReadJSON(reader io.Reader) (ard.Value, error) {
	var data ard.Value
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err == nil {
		data, _ = ard.ToMaps(data)
		return data, nil
	} else {
		return nil, err
	}
}

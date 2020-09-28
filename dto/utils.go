package dto

import "encoding/json"

// Mapper struct
type Mapper struct{}

// ModelToEntity func
// converte dto em uma entidade
func (Mapper) ModelToEntity(m, e interface{}) error {
	obj, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(obj, e)
}

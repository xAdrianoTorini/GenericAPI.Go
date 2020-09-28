package entities

type queryGeneric string

// GenericEntity - Entidade Bando Relacoinal
type GenericEntity struct {
	Paramns string `db:"Paramn" json:"Paramn"`
}

//QuerySelect recupera os dados dos Veiuculos
func (queryGeneric) QuerySelect() (str string) {
	str = ``

	return str
}

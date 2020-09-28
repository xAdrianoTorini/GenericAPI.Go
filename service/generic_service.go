package service

import (
	"log"

	"Inlog.Go.Service.Api/entities"
	"Inlog.Go.Service.Api/infraestructure/repository"
)

//Service Irta Persistir os Servicos.
var Service service

type service string

//ConsultarDados Retorna os veiculos do banco de dados
func (service) ConsultarDados() (v []entities.GenericEntity, err error) {
	rep := repository.Repository{}
	err = rep.Query(&v, entities.Query.CustomQuery.QuerySelect())
	if err != nil {
		log.Fatal(err)
	}
	return
}

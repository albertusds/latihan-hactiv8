package services

import (
	"net/http"
	"project-sesi8/models"
	"project-sesi8/params"
	"project-sesi8/repositories"
)

type PersonService struct {
	personRepo repositories.PersonRepo
}

func NewPersonService(repo repositories.PersonRepo) *PersonService {
	return &PersonService{
		personRepo: repo,
	}
}

func (p *PersonService) CreatePerson(request params.CreatePerson) *params.Response {
	model := models.Person{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	err := p.personRepo.CreatePerson(&model)
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "CREATE SUCCESS",
		Payload: request,
	}
}

func (p *PersonService) GetAllPerson() *params.Response {
	response, err := p.personRepo.GetAllPersons()
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}
}

func (p *PersonService) GetPersonById(id int) *params.Response {
	response, err := p.personRepo.GetPersonById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "DATA NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}
}

func (p *PersonService) DeletePersonById(id int) *params.Response {
	response, err := p.personRepo.DeletePersonById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "DATA NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}
}

func (p *PersonService) UpdatePersonById(req *params.UpdatePerson) *params.Response {

	var modelBase = models.Model{
		ID: req.ID,
	}
	var modelPerson = models.Person{
		Model:     modelBase,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	response, err := p.personRepo.UpdatePersonById(&modelPerson)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "DATA NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}
}

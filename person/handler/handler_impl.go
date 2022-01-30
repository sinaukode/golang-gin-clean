package handler

import (
	"github.com/gin-gonic/gin"
	"go-clean-code-gin/entity"
	"go-clean-code-gin/person"
	"net/http"
	"strconv"
)

type PersonHandler struct {
	personService person.PersonService
}

func CreatePersonHandler(r *gin.Engine, personService person.PersonService) {
	personHandler := PersonHandler{personService}

	r.POST("/person", personHandler.createPerson)
	r.GET("/person", personHandler.getPerson)
	r.GET("/person/:id", personHandler.getPersonById)
	r.PUT("/person/:id", personHandler.updatePerson)
	r.DELETE("/person/:id", personHandler.deletePerson)

}

func (p *PersonHandler) createPerson(c *gin.Context) {
	var person = entity.Person{}
	err := c.Bind(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	result, err := p.personService.Create(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}

func (p *PersonHandler) getPerson(c *gin.Context) {
	persons, err := p.personService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": persons,
	})

}

func (p *PersonHandler) getPersonById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	person, err := p.personService.GetById(id)

	c.JSON(http.StatusOK, gin.H{
		"result": person,
	})

}

func (p *PersonHandler) updatePerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	//cek data id
	_, err = p.personService.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Person ID Not Found",
		})
		return
	}

	//bind data Person by Id
	var dataPerson = entity.Person{}
	err = c.Bind(&dataPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	//update data
	result, err := p.personService.Update(id, &dataPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (p *PersonHandler) deletePerson(c *gin.Context) {
	//var person = entity.Person{}
	idStr := c.Param("id")
	id, err := checkId(idStr, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete Id Not Found",
		})
		return
	}

	err = p.personService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete Success",
	})
}

func checkId(id string, p *PersonHandler) (int, error) {
	paramId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	_, err = p.personService.GetById(paramId)
	if err != nil {
		return 0, err
	}
	return paramId, nil
}

package api

import (
	"go-study/internal/model"
	"go-study/internal/model/response"

	"github.com/gin-gonic/gin"
)

func GetSolutionList(c *gin.Context) {
	res, err := solutionService.GetSolutionList(c)
	if err != nil {
		response.Fail(500, err.Error(), c)
		return
	}
	response.OK(res, c)

}

func AddSolution(c *gin.Context) {
	var newSolution model.AddSolutionReq
	if c.ShouldBind(&newSolution) != nil {
		response.Fail(400, "Bad Request", c)
		return
	}
	res, err := solutionService.AddSolution(newSolution)
	if err != nil {
		response.Fail(500, err.Error(), c)
		return
	}
	response.OK(res, c)
}

func UpdateSolution(c *gin.Context) {
	var updateSolution model.StoreSolution
	if c.ShouldBind(&updateSolution) != nil {
		response.Fail(400, "Bad Request", c)
		return
	}
	res, err := solutionService.UpdateSolution(updateSolution)
	if err != nil {
		response.Fail(500, err.Error(), c)
		return
	}
	response.OK(res, c)
}

func DeleteSolution(c *gin.Context) {
	var deleteSolution model.StoreSolution
	if c.ShouldBind(&deleteSolution) != nil {
		response.Fail(400, "Bad Request", c)
		return
	}
	err := solutionService.DeleteSolution(deleteSolution.Id)
	if err != nil {
		c.String(500, err.Error())
	}
	response.OK(nil, c)
}

package controller

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/service"
)

type DormController interface {
	GetDorms(context *gin.Context)
	GetDormDetail(context *gin.Context)
	GetDormSuggestions(context *gin.Context)
	GetAllDormFacilities(context *gin.Context)
	GetDormZones(context *gin.Context)
}

func DormControllerHandler(dormService service.DormService) DormController {
	return &dormController{
		dormService: dormService,
	}
}

type dormController struct {
	dormService service.DormService
}

// @Summary get list of dorms
// @Description returns list of dorms filter by dorm type, zone,capacity, location, dorm and room facilities
// @Tags dorm
// @Produce json
// @Accept json
// @Success 200 {array} dto.DormDTO "OK"
// @Param dorm_filter query dto.DormFilterDTO false "Dorm Filter"
// @Router /dorm [get]
func (dormController *dormController) GetDorms(context *gin.Context) {
	var request dto.DormFilterDTO

	//TODO: fix check ignore case

	if err := context.Bind(&request); err != nil {
		if err != io.EOF {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
	}

	dormDTOs := dormController.dormService.GetDorms(request)

	context.IndentedJSON(http.StatusOK, dormDTOs)
}

// @Summary get dorm details from dorm id
// @Description returns dorm details
// @Tags dorm
// @Produce json
// @Success 200 {object} model.Dorm "OK"
// @Failure 400,404 {object} dto.ErrorResponse
// @Param id path int true "Dorm ID"
// @Router /dorm/{id} [get]
func (dormController *dormController) GetDormDetail(context *gin.Context) {
	dormID := context.Param("id")

	if _, err := strconv.Atoi(dormID); err != nil {
		context.IndentedJSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "id is not an integer"})

		return
	}

	dorm, err := dormController.dormService.GetDorm(dormID)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, &dto.ErrorResponse{Message: err.Error()})

		return
	}

	context.IndentedJSON(http.StatusOK, dorm)
}

// @Summary get dorm names from first letter
// @Description returns list of dorm names from first letter
// @Tags dorm
// @Produce json
// @Success 200 {array} dto.DormSuggestionDTO "OK"
// @Param letter path string true "First Letter"
// @Router /dorm/suggest/{letter} [get]
func (dormController *dormController) GetDormSuggestions(context *gin.Context) {
	letter := context.Param("letter")

	dormSuggestionDTOs := dormController.dormService.GetDormSuggestions(letter)

	context.IndentedJSON(http.StatusOK, dormSuggestionDTOs)
}

// @Summary get dorm facilities
// @Description returns list of dorm facilities
// @Tags dorm
// @Produce json
// @Success 200 {array} string "OK"
// @Router /dorm/facility [get]
func (dormController *dormController) GetAllDormFacilities(context *gin.Context) {
	dormFacilities := dormController.dormService.GetAllDormFacilities()

	context.IndentedJSON(http.StatusOK, dormFacilities)
}

// @Summary get dorm zones
// @Description returns list of dorm zones
// @Tags dorm
// @Produce json
// @Success 200 {array} string "OK"
// @Router /dorm/zone [get]
func (dormController *dormController) GetDormZones(context *gin.Context) {
	dormZones := dormController.dormService.GetDormZones()

	context.IndentedJSON(http.StatusOK, dormZones)
}

package handler

import (
	"fmt"
	"meli_golang_gin_basic_app/cmd/internal/dto/responsedto"
	"meli_golang_gin_basic_app/cmd/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Scan struct {
	scanUsecase usecase.IScanUsecase
}

func NewScanHandler(scanUsecase usecase.IScanUsecase) *Scan {
	return &Scan{
		scanUsecase: scanUsecase,
	}
}
// Scan
// @Summary      scan of Repository
// @Description  get jobId of scan process successfull
// @Accept       json
// @Produce      json
// @Success      200  {object}  responsedto.ResponseBody
// @Failure      400  {object}  errors.ErrorAPI
// @Failure      404  {object}  errors.ErrorAPI
// @Failure      500  {object}  errors.ErrorAPI
// @Router       /scan [post]
// FUNCION PARA REALIZAR ESCANEO DE ARCHIVOS EN UN REPOSITORIO GITHUB
func (s *Scan) Scan(c *gin.Context)  {

	var requestBody struct {
		URL string `json:"url"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewBadRequestApiError("Formato de JSON no válido"))
		return
	}
	repoURL := requestBody.URL
	if repoURL == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, NewBadRequestApiError("URL Not Found"))
		return
	}	

	res, err := s.scanUsecase.Scan(repoURL)
	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	response := responsedto.ResponseBody{
		StatusCode: http.StatusCreated,
		Body: res,
	}
	c.JSON(http.StatusOK, response)
}
// GetResult
// @Summary      Busca registros por JobId
// @Description  Obtiene los registros de escaneo por JobId
// @Accept       json
// @Produce      json
// @Param        jobid  JobId del escaneo
// @Success      200  {object}  responsedto.ResponseBody
// @Failure      400  {object}  errors.ErrorAPI
// @Failure      404  {object}  errors.ErrorAPI
// @Failure      500  {object}  errors.ErrorAPI
// @Router       /result/:job_id [get]
//  FUNCION PARA TRAER LOS RESULTADOS PREVIOS DE UN ESCANEO 
func (s *Scan) GetResult(c *gin.Context) {
	jobId := c.Param("job_id")
	if jobId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewBadRequestApiError("Job_ID Not Found"))
		return
	}
	resDB, err := s.scanUsecase.GetResult(jobId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(resDB) == 0{
		c.JSON(http.StatusOK, DataNotFoundError(fmt.Sprintf("No hay registros para: %s", jobId)))
		return
	}
	//* Metodo para parsear un objeto Go en []byte
	/* jsonData, err := json.Marshal(resDB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewJSONSerializeError("JSON Serialize Error"))
		return
	} */
	response := responsedto.ResponseBody{
		StatusCode: http.StatusOK,
		Body: resDB,
	}
	c.JSON(http.StatusOK, response)
}



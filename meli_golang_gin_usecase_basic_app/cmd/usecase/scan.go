package usecase

import (
	"fmt"
	"math/rand"
	"meli_golang_gin_basic_app/cmd/internal/dto"
	"meli_golang_gin_basic_app/cmd/internal/repository"
	"meli_golang_gin_basic_app/cmd/internal/domain"
	"time"
	//"meli_golang_gin_basic_app/cmd/internal/dto"
	"regexp"
	"strconv"
	"strings"
)

type IScanUsecase interface {
	Scan(repoURL string) (string, error)
	GetResult(jobId string) ([]domain.Entity, error)
	AnalisisFileContent(filePath string , fileContent string) (SensitiveInfo)
}

type ScanUsecase struct {
	dataRepository    repository.IDataRepository
	storageRepository		  repository.IStorageDB
}

type SensitiveInfo struct{
	FileName string 
	TypeDetectionsInfo string 
	LinesNumber string
	AmountDetections int
}
func NewScanUsecase(dataRepository repository.IDataRepository, storageRepository repository.IStorageDB ) *ScanUsecase {
	return &ScanUsecase{
		dataRepository:    dataRepository,
		storageRepository: storageRepository,
	}
}

// ************** FUNCION PARA REALIZAR EL ESCANEO EN BUSCA DE VULNERABILIDADES *************** 	|
func (s *ScanUsecase) Scan(repoURL string) (string, error) {
	files, owner, _ := s.dataRepository.GetFiles(repoURL)
	jobId := rand.Intn(999999)
	var responseDB string
	var errDb error
	
	userAndRepo := strings.Split(owner, "/")
	// ** Analiza cada archivo en busca de información sensible
	for _, file := range files {
		i := 0
		filePath := file.Path
		fileContent, err := s.dataRepository.GetDataFromFile(owner, filePath)
		if err != nil {
			return "", err
		}
		
		fileSensitiveData := s.AnalisisFileContent(filePath, fileContent)

		entity := dto.RepositoryDTO{
			RepositoryUrl: repoURL,
			File: file.Path,
			InformationType: fileSensitiveData.TypeDetectionsInfo,
			LinesNumber: fileSensitiveData.LinesNumber,
			DetectionDate: time.Now(),
			JobId: strconv.Itoa(jobId),
			RepositoryOwner: userAndRepo[0],
			AmountDetections: fileSensitiveData.AmountDetections,
		}
		// ************* Registrando en la DB ************** |
		responseDB, errDb = s.storageRepository.Create(entity)
		if errDb != nil {
			return "", errDb
		}

		if(i == 100){
			break
		}
		i++

	}
	fmt.Printf("Cantidad Files: %d\n", len(files))
	fmt.Printf("JobId: %s", responseDB)
    return responseDB, nil
}

// ********** FUNCION PARA ANALIZAR UN ARCHIVO ESPECIFICO EN BUSCA DE IP, CREDIT_CARD O EMAIL PATTERNS  ************* 	
func (s *ScanUsecase) AnalisisFileContent(filePath string , fileContent string) SensitiveInfo{

	ipv4Regex := `^(25[0-5]|2[0-4][0-9]|[0-1]?[0-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|[0-1]?[0-9]?[0-9])){3}$`
	creditCard := `^(4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$`
	email := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex1 := regexp.MustCompile(ipv4Regex)
	regex2 := regexp.MustCompile(creditCard)
	regex3 := regexp.MustCompile(email)

	var sensitiveData SensitiveInfo
	var amount_detections int
	var type_detections string = ""
	var lines_number string = ""
	lines := strings.Split(fileContent, "\n")
	for i, line := range lines {
		line = strings.Replace(line, "\t", "", -1)
		line = strings.Replace(line, " ", "", -1)
		if regex1.MatchString(line) {
			amount_detections++
			type_detections += " IP "
			lines_number += " " + strconv.Itoa(i + 1)
			
		}
		if regex2.MatchString(line) || strings.Contains(strings.ToLower(line), "credit_card") {
			amount_detections++
			type_detections += " CREDIT_CARD "
			lines_number += " " + strconv.Itoa(i + 1)
			
		}
		if regex3.MatchString(line) || strings.Contains(strings.ToLower(line), "email") {
			amount_detections++
			type_detections += " EMAIL "
			lines_number += " " + strconv.Itoa(i + 1)
			
		}
	}
	if amount_detections == 0 && type_detections == "" && lines_number == ""{
		sensitiveData = SensitiveInfo{FileName: filePath, LinesNumber: "N/A", TypeDetectionsInfo: "N/A", AmountDetections: 0}
	}

	lines_number = strings.Replace(lines_number, " ", "", 1)
	lines_number = strings.Replace(lines_number, " ", ",", -1)

	sensitiveData = SensitiveInfo{FileName: filePath, LinesNumber: lines_number, TypeDetectionsInfo: type_detections, AmountDetections: amount_detections}
	return sensitiveData
}

// ******* FUNCION PARA ENCONTRAR REGISTRO DE ESCANEOS ********** |
// *******     		PREVIOS POR MEDIO DEL JOB_ID 	   ********** |
func (s *ScanUsecase) GetResult(jobId string) ([]domain.Entity, error) {
	rerDb, err := s.storageRepository.FindByJobId(jobId)
	if err != nil {
		return nil, err
	}
	return rerDb, nil
}
package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IDataRepository interface {
	GetDataFromFile(owner string, filePath string) (string, error)
	GetFiles(reporURL string) ([]File, string, error)
	GetOwnerAndRepo(repoURL string) (string, error)
}

type DataRepository struct {
	Owner string 
	Repositorio string
}

func NewDataRepository() *DataRepository {
	return &DataRepository{}
}
// ************** FUNCION PARA OBTENER LA DATA DE UN ARCHIVO ESPECIFICO ****************** |
func (r *DataRepository) GetDataFromFile(owner string, filePath string) (string, error){
	apiURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", owner, filePath)
	response, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed file content. Status code: %d", response.StatusCode)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

type File struct{
    Path string `json:"path"`
}

// ************** FUNCION PARA OBTENER TODOS LOS ARCHIVOS DE LA RAMA PRINCIPAL *************** 	|
func (r *DataRepository) GetFiles(repoURL string) ([]File, string, error){

    owner, err := r.GetOwnerAndRepo(repoURL)
    if err != nil {
		return nil, "", err
	}
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/contents", owner)
	response, err := http.Get(apiURL)

	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("Failed to retrieve repository files. Status code: %d", response.StatusCode)
	}

	var files []File
	err = json.NewDecoder(response.Body).Decode(&files)
	if err != nil {
		return nil, "", err
	}

	return files, owner, nil
	
}
// ************** FUNCION PARA EXTRAER EL OWNER Y EL NOMBRE DEL REPOSITORIO *************** 	|
func (r *DataRepository) GetOwnerAndRepo(repoURL string) (string, error) {

    parts := strings.Split(repoURL, "/")
    if len(parts) != 5 {
        return "", fmt.Errorf("URL de repositorio no válida: %s", repoURL)
    }

    return fmt.Sprintf("%s/%s",parts[3], parts[4]), nil
}



package navigation

import (
	"encoding/json"
	"github.com/oktapascal/go-simpro/model"
	"io"
	"os"
)

type Repository struct{}

func (rpo *Repository) GetNavigation(group string) *[]model.Navigation {
	file, err := os.Open("storage/json/" + group + ".json")
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var navigation []model.Navigation
	err = json.Unmarshal(bytes, &navigation)
	if err != nil {
		panic(err)
	}

	return &navigation
}

package value_object

import (
	"errors"
	"fmt"
)

type ModelNumber struct {
	productCode string
	branch      string
	lot         string
}

func NewModelNumber(productCode, branch, lot string) ModelNumber {
	if len(productCode) == 0 {
		errors.New("product code cannot be empty")
	}
	if len(branch) == 0 {
		errors.New("branch cannot be empty")
	}
	if len(lot) == 0 {
		errors.New("lot cannot be empty")
	}
	return ModelNumber{productCode: productCode, branch: branch, lot: lot}
}
func (m ModelNumber) GetModelNumber() string {
	return m.productCode + "-" + m.branch + "-" + m.lot
}

func main() {
	modelNumber := NewModelNumber("a20421", "100", "1")
	fmt.Println(modelNumber.GetModelNumber())
}

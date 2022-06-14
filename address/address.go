package address

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const (
	empSheetName = "empaddress"
	filePath     = "./address/EmployeeAddress.xlsx"
)

// addressInter : Private Interface
type addressInter interface {
	AddEmployeeAddress()
	UpdateEmployeeAddress()
	GetEmployeeAddress()
	GetAllEmployeesAddress()
}

func GetAddress() *Address {
	return &Address{}
}

// Interface method implementation
func (a *Address) AddEmployeeAddress(add Address) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("error in reading")
	}
	defer file.Close()
	rows, err := file.GetRows(empSheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	axis := 0
	if len(rows) > 0 {
		axis = len(rows) + 1
	}

	file.SetCellValue(empSheetName, "A"+strconv.Itoa(axis), add.EmpID)
	file.SetCellValue(empSheetName, "B"+strconv.Itoa(axis), add.AddressLine1)
	file.SetCellValue(empSheetName, "C"+strconv.Itoa(axis), add.AddressLine2)
	file.SetCellValue(empSheetName, "D"+strconv.Itoa(axis), add.Street_Name)
	file.SetCellValue(empSheetName, "E"+strconv.Itoa(axis), add.City_Name)
	file.SetCellValue(empSheetName, "F"+strconv.Itoa(axis), add.State_Name)
	file.SetCellValue(empSheetName, "G"+strconv.Itoa(axis), add.Pincode)

	if err := file.Save(); err != nil {
		fmt.Println(err)
	}
}

func (a *Address) UpdateEmployeeAddress(add Address) {
	empAdd := a.GetAllEmployeesAddress()
	file, _ := fileOpen(filePath)

	got := false
	if len(empAdd) > 0 {
		for k, v := range empAdd {
			for _, v1 := range v {
				if v1 == add.EmpID {
					file.SetCellValue(empSheetName, "B"+strconv.Itoa(k+1), add.AddressLine1)
					file.SetCellValue(empSheetName, "C"+strconv.Itoa(k+1), add.AddressLine2)
					file.SetCellValue(empSheetName, "D"+strconv.Itoa(k+1), add.Street_Name)
					file.SetCellValue(empSheetName, "E"+strconv.Itoa(k+1), add.City_Name)
					file.SetCellValue(empSheetName, "F"+strconv.Itoa(k+1), add.State_Name)
					file.SetCellValue(empSheetName, "G"+strconv.Itoa(k+1), add.Pincode)
					if err := file.Save(); err != nil {
						fmt.Println("error in saving", err)
					}
					got = true
				}
				if got == true {
					break
				}
			}
			if got == true {
				break
			}
		}
	} else {
		fmt.Println("No such record under this empid")
	}
}

func (a *Address) GetEmployeeAddress(empId string) []string {
	fmt.Println("I am inside GetEmployeeAddress func")
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("error in reading")
	}
	defer file.Close()
	rows, err := file.GetRows(empSheetName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, v := range rows {
		for _, data := range v {
			if data == empId {
				return v
			}
		}
	}
	return nil
}

func (a *Address) GetAllEmployeesAddress() [][]string {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("error in reading")
	}
	defer file.Close()
	rows, err := file.GetRows(empSheetName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rows
}

func fileOpen(filename string) (*excelize.File, error) {
	file, err := excelize.OpenFile(filename)
	return file, err
}

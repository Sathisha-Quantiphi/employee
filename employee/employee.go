package employee

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const empSheetName string = "empData"
const filePath string = "./employee/EmployeeDetail.xlsx"

// employeeInter: Private Interface
type employeeInter interface {
	AddEmployee()
	UpdateEmployee()
	GetEmployeeDetail()
	GetAllEmployees()
}

func GetEmployee() *Employee {
	return &Employee{}
}

// Interface method implementation
func (e *Employee) AddEmployee(emp Employee) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	rows, err := file.GetRows(empSheetName)
	if err != nil {
		fmt.Println("GETROWS ERROR", err)
		return
	}
	axis := 0
	if len(rows) > 0 {
		axis = len(rows) + 1
	}

	file.SetCellValue(empSheetName, "A"+strconv.Itoa(axis), emp.ID)
	file.SetCellValue(empSheetName, "B"+strconv.Itoa(axis), emp.Name)
	file.SetCellValue(empSheetName, "C"+strconv.Itoa(axis), emp.Age)
	file.SetCellValue(empSheetName, "D"+strconv.Itoa(axis), emp.Sex)
	file.SetCellValue(empSheetName, "E"+strconv.Itoa(axis), emp.Designation)
	file.SetCellValue(empSheetName, "F"+strconv.Itoa(axis), emp.EmailID)

	if err := file.Save(); err != nil {
		fmt.Println("error in saving", err)
	}
}

func (e *Employee) UpdateEmployee(emp Employee) {
	empdata := e.GetAllEmployees()
	file, _ := fileOpen(filePath)
	got := false
	if len(empdata) > 0 {
		for k, v := range empdata {
			for _, v1 := range v {
				if v1 == emp.ID {
					file.SetCellValue(empSheetName, "E"+strconv.Itoa(k+1), emp.Designation)
					file.SetCellValue(empSheetName, "F"+strconv.Itoa(k+1), emp.EmailID)
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

func (e *Employee) GetAllEmployees() [][]string {
	file, openerr := fileOpen(filePath)
	if openerr != nil {
		fmt.Println(openerr)
		return nil
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

func (e *Employee) GetEmployeeDetail(empId string) []string {
	file, openerr := fileOpen(filePath)
	if openerr != nil {
		fmt.Println(openerr)
		return nil
	}
	defer file.Close()
	rows, err := file.GetRows(empSheetName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, row := range rows {
		for _, cell := range row {
			if cell == empId {
				return row
			}
		}
	}
	return nil
}

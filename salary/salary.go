package salary

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const (
	empSheetName = "empSalary"
	filePath     = "./salary/EmployeeSalary.xlsx"
)

// salaryInter: Private Interface
type salaryInter interface {
	AddEmployeeSalary()
	UpdateEmployeeSalary()
	GetEmployeeSalary()
	GetAllEmployeesSalary()
}

func GetSalary() *Salary {
	return &Salary{}
}

// Interface method implementation
func (d *Salary) AddEmployeeSalary(salry Salary) {

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

	file.SetCellValue(empSheetName, "A"+strconv.Itoa(axis), salry.EmpId)
	file.SetCellValue(empSheetName, "B"+strconv.Itoa(axis), salry.BSalary)
	file.SetCellValue(empSheetName, "C"+strconv.Itoa(axis), salry.Cross_Pay)

	if err := file.Save(); err != nil {
		fmt.Println(err)
	}
}

func (d *Salary) UpdateEmployeeSalary(salry Salary) {
	empSlry := d.GetAllEmployeesSalary()
	file, _ := fileOpen(filePath)

	got := false
	if len(empSlry) > 0 {
		for k, v := range empSlry {
			for _, v1 := range v {
				if v1 == salry.EmpId {
					file.SetCellValue(empSheetName, "B"+strconv.Itoa(k+1), salry.BSalary)
					file.SetCellValue(empSheetName, "C"+strconv.Itoa(k+1), salry.Cross_Pay)
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

func (d *Salary) GetEmployeeSalary(empId string) []string {
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

func (d *Salary) GetAllEmployeesSalary() [][]string {
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

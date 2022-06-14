package department

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const (
	empSheetName = "empdepartment"
	filePath     = "./department/EmployeeDepartment.xlsx"
)

// departmentInter: Private Interface
type departmentInter interface {
	AddEmployeeDepartment()
	UpdateEmployeeDepartment()
	GetEmployeeDepartment()
	GetAllEmployeeDepartments()
}

func GetDepartment() *Department {
	return &Department{}
}

// Interface method implementation
func (d *Department) AddEmployeeDepartment(dept Department) {

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

	file.SetCellValue(empSheetName, "A"+strconv.Itoa(axis), dept.EmpID)
	file.SetCellValue(empSheetName, "B"+strconv.Itoa(axis), dept.DeptID)
	file.SetCellValue(empSheetName, "C"+strconv.Itoa(axis), dept.DeptName)

	if err := file.Save(); err != nil {
		fmt.Println(err)
	}
}

func (d *Department) UpdateEmployeeDepartment(dept Department) {
	empDept := d.GetAllEmployeeDepartments()
	file, _ := fileOpen(filePath)

	got := false
	if len(empDept) > 0 {
		for k, v := range empDept {
			for _, v1 := range v {
				if v1 == dept.EmpID {
					file.SetCellValue(empSheetName, "B"+strconv.Itoa(k+1), dept.DeptID)
					file.SetCellValue(empSheetName, "C"+strconv.Itoa(k+1), dept.DeptName)
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

func (d *Department) GetAllEmployeeDepartments() [][]string {
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

func (d *Department) GetEmployeeDepartment(empId string) []string {
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

func fileOpen(filename string) (*excelize.File, error) {
	file, err := excelize.OpenFile(filename)
	return file, err
}

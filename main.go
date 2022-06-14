package main

import (
	"bufio"
	"employeemanagement/address"
	"employeemanagement/department"
	"employeemanagement/employee"
	"employeemanagement/salary"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("I am inside Main")

	var empdata employee.Employee
	var addressdata address.Address
	var depart department.Department
	var salarydata salary.Salary

	emp := employee.GetEmployee()
	addres := address.GetAddress()
	dept := department.GetDepartment()
	slry := salary.GetSalary()

	for {
		fmt.Println("\n 1. Add Employee Details")
		fmt.Println("\n 2. Get Employee Details")
		fmt.Println("\n 3. Get All Employee Details")
		fmt.Println("\n 4. Update Employee Details")
		fmt.Println("\n Please choose any option")
		input := readCLIInput()
		out, _ := strconv.Atoi(input)
		// fmt.Println("\n out", out, "type", reflect.TypeOf(out))
		switch out {
		case 1:
			fmt.Print("Enter Emp Name: ")
			empdata.Name = readCLIInput()

			fmt.Print("Enter Emp ID: ")
			empdata.ID = readCLIInput()
			addressdata.EmpID = empdata.ID
			depart.EmpID = empdata.ID
			salarydata.EmpId = empdata.ID

			fmt.Print("Enter Emp Age: ")
			empdata.Age = readCLIInput()

			fmt.Print("Enter Emp Sex: ")
			empdata.Sex = readCLIInput()

			fmt.Print("Enter Emp Designation: ")
			empdata.Designation = readCLIInput()

			fmt.Print("Enter Emp EmailID: ")
			empdata.EmailID = readCLIInput()

			fmt.Print("Enter Emp AddressLine1: ")
			addressdata.AddressLine1 = readCLIInput()

			fmt.Print("Enter Emp AddressLine2: ")
			addressdata.AddressLine2 = readCLIInput()

			fmt.Print("Enter Emp Street Name: ")
			addressdata.Street_Name = readCLIInput()

			fmt.Print("Enter Emp City Name: ")
			addressdata.City_Name = readCLIInput()

			fmt.Print("Enter Emp State Name: ")
			addressdata.State_Name = readCLIInput()

			fmt.Print("Enter PinCode: ")
			addressdata.Pincode = readCLIInput()

			fmt.Print("Enter Emp Basic Salary: ")
			salarydata.BSalary = readCLIInput()

			fmt.Print("Enter Emp CrossPay: ")
			salarydata.Cross_Pay = readCLIInput()

			fmt.Print("Enter Emp Department ID: ")
			depart.DeptID = readCLIInput()

			fmt.Print("Enter Department Name: ")
			depart.DeptName = readCLIInput()

			emp.AddEmployee(empdata)
			dept.AddEmployeeDepartment(depart)
			addres.AddEmployeeAddress(addressdata)
			slry.AddEmployeeSalary(salarydata)

		case 2:
			for {
				fmt.Println("\n 1. Get Employee Details")
				fmt.Println("\n 2. Get Employee Address Detail")
				fmt.Println("\n 3. Get Employee Department Detail")
				fmt.Println("\n 4. Get Employee Salary Detail")
				fmt.Println("\n 5. Press 5 to return main menu")
				addinput, _ := strconv.Atoi(readCLIInput())
				switch addinput {
				case 1:
					fmt.Println("\n Enter Employee ID")
					empid := readCLIInput()
					emp.GetEmployeeDetail(empid)
				case 2:
					fmt.Println("\n Enter Employee ID")
					empid := readCLIInput()
					addres.GetEmployeeAddress(empid)
				case 3:
					fmt.Println("\n Enter Employee ID")
					empid := readCLIInput()
					dept.GetEmployeeDepartment(empid)
				case 4:
					fmt.Println("\n Enter Employee ID")
					empid := readCLIInput()
					slry.GetEmployeeSalary(empid)
				default:
					fmt.Println("Press any to return main menu")
					break
				}
				break
			}
		case 3:
			for {
				fmt.Println("\n 1. Add All Employee Details")
				fmt.Println("\n 2. Add All Employee Address Detail")
				fmt.Println("\n 3. Add All Employee Department Detail")
				fmt.Println("\n 4. Add All Employee Salary Detail")
				fmt.Println("\n 5. Press 5 to return main menu")
				addinput, _ := strconv.Atoi(readCLIInput())
				switch addinput {
				case 1:
					data := emp.GetAllEmployees()
					for _, v := range data {
						fmt.Println("The datas are", v)
					}
				case 2:
					data := addres.GetAllEmployeesAddress()
					for _, v := range data {
						fmt.Println("The datas are", v)
					}
				case 3:
					data := dept.GetAllEmployeeDepartments()
					for _, v := range data {
						fmt.Println("The datas are", v)
					}
				case 4:
					data := slry.GetAllEmployeesSalary()
					for _, v := range data {
						fmt.Println("The datas are", v)
					}
				default:
					fmt.Println("Press any to return main menu")
					break
				}
				break
			}
		case 4:
			for {
				fmt.Println("\n 1. Update Employee Details")
				fmt.Println("\n 2. Update Employee Address Detail")
				fmt.Println("\n 3. Update Employee Department Detail")
				fmt.Println("\n 4. Update Employee Salary Detail")
				fmt.Println("\n 5. Press 5 to return main menu")
				addinput, _ := strconv.Atoi(readCLIInput())
				switch addinput {
				case 1:
					fmt.Print("Enter Emp ID: ")
					empdata.ID = readCLIInput()
					fmt.Print("Enter Emp Designation: ")
					empdata.Designation = readCLIInput()
					fmt.Print("Enter Emp EmailID: ")
					empdata.EmailID = readCLIInput()
					emp.UpdateEmployee(empdata)

				case 2:
					fmt.Print("Enter Emp ID: ")
					addressdata.EmpID = readCLIInput()
					fmt.Print("Enter Emp AddressLine1: ")
					addressdata.AddressLine1 = readCLIInput()
					fmt.Print("Enter Emp AddressLine2: ")
					addressdata.AddressLine2 = readCLIInput()
					fmt.Print("Enter Emp Street Name: ")
					addressdata.Street_Name = readCLIInput()
					fmt.Print("Enter Emp City Name: ")
					addressdata.City_Name = readCLIInput()
					fmt.Print("Enter Emp State Name: ")
					addressdata.State_Name = readCLIInput()
					fmt.Print("Enter Emp PinCode: ")
					addressdata.Pincode = readCLIInput()
					addres.UpdateEmployeeAddress(addressdata)
				case 3:
					fmt.Println("\n Enter Employee ID")
					depart.EmpID = readCLIInput()
					fmt.Print("Enter Emp Department ID: ")
					depart.DeptID = readCLIInput()
					fmt.Print("Enter Emp Department Name: ")
					depart.DeptName = readCLIInput()
					dept.UpdateEmployeeDepartment(depart)
				case 4:
					fmt.Println("\n Enter Employee ID")
					salarydata.EmpId = readCLIInput()
					fmt.Print("Enter Emp Basic Pay: ")
					salarydata.BSalary = readCLIInput()
					fmt.Print("Enter Emp CrossPay: ")
					salarydata.Cross_Pay = readCLIInput()
					slry.UpdateEmployeeSalary(salarydata)
				default:
					fmt.Println("Press any to return main menu")
					break
				}
				break
			}
		default:
			fmt.Println("\n Please enter valid input")
			break
		}
	}
}

func readCLIInput() (input string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")
	return
}

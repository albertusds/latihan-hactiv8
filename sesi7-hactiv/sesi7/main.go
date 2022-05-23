package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" //supaya libnya tetap ada, tambahin undrscore didepan
)

// config DB
const (
	DB_HOST = "localhost"
	DB_USER = "root"
	DB_PASS = "root"
	DB_NAME = "db_go_sql"
	DB_PORT = 5432
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

//Method employee for Print output
func (e *Employee) Print() {
	fmt.Println("ID", e.ID)
	fmt.Println("Fullname", e.FullName)
	fmt.Println("email", e.Email)
	fmt.Println("Division", e.Division)
	fmt.Println("")
}

func main() {
	// === CONNECT DB ===
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connected")

	// === CREATE EMPLOYEE ===
	// fmt.Println("CREATE EMPLOYEE")
	// //create employee
	// empl := Employee{
	// 	Email:    "test2@gmail.com",
	// 	FullName: "test2",
	// 	Age:      12,
	// 	Division: "Finance",
	// }

	// err = createEmployee(db, &empl)
	// if err != nil {
	// 	fmt.Println("Error Create Employee:", err.Error())
	// 	return
	// }

	// === GET ALL EMPLOYEE ===
	allEmployees, err := GetAllEmployees(db)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	for _, emp := range *allEmployees {
		emp.Print()
	}

	// === GET EMP BY ID ===
	fmt.Println("=== GET EMP BY ID ===")
	empById, err := GetEmpById(db, 2)
	if err != nil {
		fmt.Println("Error Get EMP By Id :", err.Error())
		return
	}
	empById.Print()

	//=== DELETE BY ID ===
	// fmt.Println(" === DELETE BY ID ===")
	// err = DeleteEmpById(db, 4)
	// if err != nil {
	// 	fmt.Println("Error Delete :", err.Error())
	// 	return
	// }

	// === UPDATE DATA ===
	fmt.Println("==== UPDATE DATA ====")
	updateEmp := Employee{
		ID:       2,
		FullName: "updateName1",
		Email:    "updateEmail1@gmail.com",
		Age:      22,
		Division: "UpdateDivision",
	}
	err = UpdateEmpById(db, &updateEmp)
	if err != nil {
		fmt.Println("Error when update :", err.Error())
		return
	}
}

/*
	Func connect to DB
*/
func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	fmt.Println("dsn:", dsn)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	// defer db.Close()
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

/*
	func get all employee
*/
func GetAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
	SELECT id, full_name, email, age, division
	FROM employees`

	stmt, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var employees []Employee

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee Employee //temp employee
		err := rows.Scan(
			&employee.ID, &employee.FullName, &employee.Email,
			&employee.Age, &employee.Division,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

/*
	func create employee
*/
func createEmployee(db *sql.DB, request *Employee) error {
	query := `
	INSERT INTO employees(full_name, email, age, division)
	VALUES ($1, $2, $3, $4)`

	//transaction based query
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

/*
	func get emp by id
*/
func GetEmpById(db *sql.DB, id int) (*Employee, error) {
	query := `
	SELECT id, full_name, email, age, division
	FROM employees
	WHERE id=$1`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var emp Employee

	err = row.Scan(
		&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

/*
function delete
*/
func DeleteEmpById(db *sql.DB, id int) error {
	query := `
	DELETE FROM employees
	WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Deleted data :", count)

	return err
}

/*
function update
*/
func UpdateEmpById(db *sql.DB, empl *Employee) error {
	query := `
	UPDATE employees SET full_name=$1, email=$2, age=$3, division=$4 
	WHERE id=$5
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(empl.FullName, empl.Email, empl.Age, empl.Division, empl.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Updated data :", count)

	return err
}

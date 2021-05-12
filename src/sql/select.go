package sql

import (
	"app/sql/domain"
	"database/sql"
	"log"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

type Gender int

type SelectOperator struct {
	Handler Handler
}

func (so *SelectOperator) SelectGenderWithOverAvgIncom(gender Gender) []domain.Employee {
	var employees []domain.Employee
	rows, err := so.Handler.DB.Query("SELECT name, income FROM employee WHER gender = ? and income > (SELECT AVG(income) FROM employeee)", gender)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var employee domain.Employee
		if err := rows.Scan(employee); err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}
	return employees
}

// func (so *SelectOperator) SelectTargetDepartmentInfo(dp_name string) []domain.DepartmentWithRelationsip { //relationshipで対象の名前を持つdepartmentの情報(従業員も含む)を取得する

// }

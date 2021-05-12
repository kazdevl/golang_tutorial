package selectquery

import (
	"app/selectquery/domain"
	"database/sql"
	"log"
)

type Gender int

const (
	IsFemail = iota
	IsMail   = 1
)

type SelectOperator struct {
	DB *sql.DB
}

func NewSelectOperator(db *sql.DB) *SelectOperator {
	return &SelectOperator{
		DB: db,
	}
}

func (so *SelectOperator) SelectGenderWithOverAvgIncome(gender Gender) []domain.Employee {
	var employees []domain.Employee
	rows, err := so.DB.Query("SELECT name, income FROM employee WHER gender = ? and income > (SELECT AVG(income) FROM employeee)", gender)

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

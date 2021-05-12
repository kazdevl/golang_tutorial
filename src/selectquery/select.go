package selectquery

import (
	"app/selectquery/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

func (so *SelectOperator) SelectGenderWithOverAvgIncome(gender Gender) ([]domain.Employee, error) {
	var employees []domain.Employee
	rows, err := so.DB.Query("SELECT * FROM employee WHERE gender = ? and income > (SELECT AVG(income) FROM employee)", gender)

	if err != nil {
		return employees, err
	}
	for rows.Next() {
		employee := &domain.Employee{}
		if err := rows.Scan(
			&employee.ID,
			&employee.DepartmentID,
			&employee.Income,
			&employee.Age,
			&employee.Gender,
			&employee.Name,
		); err != nil {
			return employees, err
		}
		employees = append(employees, *employee)
	}
	return employees, nil
}

// func (so *SelectOperator) SelectTargetDepartmentInfo(dp_name string) []domain.DepartmentWithRelationsip { //relationshipで対象の名前を持つdepartmentの情報(従業員も含む)を取得する

// }

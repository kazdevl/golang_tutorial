package selectquery

import (
	"database/sql"

	"github.com/kazdevl/golang_tutorial/sql/domain"

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

// 副問い合わせとは、SLQ文の中に、入子になってSQL文が書かれている事from・where・select句などで利用される
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

type incomeWithGenderGroup struct {
	gender int
	income int
}

// GROUP BYは、指定したカラム名において、同一の値を持つレコードをグルーピングして一つのレコードに集約する
func (so *SelectOperator) SelectMaxIncomeWithUnderAvgAgeAndGenderGroup() ([]incomeWithGenderGroup, error) {
	// HAVING(GROUP BYの実行後)とWHERE(GROUP BYの実行前)で条件つけできる
	rows, err := so.DB.Query("SELECT gender, MAX(income) FROM employee WHERE age < (SELECT AVG(age) FROM employee) GROUP BY gender")

	if err != nil {
		return []incomeWithGenderGroup{}, err
	}

	var data []incomeWithGenderGroup
	for rows.Next() {
		var obj incomeWithGenderGroup
		if err := rows.Scan(&obj.gender, &obj.income); err != nil {
			return []incomeWithGenderGroup{}, err
		}
		data = append(data, obj)
	}
	return data, nil
}

type departmentInfo struct {
	name      string
	maxIncome int
	avgIncome float64
}

// JOINは, table同士を結合するもの
func (so *SelectOperator) SelectDepartmentInfo() ([]departmentInfo, error) { //relationshipで対象の名前を持つdepartmentの情報(従業員関連情報も含む)を取得する
	rows, err := so.DB.Query(
		"SELECT d.name, MAX(e.income) as max_icome, AVG(e.income) as avg_income " +
			"FROM department AS d " +
			"INNER JOIN employee AS e ON d.id = e.department_id " +
			"GROUP BY e.department_id") // 内部結合後にグループ化している
	if err != nil {
		return []departmentInfo{}, err
	}

	var data []departmentInfo
	for rows.Next() {
		var row departmentInfo
		if err := rows.Scan(&row.name, &row.maxIncome, &row.avgIncome); err != nil {
			return []departmentInfo{}, err
		}
		data = append(data, row)
	}
	return data, nil
}

type employeeInfo struct {
	domain.Employee
	DepartmentName string
}

func (so *SelectOperator) SelectEmployeeInfo() ([]employeeInfo, error) {
	rows, err := so.DB.Query(
		"SELECT e.id, e.name, income, age, gender, department_id, d.name " +
			"FROM employee AS e INNER JOIN department AS d on e.department_id = d.id")
	if err != nil {
		return []employeeInfo{}, err
	}

	var data []employeeInfo
	for rows.Next() {
		var row employeeInfo
		if err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Income,
			&row.Age,
			&row.Gender,
			&row.DepartmentID,
			&row.DepartmentName,
		); err != nil {
			return []employeeInfo{}, err
		}
		data = append(data, row)
	}
	return data, nil
}

// 把握しておくと便利なもの
// 1. JOIN(INNER・OUTER(RIGHT・LEFT))
// 2. GROUP BY...同じデータをグループとして集約する
// 3. GROUP_CONCAT...グループとして集約した際に、NULL以外の値を含む文字列を連結してくれる
// 4. DISTINCT...重複レコードを一つにまとめる
// 5. 副問い合わせ...SQL文の中でSQL文を利用すること
// 6. UINION...2つ以上のSELECTの結果を、結合して表示する
// 7. その他集約関数など

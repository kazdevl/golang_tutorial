package domain

type Department struct {
	ID   int
	Name string
}

type DepartmentWithRelationsip struct {
	ID       int
	Name     string
	Employee Employee
}

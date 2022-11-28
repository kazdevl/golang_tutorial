package target

import (
	"fmt"

	"github.com/kazdevl/golang_tutorial/sql/domain"
	"github.com/kazdevl/golang_tutorial/staticlint/typechecks/target/repository"
)

type BlogService struct {
	userModelRepository repository.IFUserModelRepository
	postModelRepository repository.IFPostModelRepository
	sample              domain.Department
}

func NewBlogService(
	umr *repository.UserModelRepository,
	pmr *repository.PostModelRepository,
	sample domain.Department,
) *BlogService {
	return &BlogService{
		userModelRepository: umr,
		postModelRepository: pmr,
		sample:              sample,
	}
}

func (s *BlogService) Sample() {
	fmt.Print("Hello World")
}

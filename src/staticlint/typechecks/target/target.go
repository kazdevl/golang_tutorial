package target

import (
	"fmt"

	"github.com/kazdevl/golang_tutorial/staticlint/typechecks/target/repository"
)

type BlogService struct {
	userModelRepository repository.IFUserModelRepository
	postModelRepository repository.IFPostModelRepository
}

func NewBlogService(
	umr *repository.UserModelRepository,
	pmr *repository.PostModelRepository,
) *BlogService {
	return &BlogService{
		userModelRepository: umr,
		postModelRepository: pmr,
	}
}

func (s *BlogService) Sample() {
	fmt.Print("Hello World")
}

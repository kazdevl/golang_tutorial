package target

import "app/staticlint/target/repository"

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

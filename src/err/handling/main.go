package main

import (
	"app/err/handling/perror"
	"log"

	"github.com/pkg/errors"
)

func main() {
	r := &Repository{}
	s := Service{r: r}
	err := s.Exist(Model{ID: 1})
	// if err != nil {
	// 	perr := perror.ProductError{}
	// 	if ok := errors.As(err, &perr); ok {
	// 		log.Printf("product error: %s", perr.Error())
	// 		// なんかの処理
	// 		os.Exit(0)
	// 	}
	// 	log.Fatalf("error: %+v", err)
	// }
	if cerr := errors.Cause(err); cerr != nil {
		perr := cerr.(perror.ProductError)
		log.Printf("err cause: %+v", perr)
	}
}

type Model struct {
	ID int
}

type Repository struct{}

func (r *Repository) Find(id int) (Model, error) {
	// DBアクセスで生まれたエラー
	err := errors.New("sample")
	return Model{}, perror.Wrapf(err, perror.ErrorCase_FailedDB, "failed to get model: ID=%d", id)
}

type Service struct {
	r *Repository
}

func (s *Service) Exist(m Model) error {
	if _, err := s.r.Find(m.ID); err != nil {
		return perror.Stack(err)
	}
	return nil
}

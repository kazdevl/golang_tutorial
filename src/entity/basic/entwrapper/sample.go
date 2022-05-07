package entwrapper

import (
	"app/entity/basic/ent"
	"app/entity/basic/model/db"
	"context"
	"log"
)

func NewClient() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed: %v", err)
	}
	return client
}

type Sample struct {
	Client *ent.Client
}

func NewSample(c *ent.Client) *Sample {
	return &Sample{
		Client: c,
	}
}

func (s *Sample) Get(ctx context.Context, id int) (*db.Sample, error) {
	v, err := s.Client.Sample.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	sampleModel := &db.Sample{
		ID:   v.ID,
		Age:  v.Age,
		Name: v.Name,
	}
	return sampleModel, nil
}

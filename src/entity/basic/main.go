package main

import (
	"app/entity/basic/ent"
	"context"
	"fmt"
	"log"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func CreateSample(ctx context.Context, client *ent.Client) (*ent.Sample, error) {
	s, err := client.Sample.
		Create().SetAge(30).SetName("sample").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("sample was created: ", s)
	return s, nil
}

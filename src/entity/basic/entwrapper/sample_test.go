package entwrapper_test

import (
	"context"
	"testing"

	"github.com/kazdevl/golang_tutorial/entity/basic/ent"
	"github.com/kazdevl/golang_tutorial/entity/basic/ent/enttest"
	"github.com/kazdevl/golang_tutorial/entity/basic/entwrapper"
	"github.com/kazdevl/golang_tutorial/entity/basic/model/db"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
)

var (
	inputSamples = []*db.Sample{
		{ID: 1, Age: 10, Name: "sample"}, //IDに関してはauto incrementのため
		{ID: 2, Age: 10, Name: "sample"},
	}
)

func InputData(ctx context.Context, client *ent.Client) error {
	for _, input := range inputSamples {
		_, err := client.Sample.Create().SetAge(input.Age).SetName(input.Name).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// 欲しいデータ
	want := inputSamples[0]

	// テスト用の値入力
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1") // 内部でmigarteしてるのでtestのたびにデータは消える
	defer client.Close()
	err := InputData(ctx, client)
	assert.NoError(t, err)

	// テストしたいやつ
	sampleWrapper := entwrapper.NewSample(client)
	result, err := sampleWrapper.Get(ctx, want.ID)

	// 動作的に問題ないかの確認
	assert.NoError(t, err)
	assert.Equal(t, want.Age, result.Age)
	assert.Equal(t, want.Name, result.Name)
}

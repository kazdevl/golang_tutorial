package formock_test

import (
	"app/test/formock"
	mock_formock "app/test/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAPI(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockValidater := mock_formock.NewMockFormValidater(ctrl)
	// mockに期待する呼び出しと、返り値を定義
	// EXPECT()は呼び出し元がmockしたいオブジェクトを返す(*MockFormValidaterMockRecorder)
	// Validate()では、そのメソッドが指定した引数で呼び出されているか
	// Return()では、mockの関数が返す値を定義
	mockValidater.EXPECT().Validate("hello world").Return("hello world", nil)

	api := formock.API{Validater: mockValidater}
	if err := api.RegisterData("hello world"); err != nil {
		t.Fatal("Register Error")
	}
}

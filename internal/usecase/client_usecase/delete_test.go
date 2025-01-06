package client_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project2/internal/usecase/client_usecase/mocks"
	"testing"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	//зависимости, которые нужны для теста
	type fields struct {
		clientRepo *mocks.ClientRepo
	}

	//данные для теста
	type args struct {
		clientID int
	}

	//тесты
	tests := []struct {
		name    string
		args    args
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				clientID: 2,
			},
			wantErr: false,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().Delete(args.clientID).Return(nil)
			},
		},
		{
			name: "error deleted client",
			args: args{
				clientID: 2,
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().Delete(args.clientID).Return(errors.New("error deleted client"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			//создали зависимости
			f := fields{
				clientRepo: mocks.NewClientRepo(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.clientRepo)

			//выполнили
			err := uc.Delete(tt.args.clientID)

			//проверяем результат
			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
		})
	}

}

package client_usecase

import (
	"errors"
	"testing"

	"medicalCenter/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	//данные для теста
	type args struct {
		clientID int
		timer    *mocks.Timer
	}

	//тесты
	tests := []struct {
		name    string
		args    args
		wantErr error
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				clientID: 2,
			},
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().Delete(args.clientID).Return(nil)
			},
		},
		{
			name: "error deleted client",
			args: args{
				clientID: 2,
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().Delete(args.clientID).Return(errTest)
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
				timer:      mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.clientRepo, f.timer)

			//выполнили
			err := uc.Delete(tt.args.clientID)

			//проверяем результат
			if tt.wantErr != nil {
				a.ErrorIs(err, tt.wantErr)
				return
			}

			a.NoError(err)
		})
	}
}

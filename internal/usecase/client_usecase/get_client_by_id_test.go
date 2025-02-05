package client_usecase

import (
	"errors"
	"testing"
	"time"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetClientByID(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	//данные для теста
	type args struct {
		ID    int
		timer *mocks.Timer
	}

	//тесты
	tests := []struct {
		name    string
		args    args
		want    *domain.Client
		wantErr error
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				ID: 2,
			},
			want: &domain.Client{
				Name:        "Artem",
				BDate:       now,
				PhoneNumber: "+7888888",
				UpdatedAt:   now,
			},
			before: func(f fields, args args) {
				client := domain.NewClient("Artem", now, "+7888888", now)
				client.BDate = now
				client.UpdatedAt = now

				f.clientRepo.EXPECT().GetClientByID(args.ID).Return(client, nil)
			},
		},
		{
			name: "error get client",
			args: args{
				ID: 2,
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().GetClientByID(args.ID).Return(nil, errTest)
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
			e, err := uc.GetClientByID(tt.args.ID)

			//проверяем результат
			if tt.wantErr != nil {
				a.ErrorIs(err, tt.wantErr)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, e)
		})
	}
}

package client_usecase

import (
	"errors"
	"testing"
	"time"

	"project2/internal/domain"
	"project2/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase(t *testing.T) {
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
		req CreateClientReq
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
				req: CreateClientReq{
					Name:        "Poly",
					BDate:       now,
					PhoneNumber: "+7999999",
				},
			},
			want: &domain.Client{
				Name:        "Poly",
				BDate:       now,
				PhoneNumber: "+7999999",
				UpdatedAt:   now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)

				client := domain.NewClient(args.req.Name, args.req.BDate, args.req.PhoneNumber, now)
				f.clientRepo.EXPECT().Create(client).Return(client, nil)
			},
		},
		{
			name: "error on create client",
			args: args{
				req: CreateClientReq{
					Name:        "Poly",
					BDate:       now,
					PhoneNumber: "+7999999",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)

				client := domain.NewClient(args.req.Name, args.req.BDate, args.req.PhoneNumber, now)
				f.clientRepo.EXPECT().Create(client).Return(nil, errTest)
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
			e, err := uc.Create(tt.args.req)

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

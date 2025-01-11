package client_usecase

import (
	"errors"
	"testing"

	"project2/internal/domain"
	"project2/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		clientRepo *mocks.ClientRepo
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
					BDate:       "20.10.2010",
					PhoneNumber: "+7999999",
				},
			},
			want: &domain.Client{
				Name:        "Poly",
				BDate:       "20.10.2010",
				PhoneNumber: "+7999999",
			},
			before: func(f fields, args args) {
				client := domain.NewClient(args.req.Name, args.req.BDate, args.req.PhoneNumber)
				f.clientRepo.EXPECT().Create(client).Return(client, nil)
			},
		},
		{
			name: "error on create client",
			args: args{
				req: CreateClientReq{
					Name:        "Poly",
					BDate:       "20.10.2010",
					PhoneNumber: "+7999999",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				client := domain.NewClient(args.req.Name, args.req.BDate, args.req.PhoneNumber)
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
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.clientRepo)

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

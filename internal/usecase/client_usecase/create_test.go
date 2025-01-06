package client_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain"
	"project2/internal/usecase/client_usecase/mocks"
	"testing"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

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
		wantErr bool
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
			wantErr: false,
			before: func(f fields, args args) {
				client := domain.NewClient(args.req.Name, args.req.BDate, args.req.PhoneNumber)
				f.clientRepo.EXPECT().Create(client).Return(nil)
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
			wantErr: true,
			before: func(f fields, args args) {
				client := domain.NewClient(args.req.Name, args.req.BDate, args.req.PhoneNumber)
				f.clientRepo.EXPECT().Create(client).Return(errors.New("error on create client"))
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
			err := uc.Create(tt.args.req)

			//проверяем результат
			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
		})
	}

}

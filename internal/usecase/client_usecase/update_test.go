package client_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain"
	"project2/internal/usecase/client_usecase/mocks"
	"testing"
)

func TestUpdateUseCase(t *testing.T) {
	t.Parallel()

	//зависимости, которые нужны для теста
	type fields struct {
		clientRepo *mocks.ClientRepo
	}

	//данные для теста
	type args struct {
		req UpdateClientReq
	}

	client := domain.NewClient("Artem", "30.12.1999", "+7888888")

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
				req: UpdateClientReq{
					ID:          4,
					Name:        "Poly",
					BDate:       "20.10.2010",
					PhoneNumber: "+7999999",
				},
			},
			wantErr: false,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.ID).Return(client, nil)
				client.Name = args.req.Name
				client.BDate = args.req.BDate
				client.PhoneNumber = args.req.PhoneNumber
				f.clientRepo.EXPECT().Update(client).Return(nil)
			},
		},
		{
			name: "error on FindByID",
			args: args{
				req: UpdateClientReq{
					ID:          4,
					Name:        "Poly",
					BDate:       "20.10.2010",
					PhoneNumber: "+7999999",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.ID).Return(client, errors.New("error on FindByID"))
			},
		},
		{
			name: "error on update client",
			args: args{
				req: UpdateClientReq{
					ID:          4,
					Name:        "Poly",
					BDate:       "20.10.2010",
					PhoneNumber: "+7999999",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.ID).Return(client, nil)
				client.Name = args.req.Name
				client.BDate = args.req.BDate
				client.PhoneNumber = args.req.PhoneNumber
				f.clientRepo.EXPECT().Update(client).Return(errors.New("error on update client"))
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
			err := uc.Update(tt.args.req)

			//проверяем результат
			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
		})
	}

}

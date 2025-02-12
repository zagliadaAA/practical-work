package client_usecase

import (
	"errors"
	"testing"
	"time"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	//данные для теста
	type args struct {
		req UpdateClientReq
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
				req: UpdateClientReq{
					ID:          4,
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

				client := domain.NewClient("Artem", now, "+7888888", now)

				f.clientRepo.EXPECT().GetClientByID(args.req.ID).Return(client, nil)
				client.Name = args.req.Name
				client.BDate = args.req.BDate
				client.PhoneNumber = args.req.PhoneNumber
				f.clientRepo.EXPECT().Update(client).Return(client, nil)
			},
		},
		{
			name: "error on GetClientByID",
			args: args{
				req: UpdateClientReq{
					ID:          4,
					Name:        "Poly",
					BDate:       now,
					PhoneNumber: "+7999999",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().GetClientByID(args.req.ID).Return(nil, errTest)
			},
		},
		{
			name: "error on update client",
			args: args{
				req: UpdateClientReq{
					ID:          4,
					Name:        "Poly",
					BDate:       now,
					PhoneNumber: "+7999999",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)

				client := domain.NewClient("Artem", now, "+7888888", now)

				f.clientRepo.EXPECT().GetClientByID(args.req.ID).Return(client, nil)
				client.Name = args.req.Name
				client.BDate = args.req.BDate
				client.PhoneNumber = args.req.PhoneNumber
				f.clientRepo.EXPECT().Update(client).Return(nil, errTest)
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
			e, err := uc.Update(tt.args.req)

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

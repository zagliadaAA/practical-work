package medical_report_usecase

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"project2/internal/domain"
	"project2/internal/usecase/medical_report_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		medRepo    *mocks.MedRepo
		clientRepo *mocks.ClientRepo
	}

	//данные для теста
	type args struct {
		req CreateMedicalReportReq
	}

	now := time.Now().UTC()
	dateStr := now.Format("02.01.2006 15:04")
	createdAt, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		fmt.Println("не получилось преобразовать дату")
	}

	//тесты
	tests := []struct {
		name    string
		args    args
		want    *domain.MedicalReport
		wantErr error
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateMedicalReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			want: &domain.MedicalReport{
				DoctorName: "Ложкин",
				Diagnosis:  "A77",
				CreatedAt:  createdAt,
				IDClient:   4,
			},
			before: func(f fields, args args) {
				client := domain.NewClient("Artem", time.Date(1999, 12, 10, 0, 0, 0, 0, time.UTC), "+70000000000")
				client.ID = args.req.IDClient

				f.clientRepo.EXPECT().FindByID(args.req.IDClient).Return(client, nil)
				report := domain.NewMedicalReport(args.req.DoctorName, args.req.Diagnosis, client.ID)
				f.medRepo.EXPECT().Create(report).Return(report, nil)
			},
		},
		{
			name: "error on get client",
			args: args{
				req: CreateMedicalReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.IDClient).Return(nil, errTest)
			},
		},
		{
			name: "error on create client",
			args: args{
				req: CreateMedicalReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				client := domain.NewClient("Artem", time.Date(1999, 12, 10, 0, 0, 0, 0, time.UTC), "+70000000000")
				client.ID = args.req.IDClient

				f.clientRepo.EXPECT().FindByID(args.req.IDClient).Return(client, nil)
				report := domain.NewMedicalReport(args.req.DoctorName, args.req.Diagnosis, client.ID)
				f.medRepo.EXPECT().Create(report).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			//создали зависимости
			f := fields{
				medRepo:    mocks.NewMedRepo(t),
				clientRepo: mocks.NewClientRepo(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.medRepo, f.clientRepo)

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

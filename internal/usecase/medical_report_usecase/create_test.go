package medical_report_usecase

import (
	"errors"
	"testing"

	"project2/internal/domain"
	"project2/internal/usecase/medical_report_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	//зависимости, которые нужны для теста
	type fields struct {
		medRepo    *mocks.MedRepo
		clientRepo *mocks.ClientRepo
	}

	//данные для теста
	type args struct {
		req CreateMedicalReportReq
	}

	client := domain.NewClient("Artem", "30.12.1999", "+70000000000")

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
				req: CreateMedicalReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.IDClient).Return(client, nil)
				report := domain.NewMedicalReport(args.req.DoctorName, args.req.Diagnosis, client.ID)
				f.medRepo.EXPECT().Create(report).Return(nil)
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
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.IDClient).Return(nil, errors.New("error get client"))
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
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindByID(args.req.IDClient).Return(client, nil)
				report := domain.NewMedicalReport(args.req.DoctorName, args.req.Diagnosis, client.ID)
				f.medRepo.EXPECT().Create(report).Return(errors.New("error create client"))
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

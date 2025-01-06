package medical_report_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain"
	"project2/internal/usecase/medical_report_usecase/mocks"
	"testing"
)

func TestUpdateUseCase(t *testing.T) {
	t.Parallel()

	//зависимости, которые нужны для теста
	type fields struct {
		medRepo    *mocks.MedRepo
		clientRepo *mocks.ClientRepo
	}

	//данные для теста
	type args struct {
		req UpdateReportReq
	}

	report := domain.NewMedicalReport("Вова Лекарь", "Z.17777", 4)

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
				req: UpdateReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: false,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.req.IDClient).Return(report, nil)
				report.DoctorName = args.req.DoctorName
				report.Diagnosis = args.req.Diagnosis
				f.medRepo.EXPECT().Update(report).Return(nil)
			},
		},
		{
			name: "error on get report",
			args: args{
				req: UpdateReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.req.IDClient).Return(nil, errors.New("error on get report by id client"))
			},
		},
		{
			name: "error on update report",
			args: args{
				req: UpdateReportReq{
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.req.IDClient).Return(report, nil)
				report.DoctorName = args.req.DoctorName
				report.Diagnosis = args.req.Diagnosis
				f.medRepo.EXPECT().Update(report).Return(errors.New("error on update report"))
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
package medical_report_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project2/internal/domain"
	"project2/internal/usecase/medical_report_usecase/mocks"
	"testing"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	//зависимости, которые нужны для теста
	type fields struct {
		medRepo    *mocks.MedRepo
		clientRepo *mocks.ClientRepo
	}

	//данные для теста
	type args struct {
		clientID int
		reportID int
	}

	report := domain.MedicalReport{
		ID:         2,
		DoctorName: "Вова Лекарь",
		Diagnosis:  "Z.17777",
		IDClient:   4,
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
				clientID: 4,
				reportID: 2,
			},
			wantErr: false,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.clientID).Return(&report, nil)
				f.medRepo.EXPECT().Delete(args.reportID).Return(nil)
			},
		},
		{
			name: "error on get report",
			args: args{
				clientID: 4,
				reportID: 2,
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.clientID).Return(nil, errors.New("error on get report by id client"))
			},
		},
		{
			name: "error deleting report",
			args: args{
				clientID: 4,
				reportID: 2,
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.clientID).Return(&report, nil)
				f.medRepo.EXPECT().Delete(args.reportID).Return(errors.New("error deleting report"))
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
			err := uc.Delete(tt.args.clientID)

			//проверяем результат
			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
		})
	}

}

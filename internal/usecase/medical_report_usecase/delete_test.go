package medical_report_usecase

import (
	"errors"
	"testing"
	"time"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/medical_report_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		medRepo    *mocks.MedRepo
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	//данные для теста
	type args struct {
		clientID int
		reportID int
	}

	//тесты
	tests := []struct {
		name    string
		args    args
		wantErr error
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				clientID: 4,
				reportID: 2,
			},
			before: func(f fields, args args) {
				report := domain.NewMedicalReport("Вова Лекарь", "Z.17777", 4, now, now)
				report.ID = args.reportID

				f.medRepo.EXPECT().GetReportByIDClient(args.clientID).Return(report, nil)
				f.medRepo.EXPECT().Delete(report.ID).Return(nil)
			},
		},
		{
			name: "error on get report",
			args: args{
				clientID: 4,
				reportID: 2,
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByIDClient(args.clientID).Return(nil, errTest)
			},
		},
		{
			name: "error deleting report",
			args: args{
				clientID: 4,
				reportID: 2,
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				report := domain.NewMedicalReport("Вова Лекарь", "Z.17777", 4, now, now)
				report.ID = args.reportID

				f.medRepo.EXPECT().GetReportByIDClient(args.clientID).Return(report, nil)
				f.medRepo.EXPECT().Delete(report.ID).Return(errTest)
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
				timer:      mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.medRepo, f.clientRepo, f.timer)

			//выполнили
			err := uc.Delete(tt.args.clientID)

			//проверяем результат
			if tt.wantErr != nil {
				a.ErrorIs(err, tt.wantErr)
				return
			}

			a.NoError(err)
		})
	}
}

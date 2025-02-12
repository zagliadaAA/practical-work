package medical_report_usecase

import (
	"errors"
	"testing"
	"time"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/medical_report_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetReportByID(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("error test")

	//зависимости, которые нужны для теста
	type fields struct {
		medRepo    *mocks.MedRepo
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	//данные для теста
	type args struct {
		ID    int
		timer *mocks.Timer
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
				ID: 2,
			},
			want: &domain.MedicalReport{
				ID:         2,
				IDClient:   4,
				DoctorName: "Ложкин В",
				Diagnosis:  "A77.7",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			before: func(f fields, args args) {
				report := domain.NewMedicalReport("Ложкин В", "A77.7", 4, now, now)
				report.ID = 2

				f.medRepo.EXPECT().GetReportByID(args.ID).Return(report, nil)
			},
		},
		{
			name: "error get report",
			args: args{
				ID: 2,
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByID(args.ID).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			//создали зависимости
			f := fields{
				medRepo: mocks.NewMedRepo(t),
				timer:   mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.medRepo, f.clientRepo, f.timer)

			//выполнили
			e, err := uc.GetReportByID(tt.args.ID)

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

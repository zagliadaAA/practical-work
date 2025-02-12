package medical_report_usecase

import (
	"errors"
	"testing"
	"time"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/medical_report_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUseCase(t *testing.T) {
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
		req UpdateReportReq
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
				req: UpdateReportReq{
					ID:         2,
					IDClient:   4,
					DoctorName: "Ложкин В",
					Diagnosis:  "A77.7",
				},
			},
			want: &domain.MedicalReport{
				IDClient:   4,
				DoctorName: "Ложкин В",
				Diagnosis:  "A77.7",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)

				report := domain.NewMedicalReport("Вова Лекарь", "Z.17777", 4, now, now)

				f.medRepo.EXPECT().GetReportByID(args.req.ID).Return(report, nil)
				report.DoctorName = args.req.DoctorName
				report.Diagnosis = args.req.Diagnosis
				f.medRepo.EXPECT().Update(report).Return(report, nil)
			},
		},
		{
			name: "error on get report",
			args: args{
				req: UpdateReportReq{
					ID:         2,
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.medRepo.EXPECT().GetReportByID(args.req.ID).Return(nil, errTest)
			},
		},
		{
			name: "error on update report",
			args: args{
				req: UpdateReportReq{
					ID:         2,
					IDClient:   4,
					DoctorName: "Ложкин",
					Diagnosis:  "A77",
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)

				report := domain.NewMedicalReport("Вова Лекарь", "Z.17777", 4, now, now)

				f.medRepo.EXPECT().GetReportByID(args.req.ID).Return(report, nil)
				report.DoctorName = args.req.DoctorName
				report.Diagnosis = args.req.Diagnosis
				f.medRepo.EXPECT().Update(report).Return(nil, errTest)
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

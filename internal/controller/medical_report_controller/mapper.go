package medical_report_controller

import (
	"medicalCenter/internal/domain"
)

func mapMedicalReportToResponseForCreate(medicalReport *domain.MedicalReport) *createMedicalReportResp {
	return &createMedicalReportResp{
		ID:         medicalReport.ID,
		DoctorName: medicalReport.DoctorName,
		Diagnosis:  medicalReport.Diagnosis,
		CreatedAt:  medicalReport.CreatedAt,
		UpdatedAt:  medicalReport.UpdatedAt,
		IdClient:   medicalReport.IDClient,
	}
}

func mapMedicalReportToResponseForUpdate(medicalReport *domain.MedicalReport) *updateMedicalReportResp {
	return &updateMedicalReportResp{
		ID:         medicalReport.ID,
		DoctorName: medicalReport.DoctorName,
		Diagnosis:  medicalReport.Diagnosis,
		CreatedAt:  medicalReport.CreatedAt,
		UpdatedAt:  medicalReport.UpdatedAt,
		IdClient:   medicalReport.IDClient,
	}
}

func mapReportToResponseForGetReportByID(medicalReport *domain.MedicalReport) *getReportByIDResp {
	return &getReportByIDResp{
		ID:         medicalReport.ID,
		DoctorName: medicalReport.DoctorName,
		Diagnosis:  medicalReport.Diagnosis,
		CreatedAt:  medicalReport.CreatedAt,
		UpdatedAt:  medicalReport.UpdatedAt,
		IdClient:   medicalReport.IDClient,
	}
}

package serializer

import (
	"iip/model"
)

type ProcedureStatus struct {
	ID              uint        `json:"id"`
	ProcedureStatus string      `json:"procedure_status"`
	Procedures      []Procedure `json:"procedures"`
	Description     string      `json:"description"`
}

func BuildProcedureStatus(procedureStatus model.ProcedureStatus) ProcedureStatus {
	return ProcedureStatus{
		ID:              procedureStatus.ID,
		ProcedureStatus: procedureStatus.ProcedureStatus,
		Procedures:      BuildProcedures(procedureStatus.Procedures),
		Description:     procedureStatus.Description,
	}
}

func BuildProcedureStatuss(items []model.ProcedureStatus) (ProcedureStatuss []ProcedureStatus) {
	for _, item := range items {
		ProcedureStatus := BuildProcedureStatus(item)
		ProcedureStatuss = append(ProcedureStatuss, ProcedureStatus)
	}
	return ProcedureStatuss
}

package serializer

import (
	"iip/model"
)

type ProcedureType struct {
	ID            uint   `json:"id"`
	ProcedureType string `json:"procedure_type"`
	ProcedureName string `json:"procedure_name"`
	Decription    string `json:"description"`
}

func BuildProcedureType(procedureType model.ProcedureType) ProcedureType {
	return ProcedureType{
		ID:            procedureType.ID,
		ProcedureType: procedureType.ProcedureType,
		ProcedureName: procedureType.ProcedureName,
		Decription:    procedureType.Decription,
	}
}

func BuildProcedureTypes(items []model.ProcedureType) (ProcedureTypes []ProcedureType) {
	for _, item := range items {
		ProcedureType := BuildProcedureType(item)
		ProcedureTypes = append(ProcedureTypes, ProcedureType)
	}
	return ProcedureTypes
}

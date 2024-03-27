package serializer

import (
	"iip/model"
	"time"
)

type Procedure struct {
	ID                uint            `json:"id"`
	ProcedureType     string            `json:"procedure_type"`
	ProcedureName     string            `json:"procedure_name"`   ///string 和int
	TaskID            uint            `json:"task_id"`
	MachineID         uint            `json:"machine_id"`
	WorkpiecesNum     uint            `json:"workpieces_num"`
	ProcessedNum      uint            `json:"processed_num"`
	Sequence          uint            `json:"sequence"`
	MachineSequence   uint            `json:"machine_sequence"`
	StartDate         time.Time       `json:"start_date"`
	IsFinished        uint            `json:"is_finished"`
	ProcedureStatusID uint            `json:"procedure_status_id"`
	ProcedureItems    []ProcedureItem `json:"procedure_items"`
}

func BuildProcedure(procedure model.Procedure) Procedure {
	return Procedure{
		ID:                procedure.ID,
		ProcedureType:     procedure.ProcedureType,
		ProcedureName:     procedure.ProcedureName, 
		TaskID:            procedure.TaskID,
		MachineID:         procedure.MachineID,
		WorkpiecesNum:     procedure.WorkpiecesNum,
		ProcessedNum:      procedure.ProcessedNum,
		Sequence:          procedure.Sequence,
		MachineSequence:   procedure.MachineSequence,
		StartDate:         procedure.StartDate,
		IsFinished:        procedure.IsFinished,
		ProcedureStatusID: procedure.ProcedureStatusID,
		ProcedureItems:    BuildProcedureItems(procedure.ProcedureItems),
	}
}

func BuildProcedures(items []model.Procedure) (Procedures []Procedure) {
	for _, item := range items {
		Procedure := BuildProcedure(item)
		Procedures = append(Procedures, Procedure)
	}
	return Procedures
}

package serializer

import (
	"iip/model"
)

type ProcedureItem struct {
	ID          uint   `json:"id"`
	ProcedureID uint   `json:"procedure_id"`
	MachineID   uint   `json:"machine_id"`
	WorkerID    uint   `json:"worker_id"`
	ItemType    string `json:"item_type"`
	ItemName    string `json:"item_name"`
	Quality     string `json:"quality"` //品质分析怎么做
}

func BuildProcedureItem(procedureItem model.ProcedureItem) ProcedureItem {
	return ProcedureItem{
		ID:          procedureItem.ID,
		ProcedureID: procedureItem.ProcedureID,
		MachineID:   procedureItem.MachineID,
		WorkerID:    procedureItem.WorkerID,
		ItemType:    procedureItem.ItemType,
		ItemName:    procedureItem.ItemName,
		Quality:     procedureItem.Quality,
	}
}

func BuildProcedureItems(items []model.ProcedureItem) (ProcedureItems []ProcedureItem) {
	for _, item := range items {
		ProcedureItem := BuildProcedureItem(item)
		ProcedureItems = append(ProcedureItems, ProcedureItem)
	}
	return ProcedureItems
}

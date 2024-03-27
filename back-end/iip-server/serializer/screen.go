package serializer

import (
	"iip/model"
	"time"
)

// 车间生产状态
type ProductionBoard struct {
	MachineID     uint   `json:"machine_id"`
	MachineName   string `json:"machine_name"`
	WorkerName    string `json:"worker_name"`
	ProcedureID   uint   `json:"procrdure_id"`
	ProcedureName string `json:"procrdure_name"`
	ProcedureType string `json:"procrdure_type"`

	MachineStatusID uint    `json:"machine_status"`
	TaskID          uint    `json:"task_id"`
	Time            string  `json:"time"`
	Percent         float64 `json:"percent"`
}

type StatAnomaly struct {
	Anomalys1 []model.MachineAnomaly
	Anomalys2 []model.MachineAnomaly
	Duration1 []time.Duration
	Duration2 []time.Duration
	Count     []int64
}

// 生产实况
type StatusNum struct {
	NumOfRun     int64 `json:"num_of_run"`
	NumOfWaiting int64 `json:"num_of_waiting"`
	NumOfClose   int64 `json:"num_of_close"`
	NumOfError   int64 `json:"num_of_error"`
}
type OEE struct {
	OEEUp80   uint `json:"oee_uo_80"`
	OEEDown50 uint `json:"oee_down_50"`
	OEE50To80 uint `json:"oee_50_to_80"`
}
type MachineScreen struct {
	MachineID     uint   `json:"machine_id"`
	MachineName   string `json:"machine_name"`
	MachineStatus string `json:"machine_status"`
	StatusType    string `json:"status_type"`
	WorkpiecesNum uint   `json:"workpieces_num"`
	ProcessedNum  uint   `json:"processed_num"`
	Progress      uint   `json:"progress"`
}
type AnomalyScreen struct {
	MachineID   uint   `json:"machine_id"`
	MachineName string `json:"machine_name"`
	Reason      string `json:"reason"`
	StatusType  string `json:"status_type"`
	Date        string `json:"date"`
	// Time        string `json:"time"`
}
type ProductionStatus struct {
	StatusNum StatusNum       `json:"status_num"`
	OEE       OEE             `json:"oee"`
	Machines  []MachineScreen `json:"machines"`
	Anomalys  []AnomalyScreen `json:"anomalys"`
}

// 设备综合效率OEE分析
type OEERankItem struct {
	MachineID uint64
	Run       float64
	Wating    float64
	Closed    float64
	Error     float64
}
type OEELossItem struct {
	Date  string
	Value int64
}
type OEERankItems []OEERankItem

func (items OEERankItems) Len() int {
	return len(items)
}

// Less 方法比较两个元素的大小，这里我们根据 Run 的值来比较
func (items OEERankItems) Less(i, j int) bool {
	return items[i].Run > items[j].Run // 降序排序
}

// Swap 方法交换两个元素的位置
func (items OEERankItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

type OEEReasonItem struct {
	Reason string
	Value  int64
}
type OEETrendItem struct {
	Date  string
	Value float64
}
type OEEAnalysis struct {
	OEERank           OEERankItems
	OEERankPercent    OEERankItems
	OEEStatusCount    []StatusCount
	OEEStatusDayCount []StatusDayCount
	OEELoss           []OEELossItem
	OEEReason         []OEEReasonItem
	OEETrend          []OEETrendItem
}

// 异常分析
type AnomalyTimeItem struct {
	WatingTime  string  `json:"wating_time"`
	ProcessTime float64 `json:"process_time"`
}
type AnomalyMachineItem struct {
	MachineID   uint   `json:"machine_id"`
	MachineName string `json:"machine_name"`
	Reason      string `json:"reason"`
	StatusType  string `json:"status_type"`
	Date        string `json:"date"`
	// Time        string `json:"time"`
}
type AnomalyTypeItem struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}
type AnomalyAnalysis struct {
	AnomalyTime    []AnomalyTimeItem    `json:"anomaly_time"`
	AnomalyMachine []AnomalyMachineItem `json:"anomaly_machine"`
	AnomalyType    []AnomalyTypeItem    `json:"anomaly_type"`
}

// 工单信息
type Order struct {
	TaskName          string
	WorkshopName      string
	StartTime         string
	EffectiveTime     string
	TotalNum          string
	FinishNum         string
	ProcedureStatusID string
	IsFinished        string
	Description       string
}
type OrderAnalysis struct {
	Orders []Order
}

type StatusCount struct {
	MachineStatusID uint   `json:"machine_status_id"`
	StatusCount     uint   `json:"status_count"`
	StatusName      string `json:"status_name"`
}

type StatusDayCount struct {
	MachineStatusID uint      `json:"machine_status_id"`
	StatusCount     uint      `json:"status_count"`
	StatusName      string    `json:"status_name"`
	Date            time.Time `json:"date"`
}

type ItemCount1 struct {
	ItemType     string  `json:"item_type"`
	TotalCount   int     `json:"total_count"`
	PerfectCount int     `json:"perfect_count"`
	Percent      float64 `json:"percent"`
}
type ItemCount2 struct {
	ItemType  string    `json:"item_type"`
	CreatedAt time.Time `json:"created_at"`
	Count     int       `json:"count"`
}

type ItemCount struct{
	ItemCount1 []ItemCount1
	ItemCount2 []ItemCount2
}

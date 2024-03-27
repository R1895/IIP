package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/serializer"
	"sort"
	"time"
)

// 车间生产状态
type ProductionBoardService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
	// TaskID     uint `form:"task_id" json:"task_id""`
}

func (service *ProductionBoardService) ProductionBoard() serializer.Response {

	// TODO 性能优化

	//首先得到workshop中所有设备
	var workshop model.Workshop
	code := e.SUCCESS
	err := model.DB.Model(model.Workshop{}).
		Preload("Machines").
		Preload("Tasks").
		First(&workshop, service.WorkshopID).Error

	var result []serializer.ProductionBoard

	for _, machine := range workshop.Machines {

		var workerName string
		sql := `
		SELECT  w.worker_name
		FROM worker AS w
		JOIN machine_worker AS mw ON w.id = mw.worker_id
		WHERE mw.machine_id = ?
		LIMIT 1
   	 	`
		err1 := model.DB.Raw(sql, machine.ID).Scan(&workerName).Error

		if err1 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		var procedures []model.Procedure
		err2 := model.DB.Model(model.Procedure{}).
			Where("machine_id = ?", machine.ID).
			Where("is_finished= ?", 2).        //表示工序任务当前未完成
			Where("procedure_status_id=?", 1). //表示工序任务当前正在进行
			Find(&procedures).Error

		if err2 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		var p model.Procedure
		var item serializer.ProductionBoard
		if len(procedures) > 0 {
			p = procedures[0]
			item = serializer.ProductionBoard{
				MachineID:       machine.ID,
				MachineName:     machine.MachineName,
				Time:            machine.UpdatedAt.String(),
				MachineStatusID: machine.MachineStatusID,
				WorkerName:      workerName,
				TaskID:          machine.CurrentTaskID,
				ProcedureID:     p.ID,
				ProcedureName:   p.ProcedureName,
				ProcedureType:   p.ProcedureType,
				Percent:         float64(p.ProcessedNum) / float64(p.WorkpiecesNum),
			}
		} else {
			item = serializer.ProductionBoard{
				MachineID:       machine.ID,
				MachineName:     machine.MachineName,
				TaskID:          machine.CurrentTaskID,
				Time:            machine.UpdatedAt.String(),
				MachineStatusID: machine.MachineStatusID,
				WorkerName:      workerName,
			}
		}

		result = append(result, item)
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}

// 生产实况
type ProductionStatusService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
}

func (service *ProductionStatusService) ProductionStatus() serializer.Response {
	var workshop model.Workshop
	code := e.SUCCESS
	err := model.DB.Model(model.Workshop{}).
		Preload("Machines").
		Preload("Tasks").
		First(&workshop, service.WorkshopID).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var statusNum serializer.StatusNum
	var oee serializer.OEE
	var machines []serializer.MachineScreen
	var anomalys []serializer.AnomalyScreen
	var count1, count2, count3, countGreater3 int64
	// machine_status_id = 1 的数量
	model.DB.Model(&model.Machine{}).Where("machine_status_id = ?", 1).Count(&count1)

	// machine_status_id = 2 的数量
	model.DB.Model(&model.Machine{}).Where("machine_status_id = ?", 2).Count(&count2)

	// machine_status_id = 3 的数量
	model.DB.Model(&model.Machine{}).Where("machine_status_id = ?", 3).Count(&count3)

	// machine_status_id > 3 的数量
	model.DB.Model(&model.Machine{}).Where("machine_status_id > ?", 3).Count(&countGreater3)
	statusNum.NumOfRun = count1
	statusNum.NumOfWaiting = count2
	statusNum.NumOfClose = count3
	statusNum.NumOfError = countGreater3

	// TODO OEE

	err = model.DB.Table("machine").
	Select("machine.id as machine_id, machine.machine_name, machine_status.status_name as machine_status, machine_status.status_type, procedure.workpieces_num, procedure.processed_num").
	Joins("left join `procedure` on machine.id = `procedure`.machine_id").
	Joins("left join machine_status on machine.machine_status_id = machine_status.id").
	Scan(&machines).Error

	for i := range machines {
		if machines[i].WorkpiecesNum != 0 {
			machines[i].Progress = (machines[i].ProcessedNum * 100) / machines[i].WorkpiecesNum
		} else {
			// 处理分母为零的情况，防止除零错误
			machines[i].Progress = 0
		}
	}
	// 构建并执行查询
	err = model.DB.Model(&model.Machine{}).
		Select("machine.id as machine_id, machine.machine_name, machine_status.status_name as reason, machine_status.status_type, machine.updated_at as date").
		Joins("left join machine_status on machine.machine_status_id = machine_status.id").
		Where("machine.machine_status_id > ?", 3).
		Scan(&anomalys).Error

	result := serializer.ProductionStatus{
		StatusNum: statusNum,
		OEE:       oee,
		Machines:  machines,
		Anomalys:  anomalys,
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}

// 设备综合效率OEE分析
type OEEAnalysisService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
}

func (service OEEAnalysisService) OEEAnalysis() serializer.Response {
	code := e.SUCCESS
	var oeeRank serializer.OEERankItems
	var oeeRankpercent serializer.OEERankItems
	var oeeLoss []serializer.OEELossItem
	var oeeReason []serializer.OEEReasonItem
	var oeeTrend []serializer.OEETrendItem

	type MachineInfo struct {
		MachineID       uint64
		Time            time.Time
		MachineStatusID uint64
	}
	var machineInfos []MachineInfo

	err := model.DB.Model(&model.Machine{}).
		Select("machine.id as machine_id, machine_log.time, machine_log.machine_status_id").
		Joins("left join machine_log on machine.id = machine_log.machine_id").
		Where("machine.workshop_id = ?", service.WorkshopID).
		Scan(&machineInfos).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 1、OEE设备排名，各状态的所占比
	// total := len(machineInfos)
	counts := make(map[uint64]map[uint64]int)
	// 遍历 machineInfos 并更新计数
	for _, mi := range machineInfos {
		if _, ok := counts[mi.MachineID]; !ok {
			counts[mi.MachineID] = make(map[uint64]int)
		}
		counts[mi.MachineID][mi.MachineStatusID]++
	}
	for machineID, statusCounts := range counts {
		item := serializer.OEERankItem{
			MachineID: machineID,
		}
		for statusID, count := range statusCounts {
			if statusID == 1 {
				item.Run = float64(count)
			}
			if statusID == 2 {
				item.Wating = float64(count)
			}
			if statusID == 3 {
				item.Closed = float64(count)
			}
			if statusID > 3 {
				item.Error += float64(count)
			}
		}
		oeeRank = append(oeeRank, item)
	}
	sort.Slice(oeeRank, func(i, j int) bool {
		return oeeRank[i].MachineID < oeeRank[j].MachineID
	})

	// 计算百分比并添加到 oeeRank
	for machineID, statusCounts := range counts {
		item := serializer.OEERankItem{
			MachineID: machineID,
		}
		total := int64(0)
		for _, oeeItem := range oeeRank {
			if oeeItem.MachineID == machineID {
				item = oeeItem
				break
			}
		}
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		total = int64(item.Closed) + int64(item.Run) + int64(item.Wating) + int64(item.Error)
		for statusID, count := range statusCounts {
			if total != 0 {
				percentage := float64(count) / float64(total)

				switch statusID {
				case 1:
					item.Run = percentage
				case 2:
					item.Wating = percentage
				case 3:
					item.Closed = percentage
				default:
				}
			}
			item.Error = 1 - item.Run - item.Wating - item.Closed
		}

		oeeRankpercent = append(oeeRankpercent, item)
	}
	sort.Slice(oeeRankpercent, func(i, j int) bool {
		return oeeRankpercent[i].Run > oeeRankpercent[j].Run
	})
	//类型-数量
	var statusCounts []serializer.StatusCount
	err2 := model.DB.Model(&model.MachineAnomaly{}).
		Select("machine_anomaly.machine_status_id, COUNT(*) as status_count, machine_status.status_name").
		Joins("JOIN machine_status ON machine_anomaly.machine_status_id = machine_status.id").
		Where("machine_anomaly.machine_status_id IN (5, 6, 7, 8, 9)").
		Group("machine_anomaly.machine_status_id, machine_status.status_name").
		Scan(&statusCounts).Error
	if err2 != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	sort.Slice(statusCounts, func(i, j int) bool {
		return statusCounts[i].MachineStatusID < statusCounts[j].MachineStatusID
	})

	//类型-数量-日期

	var statusDayCounts []serializer.StatusDayCount

	err3 := model.DB.Model(&model.MachineAnomaly{}).
		Select("DATE(time) as date, machine_anomaly.machine_status_id, COUNT(*) as status_count, machine_status.status_name").
		Joins("JOIN machine_status ON machine_anomaly.machine_status_id = machine_status.id").
		Where("machine_anomaly.machine_status_id IN (5, 6, 7, 8, 9)").
		Group("date, machine_anomaly.machine_status_id, machine_status.status_name").
		Scan(&statusDayCounts).Error

	if err3 != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	sort.Slice(statusDayCounts, func(i, j int) bool {
		return statusDayCounts[i].MachineStatusID < statusDayCounts[j].MachineStatusID
	})

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 2、设备运行时长每日损失量

	// 3、设备OEE综合趋势

	// 4、设备运行时间损失原因分析

	result := serializer.OEEAnalysis{
		OEERank:           oeeRank,
		OEERankPercent:    oeeRankpercent,
		OEEStatusCount:    statusCounts,
		OEEStatusDayCount: statusDayCounts,
		OEELoss:           oeeLoss,
		OEEReason:         oeeReason,
		OEETrend:          oeeTrend,
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   result,
	}
}

// 异常分析
type AnomalyAnalysisService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
}

func (service AnomalyAnalysisService) AnomalyAnalysis() serializer.Response {
	code := e.SUCCESS
	workshopid := service.WorkshopID
	var machines []model.Machine

	//获得workshopid对应的machine列表
	err := model.DB.Model(&model.Machine{}).
		Where("workshop_id = ?", workshopid).
		Find(&machines).
		Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var counts []int64
	var count int64
	for i := 5; i <= 14; i++ {
		err := model.DB.Model(&machines).Where("machine_status_id = ?", i).Count(&count).Error
		counts = append(counts, count)

		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	// 查询MachineAnomaly表，获取每台机器的异常记录

	var anomalies1 []model.MachineAnomaly
	for _, machine := range machines {
		var tmpanomaly model.MachineAnomaly
		err := model.DB.Model(&model.MachineAnomaly{}).
			Where("machine_id = ?", machine.ID). //TODO 去掉for循环
			Order("time DESC").
			Limit(1).
			Find(&tmpanomaly).
			Error

		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		anomalies1 = append(anomalies1, tmpanomaly)
	}

	var anomalies2 []model.MachineAnomaly
	for _, anomaly := range anomalies1 {
		// 如果machine_status_id的值在10到14之间，就往前查找一个状态
		if anomaly.MachineStatusID >= 10 && anomaly.MachineStatusID <= 14 {
			var previousAnomaly model.MachineAnomaly
			// 查询前一个状态
			err := model.DB.Model(anomalies1).
				Where("machine_id = ? AND time < ?", anomaly.MachineID, anomaly.Time).
				Order("time DESC").
				Limit(1).
				Find(&previousAnomaly).
				Error

			if err != nil {
				code = e.ErrorDatabase
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
					Error:  err.Error(),
				}
			}

			// 如果找到了前一个状态，将其添加到anomalies2中
			if previousAnomaly.ID != 0 {
				anomalies2 = append(anomalies2, previousAnomaly)
			}
		}
	}

	//计算时间
	var durations1 []time.Duration
	for i := 0; i <= len(anomalies1)-1; i++ {
		duration := time.Now().Sub(anomalies1[i].Time)
		durations1 = append(durations1, duration)
	}

	var durations2 []time.Duration
	var correspondingTime time.Time
	for i := 0; i <= len(anomalies2)-1; i++ {
		// 使用与anomalies1中相同machine_id对应的时间
		for _, anomaly := range anomalies1 {
			if anomaly.MachineID == anomalies2[i].MachineID {
				correspondingTime = anomaly.Time
			}
		}
		if !correspondingTime.IsZero() {
			duration := anomalies2[i].Time.Sub(correspondingTime)
			durations2 = append(durations2, duration)
		}
	}
	statanomaly := serializer.StatAnomaly{
		Anomalys1: anomalies1,
		Anomalys2: anomalies2,
		Duration1: durations1,
		Duration2: durations2,
		Count:     counts,
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   statanomaly,
	}
}

// 工单信息
type OrderAnalysisService struct {
	WorkshopID int `form:"workshop_id" json:"workshop_id" binding:"required"`
	Limit      int `form:"limit" json:"limit"`
	Start      int `form:"start" json:"start" binding:"required"`
}

func (service OrderAnalysisService) OrderAnalysis() serializer.Response {
	var orders []serializer.Order
	var orderAnalysis serializer.OrderAnalysis
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}

	tx := model.DB.Model(&model.Task{}).
		Select("task.task_name, task.start_date as start_time, task.effective_time, workshop.workshop_name, procedure.workpieces_num as total_num, procedure.processed_num as finish_num,procedure.procedure_status_id , task.is_finished, task.description").
		Joins("left join workshop on task.workshop_id = workshop.id").
		Joins("left join `procedure` on task.id = `procedure`.task_id")

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).Scan(&orders)

	orderAnalysis.Orders = orders
	//如果这种写法不行就需要添加BuildListOrder函数
	return serializer.BuildListResponse(orderAnalysis, uint(total))

}

// 工序甘特图
type ProcedureGanttService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
}

func (service ProcedureGanttService) ProcedureGantt() serializer.Response {

	err := e.SUCCESS
	return serializer.Response{
		Status: err,
		Msg:    e.GetMsg(err),
	}
}

// 工时排行 本月、年度、今日
type WorkerAttendanceService struct {
	WorkShopID uint   `form:"workshop_id" json:"workshop_id"`
	StartDate  string `form:"start_date"  json:"start_date"  binding:"required"`
	EndDate    string `form:"end_date" json:"end_date" binding:"required"`
}

// 现在返回的是多少秒
func (service *WorkerAttendanceService) WorkerAttendance() serializer.Response {

	var workerattendances []model.WorkerAttendance
	code := e.SUCCESS
	startdate := service.StartDate
	enddate := service.EndDate
	workshopid := service.WorkShopID
	var statattendances []model.StatAttendance

	// TODO 可以先查workshop， workshop.workers
	rows, err := model.DB.Raw("SELECT worker_id FROM worker_workshop WHERE workshop_id =?", workshopid).Rows()
	defer rows.Close()
	var workerids []uint
	for rows.Next() {
		var id uint
		rows.Scan(&id)
		workerids = append(workerids, id)
	}

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Model(model.WorkerAttendance{}).
		Where(`date(date)>=?`, startdate).
		Where(`date(date)<=?`, enddate).
		Where("worker_id IN (?)", workerids).
		Order("worker_id ASC").
		Find(&workerattendances).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Model(&model.WorkerAttendance{}).
		Select("worker_attendance.worker_id, SEC_TO_TIME(SUM(TIMESTAMPDIFF(SECOND, worker_attendance.arrival_time, worker_attendance.departure_time))) as total_time, worker.worker_name").
		Joins("JOIN worker ON worker_attendance.worker_id = worker.id").
		Where("worker_attendance.type IS NOT NULL").
		Where(`date(date)>=?`, startdate).
		Where(`date(date)<=?`, enddate).
		Group("worker_attendance.worker_id, worker.worker_name").
		Order("total_time DESC").
		Find(&statattendances).Error

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildStatAttendances(statattendances),
	}
}

// 库存管理
type StatIentoryService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start" binding:"required"`
}

func (service *StatIentoryService) StatIentory() serializer.Response {
	code := e.SUCCESS
	var itemTypes []string
	var total int64
	err := model.DB.Model(&model.Iventory{}).
		Distinct("item_type").
		Pluck("item_type", &itemTypes).
		Count(&total).
		Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
		}
	}
	var statIventorys []serializer.StatIventory
	start := (service.Start - 1) * service.Limit
	end := start + service.Limit
	if end > int(total) {
		end = int(total)
	}
	itemTypes = itemTypes[start:end]
	for _, itemType := range itemTypes {
		var tmpiventory model.Iventory
		var iventorysUp []model.Iventory
		var iventorysDown []model.Iventory

		err1 := model.DB.Model(model.Iventory{}).Where(`item_type = ?`, itemType).Limit(1).Find(&tmpiventory).Error
		if err1 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
			}
		}
		field1 := tmpiventory.Field1

		var up, down uint
		err1 = model.DB.Model(model.Iventory{}).Where(`item_type = ?`, itemType).Where(`iventory_type = ?`, 1).Find(&iventorysUp).Error
		if err1 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
			}
		}

		err2 := model.DB.Model(model.Iventory{}).Where(`item_type = ?`, itemType).Where(`iventory_type = ?`, 2).Find(&iventorysDown).Error
		if err2 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
			}
		}

		if len(iventorysUp) > 0 {
			up = 0
			for _, iventory := range iventorysUp {
				up += iventory.IventoryNum
			}
		}
		if len(iventorysDown) > 0 {
			down = 0

			for _, iventory := range iventorysDown {
				down += iventory.IventoryNum
			}
		}

		statIventory := serializer.StatIventory{
			Field1:   field1,
			ItemType: itemType,
			StatNum:  up - down,
		}

		// 将计算得到的 StatIventory 添加到 statIventrys 中
		statIventorys = append(statIventorys, statIventory)
	}

	stat := serializer.StatIventorys{
		StatIventorys: statIventorys,
		Total:         total,
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   stat,
	}
}

// 产能 优品率
type ProductivityService struct {
	WorkshopID uint
}

func (service ProductivityService) Productivity() serializer.Response {
	code := e.SUCCESS
	var itemCounts1 []serializer.ItemCount1
	//TODO 未加worksopid，只需要在搜索中现根据workshop_id找到相应的machine_id再在查询中加入where判断即可
	err := model.DB.Model(&model.ProcedureItem{}).
		Select("item_type, COUNT(*) as total_count, SUM(CASE WHEN quality = 'perfect' THEN 1 ELSE 0 END) as perfect_count, AVG(CASE WHEN quality = 'perfect' THEN 100.0 ELSE 0 END) as percent").
		Group("item_type").
		Scan(&itemCounts1).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
		}
	}
	var itemCounts2 []serializer.ItemCount2

	err = model.DB.Model(&model.ProcedureItem{}).
		Select("item_type, DATE(created_at) as created_at, COUNT(*) as count").
		Group("item_type, DATE(created_at)").
		Scan(&itemCounts2).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
		}
	}
	var itemCount serializer.ItemCount
	itemCount.ItemCount1 = itemCounts1
	itemCount.ItemCount2 = itemCounts2
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   itemCount,
	}
}

// 车间的 环境、气味、噪音、安全 信息
type NumDayEnvironmentService struct {
	WorkshopId uint   `form:"workshop_id" json:"workshop_id"  binding:"required" `
	StartDate  string `form:"start_date"  json:"start_date"  binding:"required"`
	EndDate    string `form:"end_date" json:"end_date" binding:"required"`
}

func (service *NumDayEnvironmentService) NumDay() serializer.Response {
	var environments []model.Environment
	code := e.SUCCESS
	startdate := service.StartDate
	enddate := service.EndDate
	workshopid := service.WorkshopId
	// TODO 时间
	err := model.DB.Model(model.Environment{}).
		Where(`date(date)>=?`, startdate).
		Where(`date(date)<=?`, enddate).
		Where("workshop_id= ?", workshopid).
		Order("date ASC").Find(&environments).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	count := len(environments)

	return serializer.BuildListResponse(serializer.BuildEnvironments(environments), uint(count))
}

type StatDaySmellService struct {
	WorkshopID uint   `form:"workshop_id" json:"workshop_id"  binding:"required" `
	StartDate  string `form:"start_date"  json:"start_date"  binding:"required"`
	EndDate    string `form:"end_date" json:"end_date" binding:"required"`
}

func (service *StatDaySmellService) StatDay() serializer.Response {
	var smells []model.Smell
	code := e.SUCCESS
	startdate := service.StartDate
	enddate := service.EndDate
	workshopid := service.WorkshopID

	err := model.DB.Model(model.Smell{}).
		Where(`date(date)>=?`, startdate).
		Where(`date(date)<=?`, enddate).
		Where("workshop_id= ?", workshopid).
		Order("date ASC").Find(&smells).Error
	if err != nil {
		//
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var sum float64
	var max, min float64

	if len(smells) > 0 {
		sum = 0
		max = smells[0].Value
		min = smells[0].Value

		for _, smell := range smells {
			sum += smell.Value

			if smell.Value > max {
				max = smell.Value
			}

			if smell.Value < min {
				min = smell.Value
			}
		}
	} else {
		// 处理当噪声数据为空的情况
		sum = 0
		max = 0
		min = 0
	}
	// 计算平均值
	avg := sum / float64(len(smells))
	count := len(smells)
	// 构建响应
	smellStats := serializer.DaySmell{
		Workshop_ID: workshopid,
		Min:         min,
		Max:         max,
		Avg:         avg,
		Count:       uint(count),
		Smells:      serializer.BuildSmells(smells),
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   smellStats,
	}
}

type StatDayNoiseService struct {
	WorkshopID uint   `form:"workshop_id" json:"workshop_id"  binding:"required" `
	StartDate  string `form:"start_date"  json:"start_date"  binding:"required"`
	EndDate    string `form:"end_date" json:"end_date" binding:"required"`
}

func (service *StatDayNoiseService) StatDay() serializer.Response {

	var noises []model.Noise
	code := e.SUCCESS
	startdate := service.StartDate
	enddate := service.EndDate
	workshopid := service.WorkshopID

	err := model.DB.Model(model.Noise{}).
		Where(`date(date)>=?`, startdate).
		Where(`date(date)<=?`, enddate).
		Where("workshop_id= ?", workshopid).
		Order("date ASC").Find(&noises).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var sum float64
	var max, min float64

	if len(noises) > 0 {
		sum = 0
		max = noises[0].Value
		min = noises[0].Value

		for _, noise := range noises {
			sum += noise.Value

			if noise.Value > max {
				max = noise.Value
			}

			if noise.Value < min {
				min = noise.Value
			}
		}
	} else {
		// 处理当噪声数据为空的情况
		sum = 0
		max = 0
		min = 0
	}
	// 计算平均值
	avg := sum / float64(len(noises))
	count := len(noises)
	// 构建响应
	noiseStats := serializer.NoiseStats{
		Workshop_ID: workshopid,
		Min:         min,
		Max:         max,
		Avg:         avg,
		Count:       uint(count),
		Noises:      serializer.BuildNoises(noises),
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   noiseStats,
	}

}

type NumDaySafeService struct {
	WorkshopID uint   `form:"workshop_id" json:"workshop_id"  binding:"required" `
	StartDate  string `form:"start_date"  json:"start_date"  binding:"required"`
	EndDate    string `form:"end_date" json:"end_date" binding:"required"`
}

func (service *NumDaySafeService) NumDay() serializer.Response {
	var safeinformations []model.SafeInformation
	code := e.SUCCESS
	startdate := service.StartDate
	enddate := service.EndDate
	workshopid := service.WorkshopID
	var count uint
	var total int64
	err := model.DB.Model(model.Safe{}).
		Where(`date(date)>=?`, startdate).
		Where(`date(date)<=?`, enddate).
		Where("workshop_id= ?", workshopid).
		Select("type, SUM(value) as total_value").
		Group("type").
		Order("type ASC").
		Count(&total).
		Find(&safeinformations).Error

	if err != nil {
		//
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
		}
	}
	if len(safeinformations) > 0 {
		count = 0

		for _, safeinformation := range safeinformations {
			count += uint(safeinformation.TotalValue)
		}
	}

	numDay := serializer.NumDaySafe{
		Workshop_ID:      workshopid,
		Total:            uint(total),
		SafeInformations: serializer.BuildSafeInformations(safeinformations),
		Count:            count,
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   numDay,
	}
}

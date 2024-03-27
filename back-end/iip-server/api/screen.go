package api

import (
	"iip/pkg/util"
	"iip/service"

	"github.com/gin-gonic/gin"
)

// 车间生产状态
func ProductionBoard(c *gin.Context) {
	productionBoardService := service.ProductionBoardService{}

	if err := c.ShouldBind(&productionBoardService); err == nil {
		res := productionBoardService.ProductionBoard()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 生产实况
func ProductionStatus(c *gin.Context) {
	productionStatusService := service.ProductionStatusService{}

	if err := c.ShouldBind(&productionStatusService); err == nil {
		res := productionStatusService.ProductionStatus()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 设备综合效率OEE
func OEEAnalysis(c *gin.Context) {
	oeeAnalysisService := service.OEEAnalysisService{}

	if err := c.ShouldBind(&oeeAnalysisService); err == nil {
		res := oeeAnalysisService.OEEAnalysis()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 异常分析
func AnomalyAnalysis(c *gin.Context) {
	anomalyAnalysisService := service.AnomalyAnalysisService{}

	if err := c.ShouldBind(&anomalyAnalysisService); err == nil {
		res := anomalyAnalysisService.AnomalyAnalysis()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 工单信息
func OrderAnalysis(c *gin.Context) {
	orderAnalysisService := service.OrderAnalysisService{}

	if err := c.ShouldBind(&orderAnalysisService); err == nil {
		res := orderAnalysisService.OrderAnalysis()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 工序甘特图
func ProcedureGantt(c *gin.Context) {
	procedureGanttService := service.ProcedureGanttService{}

	if err := c.ShouldBind(&procedureGanttService); err == nil {
		res := procedureGanttService.ProcedureGantt()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 工时排行
func WorkerAttendance(c *gin.Context) {
	workerAttendanceService := service.WorkerAttendanceService{}

	if err := c.ShouldBind(&workerAttendanceService); err == nil {
		res := workerAttendanceService.WorkerAttendance()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}



// 库存管理
func IentoryAnalysis(c *gin.Context) {
	statIentoryService := service.StatIentoryService{}

	if err := c.ShouldBind(&statIentoryService); err == nil {
		res := statIentoryService.StatIentory()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 车间信息
// 环境
func EnvironmentAnalysis(c *gin.Context) {
	numDayEnvironmentService := service.NumDayEnvironmentService{}

	if err := c.ShouldBind(&numDayEnvironmentService); err == nil {
		res := numDayEnvironmentService.NumDay()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 气味
func SmellAnalysis(c *gin.Context) {
	statDaySmellService := service.StatDaySmellService{}

	if err := c.ShouldBind(&statDaySmellService); err == nil {
		res := statDaySmellService.StatDay()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 噪音 取最大最小平均值 stat 某一天
func NoiseAnalysis(c *gin.Context) {
	statDayNoiseService := service.StatDayNoiseService{}

	if err := c.ShouldBind(&statDayNoiseService); err == nil {
		res := statDayNoiseService.StatDay()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// 安全
func SafeAnalysis(c *gin.Context) {
	numDaySafeService := service.NumDaySafeService{}

	if err := c.ShouldBind(&numDaySafeService); err == nil {
		res := numDaySafeService.NumDay()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func Productivity(c *gin.Context) {
	productivityService := service.ProductivityService{}

	if err := c.ShouldBind(&productivityService); err == nil {
		res := productivityService.Productivity()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
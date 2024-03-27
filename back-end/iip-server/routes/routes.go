package routes

import (
	"iip/api"
	"iip/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 跨域
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
		// Token 鉴权相关
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			user := authed.Group("user")
			{
				user.POST("", api.CreateUser)
				user.GET("/users", api.ListUsers)
				user.GET("/:id", api.ShowUser)
				user.DELETE("/:id", api.DeleteUser)
				user.PUT("/:id", api.UpdateUser)
			}
			// 角色操作
			role := authed.Group("role")
			{
				role.GET("/:id", api.ShowRole)
				role.GET("/roles", api.ListRoles)
				role.POST("", api.CreateRole)
				role.DELETE("/:id", api.DeleteRole)
				role.PUT("/:id", api.UpdateRole)
			}
			// 权限操作
			permission := authed.Group("permission")
			{
				permission.GET("/:id", api.ShowPermission)
				permission.GET("/permissions", api.ListPermissions)
				permission.POST("", api.CreatePermission)
				permission.DELETE("/:id", api.DeletePermission)
				permission.PUT("/:id", api.UpdatePermission)
			}
			// 车间操作
			workshop := authed.Group("workshop")
			{
				workshop.GET("/:id", api.ShowWorkshop)
				workshop.GET("/workshops", api.ListWorkshops)
				workshop.POST("", api.CreateWorkshop)
				workshop.DELETE("/:id", api.DeleteWorkshop)
				workshop.PUT("/:id", api.UpdateWorkshop)
				workshop.POST("/worker", api.AddWorker)      // 车间添加工人
				workshop.DELETE("/worker", api.RemoveWorker) // 车间删除工人
			}
			// 工人类别
			workerType := authed.Group("workerType")
			{
				workerType.GET("/:id", api.ShowWorkerType)
				workerType.GET("/workerTypes", api.ListWorkerTypes)
				workerType.POST("", api.CreateWorkerType)
				workerType.DELETE("/:id", api.DeleteWorkerType)
				workerType.PUT("/:id", api.UpdateWorkerType)
			}
			// 工人操作
			worker := authed.Group("worker")
			{
				worker.GET("/:id", api.ShowWorker)
				worker.GET("/workers", api.ListWorkers)
				worker.POST("", api.CreateWorker)
				worker.DELETE("/:id", api.DeleteWorker)
				worker.PUT("/:id", api.UpdateWorker)
			}
			// 机器设备操作
			machine := authed.Group("machine")
			{
				machine.GET("/:id", api.ShowMachine)
				machine.GET("/machines", api.ListMachines)
				machine.POST("", api.CreateMachine)
				machine.DELETE("/:id", api.DeleteMachine)
				machine.PUT("/:id", api.UpdateMachine)
				machine.GET("/status/:mid/:sid", api.ModifyMachineStatus)
			}
			// 机器设备异常
			machineAnomaly := authed.Group("machineAnomaly")
			{
				machineAnomaly.GET("/anomaly/:mid/:id", api.CreateMachineAnomaly2)
				machineAnomaly.GET("/:id", api.ShowMachineAnomaly)
				machineAnomaly.GET("/machineAnomalys", api.ListMachineAnomalys)
				machineAnomaly.POST("", api.CreateMachineAnomaly)
				machineAnomaly.DELETE("/:id", api.DeleteMachineAnomaly)
				machineAnomaly.PUT("/:id", api.UpdateMachineAnomaly)
			}
			// 机器分配工人
			machineWorker := authed.Group("machineWorker")
			{
				machineWorker.GET("/:id", api.ShowMachineWorker)
				machineWorker.GET("/machineWorkers", api.ListMachineWorkers)
				machineWorker.POST("", api.CreateMachineWorker)
				machineWorker.DELETE("/:id", api.DeleteMachineWorker)
				machineWorker.PUT("/:id", api.UpdateMachineWorker)
			}
			// 机器状态管理
			machineStatus := authed.Group("machineStatus")
			{
				machineStatus.GET("/:id", api.ShowMachineStatus)
				machineStatus.GET("/machineStatuss", api.ListMachineStatuss)
				machineStatus.POST("", api.CreateMachineStatus)
				machineStatus.DELETE("/:id", api.DeleteMachineStatus)
				machineStatus.PUT("/:id", api.UpdateMachineStatus)
			}
			// 机器日志
			machineLog := authed.Group("machineLog")
			{
				machineLog.GET("/:id", api.ShowMachineLog)
				machineLog.GET("/machineLogs", api.ListMachineLogs)
				machineLog.POST("", api.CreateMachineLog)
				machineLog.DELETE("/:id", api.DeleteMachineLog)
				machineLog.PUT("/:id", api.UpdateMachineLog)
			}
			// 任务/产线管理
			task := authed.Group("task")
			{
				task.GET("/:id", api.ShowTask)
				task.GET("/tasks", api.ListTasks)
				task.POST("", api.CreateTask)
				task.DELETE("/:id", api.DeleteTask)
				task.PUT("/:id", api.UpdateTask)
			}
			// 工时管理
			workerAttendance := authed.Group("workerAttendance")
			{
				workerAttendance.GET("/:id", api.ShowWorkerAttendance)
				workerAttendance.GET("/workerAttendances", api.ListWorkerAttendances)
				workerAttendance.POST("", api.CreateWorkerAttendance)
				workerAttendance.DELETE("/:id", api.DeleteWorkerAttendance)
				workerAttendance.PUT("/:id", api.UpdateWorkerAttendance)
			}
			// 消息管理
			message := authed.Group("message")
			{
				message.GET("/:id", api.ShowMessage)
				message.GET("/messages", api.ListMessages)
				message.POST("", api.CreateMessage)
				message.DELETE("/:id", api.DeleteMessage)
				message.PUT("/:id", api.UpdateMessage)
			}
			// 数据字典管理
			dictionaryItem := authed.Group("dictionaryItem")
			{
				dictionaryItem.GET("/:id", api.ShowDictionaryItem)
				dictionaryItem.GET("/dictionaryItems", api.ListDictionaryItems)
				dictionaryItem.POST("", api.CreateDictionaryItem)
				dictionaryItem.DELETE("/:id", api.DeleteDictionaryItem)
				dictionaryItem.PUT("/:id", api.UpdateDictionaryItem)
			}
			// 用户日志
			log := authed.Group("log")
			{
				log.GET("/:id", api.ShowLog)
				log.GET("/logs", api.ListLogs)
				log.POST("", api.CreateLog)
				log.DELETE("/:id", api.DeleteLog)
				log.PUT("/:id", api.UpdateLog)
			}
			// 工厂噪音
			noise := authed.Group("noise")
			{
				noise.GET("/:id", api.ShowNoise)
				noise.GET("/noises", api.ListNoises)
				noise.POST("", api.CreateNoise)
				noise.DELETE("/:id", api.DeleteNoise)
				noise.PUT("/:id", api.UpdateNoise)
			}
			// 工厂气味
			smell := authed.Group("smell")
			{
				smell.GET("/:id", api.ShowSmell)
				smell.GET("/smells", api.ListSmells)
				smell.POST("", api.CreateSmell)
				smell.DELETE("/:id", api.DeleteSmell)
				smell.PUT("/:id", api.UpdateSmell)
			}
			// 工厂安全
			safe := authed.Group("safe")
			{
				safe.GET("/:id", api.ShowSafe)
				safe.GET("/safes", api.ListSafes)
				safe.POST("", api.CreateSafe)
				safe.DELETE("/:id", api.DeleteSafe)
				safe.PUT("/:id", api.UpdateSafe)
			}
			// 工厂环境
			environment := authed.Group("environment")
			{
				environment.GET("/:id", api.ShowEnvironment)
				environment.GET("/environments", api.ListEnvironments)
				environment.POST("", api.CreateEnvironment)
				environment.DELETE("/:id", api.DeleteEnvironment)
				environment.PUT("/:id", api.UpdateEnvironment)
			}
			// 工序管理
			procedureType := authed.Group("procedureType")
			{
				procedureType.GET("/:id", api.ShowProcedureType)
				procedureType.GET("/procedureTypes", api.ListProcedureTypes)
				procedureType.POST("", api.CreateProcedureType)
				procedureType.DELETE("/:id", api.DeleteProcedureType)
				procedureType.PUT("/:id", api.UpdateProcedureType)
			}
			// 工件管理
			procedureItem := authed.Group("procedureItem")
			{
				procedureItem.GET("/:id", api.ShowProcedureItem)
				procedureItem.GET("/procedureItems", api.ListProcedureItems)
				procedureItem.POST("", api.CreateProcedureItem)
				procedureItem.DELETE("/:id", api.DeleteProcedureItem)
				procedureItem.PUT("/:id", api.UpdateProcedureItem)
			}
			// 工序状态
			procedureStatus := authed.Group("procedureStatus")
			{
				procedureStatus.GET("/:id", api.ShowProcedureStatus)
				procedureStatus.GET("/procedureStatuss", api.ListProcedureStatuss)
				procedureStatus.POST("", api.CreateProcedureStatus)
				procedureStatus.DELETE("/:id", api.DeleteProcedureStatus)
				procedureStatus.PUT("/:id", api.UpdateProcedureStatus)
			}
			// 工序管理
			procedure := authed.Group("procedure")
			{
				procedure.GET("/:id", api.ShowProcedure)
				procedure.GET("/procedures", api.ListProcedures)
				procedure.POST("", api.CreateProcedure)
				procedure.DELETE("/:id", api.DeleteProcedure)
				procedure.PUT("/:id", api.UpdateProcedure)
			}
			// 任务状态
			taskStatus := authed.Group("taskStatus")
			{
				taskStatus.GET("/:id", api.ShowTaskStatus)
				taskStatus.GET("/taskStatuss", api.ListTaskStatuss)
				taskStatus.POST("", api.CreateTaskStatus)
				taskStatus.DELETE("/:id", api.DeleteTaskStatus)
				taskStatus.PUT("/:id", api.UpdateTaskStatus)
			}
			// 任务类别
			taskType := authed.Group("taskType")
			{
				taskType.GET("/:id", api.ShowTaskType)
				taskType.GET("/taskTypes", api.ListTaskTypes)
				taskType.POST("", api.CreateTaskType)
				taskType.DELETE("/:id", api.DeleteTaskType)
				taskType.PUT("/:id", api.UpdateTaskType)
			}
			// 库存管理
			inventory := authed.Group("inventory")
			{
				inventory.GET("/:id", api.ShowIventory)
				inventory.GET("/inventorys", api.ListIventorys)
				inventory.POST("", api.CreateIventory)
				inventory.DELETE("/:id", api.DeleteIventory)
				inventory.PUT("/:id", api.UpdateIventory)
			}
		}

		screen := v1.Group("/screen")
		{
			// 车间生产状态
			screen.POST("/productionBoard", api.ProductionBoard)
			// 生产实况
			screen.POST("/productionStatus", api.ProductionStatus)
			// 设备综合效率OEE
			screen.POST("/oee", api.OEEAnalysis)
			// 异常分析
			screen.POST("/anomaly", api.AnomalyAnalysis)
			// 工单信息
			screen.POST("/order", api.OrderAnalysis)
			// 工序甘特图
			screen.POST("/procedure", api.ProcedureGantt)
			// 工时排行
			screen.POST("/workerAttendance", api.WorkerAttendance)
			// 质量分析 优品率
			screen.POST("/productivity", api.Productivity)
			// 库存
			screen.POST("/iventory", api.IentoryAnalysis)
			// 车间信息 环境 气味 噪音 安全
			screen.POST("/noise", api.NoiseAnalysis)
			screen.POST("/smell", api.SmellAnalysis)
			screen.POST("/safe", api.SafeAnalysis)
			screen.POST("/environment", api.EnvironmentAnalysis)
		}

	}
	return r
}

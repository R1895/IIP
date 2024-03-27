package model

// 执行数据迁移
func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
		// &User{},
		// &Role{},
		// &Permission{},
		// &Message{},
		// &WorkerType{},
		// &Worker{},
		// &Workshop{},
		// &Noise{},
		// &Smell{},
		// &Safe{},
		// &Environment{},
		// &WorkerAttendance{},
		// &MachineWorker{},

		// &MachineStatus{},
		// &Machine{},
		// &MachineAnomaly{},
		// &MachineLog{},
		// &TaskStatus{},
		// &TaskType{},
		// &Task{},
		// &ProcedureItem{},
		// &ProcedureStatus{},
		// &ProcedureType{},
		// &Procedure{},
		// &DictionaryItem{},
		// &Log{},
		// &Iventory{},
		)
	if err != nil {
		return
	}

	// DB.Model(&CreditCard{}).AddForeignKey("pid", "People(id)", "CASCADE", "CASCADE")
}

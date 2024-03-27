import http from "@/utils/request";
import request from "@/utils/request";
// 用户管理，权限管理，角色管理

// 1.设备管理
// 查询
export function getDeviceManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/machine/machines`;
  } else {
    url = `/machine/${params.id}`;
  }
  console.log(url);
  return request({
    url: url,
    method: "get",
    params
  });
}
// 添加
export function addDeviceManagerCol(params) {
  return request({
    url: `/machine`,
    method: "post",
    data: params,
  });
}
// 更新
export function updateDeviceManagerCol(params) {
  return request({
    url: `/machine/${params.id}`,
    method: "put",
    data:params,
  });
}
// 删除
export function deleteDeviceManagerCol(params) {
  return request({
    url: `/machine/${params.id}`,
    method: "delete",
  });
}

// 2.车间管理
// 查询
export function getWorkshopManagerCols(params) {
  let url;
  console.log(params.id)
  if(params.id==void 0 ||params.id==""){
    url = `/workshop/workshops`
  } else {

    url = `/workshop/${params.id}`
  }
  return request({
    url: url,
    method: "get",
    params,
  });
}
// 添加
export function addWorkshopManagerCol(params) {
  return request({
    url: `/workshop`,
    method: "post",
    params,
  });
}
// 更新
export function updateWorkshopManagerCol(params) {
  return request({
    url: `/workshop/${params.id}`,
    method: "put",
    params,
  });
}
// 删除
export function deleteWorkshopManagerCol(params) {
  return request({
    url: `/workshop/${params.id}`,
    method: "delete",
  });
}

// 查询
export function getMachineAnomalyManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/machineAnomaly/machineAnomalys?limit=${params.limit}&start=${params.start}`;
  } else {
    url = `/machineAnomaly/${params.id}`;
  }
  return request({
    url: url,
    method: "get",
  });
}
// 添加
export function addMachineAnomalyManagerCol(params) {
  return request({
    url: `/machineAnomaly`,
    method: "post",
    data: params,
  });
}
// 更新
export function updateMachineAnomalyManagerCol(params) {
  return request({
    url: `/machineAnomaly/${params.id}`,
    method: "put",
    params,
  });
}
// 删除
export function deleteMachineAnomalyManagerCol(params) {
  return request({
    url: `/machineAnomaly/${params.id}`,
    method: "delete",
  });
}

// 工序类型
// 查询
export function getProceduretypeManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/procedureType/procedureTypes?limit=${params.limit}&start=${params.start}`;
  } else {
    url = `/procedureType/${params.id}`;
  }
  return request({
    url: url,
    method: "get",
  });
}
// 添加
export function addProceduretypeManagerCol(params) {
  return request({
    url: `/procedureType`,
    method: "post",
    params,
  });
}
// 更新
export function updateProceduretypeManagerCol(params) {
  return request({
    url: `/procedureType/${params.id}`,
    method: "put",
    data: params,
  });
}
// 删除
export function deleteProceduretypeManagerCol(params) {
  return request({
    url: `/procedureType/${params.id}`,
    method: "delete",
  });
}

// 工序

// 查询
export function getProcedureManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/procedure/procedures`;
  } else {
    url = `/procedure/${params.id}`;
  }
  return request({
    url: url,
    method: "get",
    params
  });
}
// 添加
export function addProcedureManagerCol(params) {
  return request({
    url: `/procedure`,
    method: "post",
    data:params,
  });
}
// 更新
export function updateProcedureManagerCol(params) {
  return request({
    url: `/procedure/${params.id}`,
    method: "put",
    data: params,
  });
}
// 删除
export function deleteProcedureManagerCol(params) {
  return request({
    url: `/procedure/${params.id}`,
    method: "delete",
  });
}

export function getMachineStatusManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/machineStatus/machineStatuss?limit=${params.limit}&start=${params.start}`;
  } else {
    url = `/machineStatus/${params.id}`;
  }
  console.log(url);
  return request({
    url: url,
    method: "get",
  });
}
// 添加
export function addMachineStatusManagerCol(params) {
  return request({
    url: `/machineStatus`,
    method: "post",
    data: params,
  });
}
// 更新
export function updateMachineStatusManagerCol(params) {
  return request({
    url: `/machineStatus/${params.id}`,
    method: "put",
    data: params,
  });
}
// 删除
export function deleteMachineStatusManagerCol(params) {
  return request({
    url: `/machineStatus/${params.id}`,
    method: "delete",
  });
}

export function getMachineWorkerTypeManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/machineWorkerType/machineWorkerTypes?limit=${params.limit}&start=${params.start}`;
  } else {
    url = `/machineWorkerType/${params.id}`;
  }
  console.log(url);
  return request({
    url: url,
    method: "get",
  });
}
// 添加
export function addMachineWorkerTypeManagerCol(params) {
  return request({
    url: `/machineWorkerType`,
    method: "post",
    data: params,
  });
}
// 更新
export function updateMachineWorkerTypeManagerCol(params) {
  return request({
    url: `/machineWorkerType/${params.id}`,
    method: "put",
    params,
  });
}
// 删除
export function deleteMachineWorkerTypeManagerCol(params) {
  return request({
    url: `/machineWorkerType/${params.id}`,
    method: "delete",
  });
}

export function getTaskTypeManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/taskType/taskTypes?limit=${params.limit}&start=${params.start}`;
  } else {
    url = `/taskType/${params.id}`;
  }
  return request({
    url: url,
    method: "get",
  });
}
// 添加
export function addTaskTypeManagerCol(params) {
  return request({
    url: `/taskType`,
    method: "post",
    data: params,
  });
}
// 更新
export function updateTaskTypeManagerCol(params) {
  return request({
    url: `/taskType/${params.id}`,
    method: "put",
    data: params,
  });
}
// 删除
export function deleteTaskTypeManagerCol(params) {
  return request({
    url: `/taskType/${params.id}`,
    method: "delete",
  });
}

// 查询
export function getMachineWorkerManagerCols(params) {
  let url;
  if (params.id == "" || params.id == void 0) {
    url = `/machineWorker/machineWorkers`;
  } else {
    url = `/machineWorker/${params.id}`;
  }
  console.log(url);
  return request({
    url: url,
    method: "get",
    params
  });
  
}
// 添加
export function addMachineWorkerManagerCol(params) {
  return request({
    url: `/machineWorker`,
    method: "post",
    params,
  });
}
// 更新
export function updateMachineWorkerManagerCol(params) {
  return request({
    url: `/machineWorker/${params.id}`,
    method: "put",
    params,
  });
}
// 删除
export function deleteMachineWorkerManagerCol(params) {
  return request({
    url: `/machineWorker/${params.id}`,
    method: "delete",
  });
}



export function getTaskStatusManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/taskStatus/taskStatuss`
    } else {
        url = `/taskStatus/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
        params
    })
}
// 添加
export function addTaskStatusManagerCol(params) {
    return request({
        url: `/taskStatus`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateTaskStatusManagerCol(params) {
    return request({
        url: `/taskStatus/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteTaskStatusManagerCol(params) {
    return request({
        url: `/taskStatus/${params.id}`,
        method: 'delete'
    })
}

export function getProcedureStatusManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/procedureStatus/procedureStatuss?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/procedureStatus/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
    })
}
// 添加
export function addProcedureStatusManagerCol(params) {
    return request({
        url: `/procedureStatus`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateProcedureStatusManagerCol(params) {
    return request({
        url: `/procedureStatus/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteProcedureStatusManagerCol(params) {
    return request({
        url: `/procedureStatus/${params.id}`,
        method: 'delete'
    })
}



export function getInventoryManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/inventory/inventorys?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/inventory/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
    })
}
// 添加
export function addInventoryManagerCol(params) {
    return request({
        url: `/inventory`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateInventoryManagerCol(params) {
    return request({
        url: `/inventory/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteInventoryManagerCol(params) {
    return request({
        url: `/inventory/${params.id}`,
        method: 'delete'
    })
}


//为车间添加工人
export function addWorkerToWorkshop(params) {
    return request({
        url: `/workshop/worker`,
        method: 'post',
        data:params
    })
}

export function deleteWorkerToWorkshop(params) {
    return request({
        url: `/workshop/worker`,
        method: 'delete',
        params
    })
}

//为机器添加工人
export function addWorkerToDevice(params) {
    return request({
        url: `/machine/worker`,
        method: 'post',
        data:params
    })
}
export function deleteWorkerToDevice(params) {
    return request({
        url: `/machine/worker`,
        method: 'delete',
        data:params
    })
}



export function getWorkerTypeManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/workerType/workerTypes?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/workerType/${params.id}`
    }
    console.log(url)
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addWorkerTypeManagerCol(params) {
    return request({
        url: `/workerType`,
        method: 'post',
        params
    })
}
// 更新
export function updateWorkerTypeManagerCol(params) {
    return request({
        url: `/workerType/${params.id}`,
        method: 'put',
        params
    })
}
// 删除
export function deleteWorkerTypeManagerCol(params) {
    return request({
        url: `/workerType/${params.id}`,
        method: 'delete'
    })
}



// 查询
export function getProcedureItemManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/procedureItem/procedureItems`
    } else {
        url = `/procedureItem/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
        params
    })
}
// 添加
export function addProcedureItemManagerCol(params) {
    return request({
        url: `/procedureItem`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateProcedureItemManagerCol(params) {
    return request({
        url: `/procedureItem/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteProcedureItemManagerCol(params) {
    return request({
        url: `/procedureItem/${params.id}`,
        method: 'delete'
    })
}



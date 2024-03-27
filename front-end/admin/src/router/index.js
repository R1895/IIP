import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

/* Layout */
import Layout from "@/layout";

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: "/login",
    component: () => import("@/views/login/index"),
    hidden: true,
  },
  {
    path: "/",
    component: Layout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: () => import("@/views/dashboard/index"),
        meta: { title: "Dashboard", icon: "dashboard" },
      },
    ],
  },
  {
    path: "/users",
    component: Layout,
    name: "User",
    meta: { title: "人员管理", icon: "el-icon-user" },
    children: [
      {
        path: "userManager",
        name: "UserManager",
        component: () => import("@/views/user/userManager"),
        meta: { title: "用户管理", icon: "el-icon-user" },
      },
      {
        path: "roleManager",
        name: "RoleManager",
        component: () => import("@/views/user/roleManager"),
        meta: { title: "角色管理", icon: "el-icon-user" },
      },
      {
        path: "permission",
        name: "PermissionManager",
        component: () => import("@/views/user/permissionManager"),
        meta: { title: "权限管理", icon: "el-icon-user" },
      },
    ],
  },
  {
    path: "/taskMan",
    component: Layout,
    hidden: true,
    children: [
      {
        path: "addProcedureToTask",
        name: "AddProcedureToTask",
        component: () => import("@/views/industry/addProcedureToTask"),
        meta: { title: "为任务添加产线", icon: "dashboard" },
      },
      {
        path: "updateTask",
        name: "UpdateTask",
        component: () => import("@/views/industry/updateTask"),
        meta: { title: "修改任务", icon: "dashboard" },
      },
      {
        path: "workerMachineDetail",
        name: "WorkerMachineDetail",
        component: () => import("@/views/industry/workerMachineDetail"),
        meta: { title: "查看工人工作的机器详情", icon: "dashboard" },
      },
      {
        path: "workerTaskDetail",
        name: "WorkerTaskDetail",
        component: () => import("@/views/industry/workerTaskDetail"),
        meta: { title: "查看工人任务详情", icon: "dashboard" },
      },
      {
        path: "allocateWorkerToMachine",
        name: "AllocateWorkerToMachine",
        component: () => import("@/views/industry/allocateWorkerToMachine"),
        meta: { title: "为机器分配工人", icon: "dashboard" },
      },
      {
        path: "addWorkerToWorkshop",
        name: "AddWorkerToWorkshop",
        component: () => import("@/views/industry/addWorkerToWorkshop"),
        meta: { title: "为车间分配工人", icon: "dashboard" },
      },
    ],
  },
  {
    path: "/industry",
    component: Layout,
    name: "Industry",
    meta: { title: "产线管理", icon: "el-icon-s-marketing" },
    children: [
      {
        path: "workshop",
        name: "Workshop",
        component: () => import("@/views/industry/workshop"),
        meta: { title: "车间管理", icon: "el-icon-s-ticket" },
      },
      {
        path: "task",
        name: "Task",
        component: () => import("@/views/industry/task"),
        meta: { title: "产线管理", icon: "el-icon-s-claim" },
      },
      {
        path: "device",
        name: "Device",
        component: () => import("@/views/industry/device"),
        meta: { title: "设备管理", icon: "el-icon-s-management" },
      },
      {
        path: "worker",
        name: "Worker",
        component: () => import("@/views/industry/worker"),
        meta: { title: "工人管理", icon: "el-icon-s-custom" },
      },
      {
        path: "procedure",
        name: "Procedure",
        component: () => import("@/views/industry/procedure"),
        meta: { title: "工序管理", icon: "el-icon-s-ticket" },
      },
      {
        path: "procedureItem",
        name: "ProcedureItem",
        component: () => import("@/views/industry/procedureItem"),
        meta: { title: "工序产品管理", icon: "el-icon-s-ticket" },
      }
    ],
  },
  {
    path: "/typeManager",
    component: Layout,
    name: "TypeManager",
    meta: { title: "类别管理", icon: "el-icon-s-marketing" },
    children: [
    {
        path: "taskStatus",
        name: "TaskStatus",
        component: () => import("@/views/industry/taskStatus"),
        meta: { title: "产线状态管理", icon: "el-icon-s-claim" },
        },
      {
        path: "taskType",
        name: "TaskType",
        component: () => import("@/views/industry/taskType"),
        meta: { title: "产线类型管理", icon: "el-icon-s-claim" },
      },
      {
        path: "deviceStatus",
        name: "DeviceStatus",
        component: () => import("@/views/industry/deviceStatus"),
        meta: { title: "设备状态管理", icon: "el-icon-s-management" },
      },
      {
        path: "workerType",
        name: "WorkerType",
        component: () => import("@/views/industry/workerType"),
        meta: { title: "工人类别管理", icon: "el-icon-s-custom" },
      },
      {
        path: "procedureStatus",
        name: "ProcedureStatus",
        component: () => import("@/views/industry/procedureStatus"),
        meta: { title: "工序状态管理", icon: "el-icon-s-claim" },
        },
      {
        path: "procedureType",
        name: "ProcedureType",
        component: () => import("@/views/industry/procedureType"),
        meta: { title: "工序类型管理", icon: "el-icon-s-ticket" },
      },
    ],
  },
  {
    path: "/environmentManager",
    component: Layout,
    name: "environmentManager",
    meta: { title: "车间环境信息管理", icon: "el-icon-s-marketing" },
    children: [
      {
        path: "saferate",
        name: "Saferate",
        component: () => import("@/views/environment/saferate"),
        meta: { title: "安全事故率管理", icon: "el-icon-s-management" },
      },
      {
        path: "noise",
        name: "Noise",
        component: () => import("@/views/environment/noise"),
        meta: { title: "噪音管理", icon: "el-icon-s-ticket" },
      },
      {
        path: "smell",
        name: "Smell",
        component: () => import("@/views/environment/smell"),
        meta: { title: "气味管理", icon: "el-icon-s-claim" },
      },
      {
        path: "environment",
        name: "Environment",
        component: () => import("@/views/environment/environment"),
        meta: { title: "环境管理", icon: "el-icon-s-custom" },
      },
    ],
  },
  {
    path: "/camMan",
    component: Layout,
    name: "CamMan",
    children: [
      {
        path: "camToDevice",
        name: "CamToDevice",
        component: () => import("@/views/camToDevice/index"),
        meta: { title: "设备监控", icon: "el-icon-menu" },
      },
    ],
  },
  {
    path: "/manager",
    component: Layout,
    name: "Manager",
    children: [
      {
        path: "inventory",
        name: "Inventory",
        component: () => import("@/views/industry/Inventory"),
        meta: { title: "库存管理", icon: "el-icon-menu" },
      },
    ],
  },
  {
    path: "/team",
    component: Layout,
    name: "Team",
    meta: { title: "团队管理", icon: "el-icon-s-grid" },
    children: [
      {
        path: "workHour",
        name: "WorkHour",
        component: () => import("@/views/team/workAttendence"),
        meta: { title: "签到管理", icon: "el-icon-menu" },
      },
    ],
  },
  {
    path: "/anomaly",
    component: Layout,
    children: [
      {
        path: "anomaly",
        name: "Anomaly",
        component: () => import("@/views/anomaly/index"),
        meta: { title: "设备异常管理", icon: "el-icon-s-order" },
      },
    ],
  },
  {
    path: "/dictionaryItem",
    component: Layout,
    children: [
      {
        path: "dictionaryItem",
        name: "DictionaryItem",
        component: () => import("@/views/dictionaryItem/index"),
        meta: { title: "数据字典", icon: "el-icon-document-copy" },
      },
    ],
  },
  {
    path: "/message",
    component: Layout,
    children: [
      {
        path: "message",
        name: "Message",
        component: () => import("@/views/message/index"),
        meta: { title: "消息管理", icon: "el-icon-chat-round" },
      },
    ],
  },

  {
    path: "/log",
    component: Layout,
    children: [
      {
        path: "log",
        name: "Log",
        component: () => import("@/views/log/index"),
        meta: { title: "日志管理", icon: "el-icon-tickets" },
      },
    ],
  },

  // 404 page must be placed at the end !!!
  { path: "*", redirect: "/404", hidden: true },
];

const createRouter = () =>
  new Router({
    mode: "history", // require service support
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes,
  });

const router = createRouter();

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter();
  router.matcher = newRouter.matcher; // reset router
}

export default router;

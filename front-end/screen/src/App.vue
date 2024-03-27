<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style lang="scss">
.dv-scroll-board .rows .row-item {
  font-size: 20px !important;
}

.dv-scroll-board .header {
  font-size: 18px !important;
  font-weight: 550;
  margin-bottom: 5px;
}

.dv-capsule-chart .unit-label {
  font-size: 20px !important;
}

.dv-capsule-chart .capsule-item .capsule-item-column .capsule-item-value {
  font-size: 20px !important;
}

.dv-capsule-chart .label-column .label-item {
  font-size: 20px !important;
}

.dv-capsule-chart .unit-text {
  font-size: 18px !important;
}

.dv-capsule-chart .label-column {
  font-size: 20px !important;
}

.el-select-dropdown {
  background-color: #0f1325 !important;
  font-size: 20px !important;
  font-weight: bold !important;
}

.el-input__inner {
  background-color: #0f1325 !important;
  font-size: 20px !important;
  font-weight: bold !important;
}

.el-pager li {
  background-color: #0f1325 !important;
}

.el-pagination button:disabled {
  background-color: #0f1325 !important;

}

.el-pagination .btn-next, .el-pagination .btn-prev {
  background-color: #0f1325 !important;
  color: #ecf5ff !important;
}

.el-pagination {
  color: #abcaf1 !important;
}

.el-table th.el-table__cell {
  background-color: #000c !important;
}

.el-table tr {
  background-color: #000c !important;
  font-size: 20px !important;
  color: white !important;
}

.el-table--enable-row-hover .el-table__body tr:hover > td.el-table__cell {
  background-color: #606266 !important;
}
.el-table__empty-block{
  background-color: #000c !important;
}

#app {
  width: 100vw;
  height: 100vh;
  background-color: #020308;
  overflow: hidden;
}
</style>

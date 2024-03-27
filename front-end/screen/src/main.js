import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import dataV from '@jiaminghi/data-view';
// 引入全局css
import './assets/scss/style.scss';
// 按需引入vue-awesome图标
import './assets/iconfont/iconfont.css';
import Icon from 'vue-awesome/components/Icon';
import 'vue-awesome/icons/chart-bar.js';
import 'vue-awesome/icons/chart-area.js';
import 'vue-awesome/icons/chart-pie.js';
import 'vue-awesome/icons/chart-line.js';
import 'vue-awesome/icons/align-left.js';
// 引入element-ui
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
//引入echart
//4.x 引用方式
// import echarts from 'echarts'
//5.x 引用方式为按需引用
//希望使用5.x版本的话,需要在package.json中更新版本号,并切换引用方式
import * as echarts from 'echarts'
import Highcharts from 'highcharts'
import HighchartsVue from 'highcharts-vue'
import countTo from 'vue-count-to';

Vue.prototype.$echarts = echarts
Vue.config.productionTip = false;


// 全局注册
Vue.component('countTo', countTo)
Vue.use(Highcharts)
Vue.use(HighchartsVue)
Vue.use(ElementUI);
Vue.component('icon', Icon);
Vue.use(dataV);

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount('#app');

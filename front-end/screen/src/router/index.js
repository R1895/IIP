import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'live',
    component: () => import('../views/live/live.vue')
},
    {
        path: '/index',
        name: 'index',
        component: () => import('../views/index.vue')
    },
    {
        path: '/status',
        name: 'status',
        component: () => import('../views/status/status.vue')
    },
    {
        path: '/oee',
        name: 'oee',
        component: () => import('../views/oee/oee.vue')
    },
    {
        path: '/light',
        name: 'light',
        component: () => import('../views/light/light.vue')
    },
    {
        path: '/info',
        name: 'info',
        component: () => import('../views/info/info.vue')
    },
    {
        path: '/rank',
        name: 'rank',
        component: () => import('../views/working_hours_rank/rank.vue')
    },
    {
        path: '/order',
        name: 'order',
        component: () => import('../views/working_order/order.vue')
    },
    {
        path: '/procedure',
        name: 'procedure',
        component: () => import('../views/procedure/procedure.vue')
    },
    {
        path: '/stock',
        name: 'stock',
        component: () => import('../views/stock/stock.vue')
    }
]
const router = new VueRouter({
    routes
})

export default router

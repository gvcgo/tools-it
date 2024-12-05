import { createMemoryHistory, createRouter } from 'vue-router'

const routes = [
  { 
    path: '/', 
    name: "home",
    component: ()=>import('../components/layout/Layout.vue'),
    children: [
      {
        path: '',
        name: "default",
        component: ()=>import('../views/Default.vue')
      },
      {
        path:"/clocQuery",
        name:"clocQuery",
        component: ()=>import('../views/cloc/ClocQuery.vue')
      },
      {
        path:"/clocResult",
        name:"clocResult",
        component: ()=>import('../views/cloc/ClocResult.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router

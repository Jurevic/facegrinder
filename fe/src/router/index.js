import Vue from 'vue'
import Router from 'vue-router'

const routerOptions = [
  { path: '/', component: 'Landing' },
  { path: '/signup', component: 'Signup' },
  { path: '/signin', component: 'Signin' },
  { path: '/home', component: 'Home', meta: { requiresAuth: true } },
  { path: '/channels', component: 'Channels', meta: { requiresAuth: true } },
  { path: '/faces', component: 'Faces', meta: { requiresAuth: true } },
  { path: '/processors', component: 'Processors', meta: { requiresAuth: true } }
]

const routes = routerOptions.map(route => {
  return {
    ...route,
    component: () => import(`@/components/${route.component}.vue`)
  }
})

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes
})

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const isAuthenticated = localStorage.getItem('t')
  if (requiresAuth && !isAuthenticated) {
    next('/signin')
  } else {
    next()
  }
})

export default router

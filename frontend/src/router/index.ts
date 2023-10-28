import { createRouter, createWebHistory } from '@ionic/vue-router';
import { RouteRecordRaw } from 'vue-router';

const routes: Array<RouteRecordRaw> = [
  {
    path: '',
    redirect: '/glicemia'
  },
  {
    path: '/glicemia',
    component: () => import('@/views/BloodGlucose.vue')
  },
  {
    path: '/pressao',
    component: () => import('@/views/BloodPressure.vue')
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router

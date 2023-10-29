import { createRouter, createWebHistory } from '@ionic/vue-router';
import { RouteRecordRaw } from 'vue-router';
import BloodPressureView from '@/views/BloodPressure/BloodPressureView.vue'
import BloodPressureForm from '@/views/BloodPressure/BloodPressureForm.vue'
import BloodGlucoseView from '@/views/BloodGlucose/BloodGlucoseView.vue'
import BloodGlucoseForm from '@/views/BloodGlucose/BloodGlucoseForm.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '',
    redirect: '/pressao'
  },
  {
    path: '/glicemia',
    name: 'BloodGlucoseView',
    component: BloodGlucoseView
  },
  {
    path: '/glicemia/registrar',
    name: 'BloodGlucoseForm',
    component: BloodGlucoseForm
  },
  {
    path: '/pressao',
    name: 'BloodPressureView',
    component: BloodPressureView
  },
  {
    path: '/pressao/registrar',
    name: 'BloodPressureForm',
    component: BloodPressureForm
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router

import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
// import {createInertiaApp} from '@inertiajs/vue3'
// import {resolvePageComponent} from ''
// import {ZiggyVue} from '../../vendor/tightenco/ziggy/dist/vue.m'

import App from './App.vue'
import router from './router'
import piniaPluginPersistedState from 'pinia-plugin-persistedstate'

const pinia = createPinia()
pinia.use(piniaPluginPersistedState)

const app = createApp(App)

app.use(pinia)
app.use(router)

app.mount('#app')

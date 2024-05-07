import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedState from "pinia-plugin-persistedstate"
import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
import router from '@/router'
import App from '@/App.vue'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedState)

app.use(ElementPlus)
app.use(router)
app.use(pinia)
app.mount('#app')
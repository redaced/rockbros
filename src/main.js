import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
import App from './App.vue'
import router from './router'
import axios from 'axios'


const app = createApp(App).use(ElementPlus).use(router)
app.config.globalProperties.$axios = axios
app.mount('#app')
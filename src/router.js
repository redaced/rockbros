import { createRouter, createWebHistory } from 'vue-router'
import Home from './pages/HomePage.vue'
import Login from './pages/LoginPage.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: { layout: 'main-layout' }
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: { layout: 'empty-layout' }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
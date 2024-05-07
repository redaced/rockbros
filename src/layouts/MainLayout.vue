<template>
    <div v-loading="isLoading">
        <el-header>
            <router-link to="/">Home</router-link>
            <router-link v-if="!authStore.isLoggedIn" to="/login">Login</router-link>

            <div v-else>
                <p>{{ authStore.user.username }} </p>
                <a @click="logout">Logout</a>
            </div>
        </el-header>
        <main>
            <slot />
        </main>
        <footer>
            Main hol
        </footer>
    </div>
</template>

<script>
import { ref } from 'vue'
import { useAuthStore } from '@/store/auth'


export default {
    setup() {
        const authStore = useAuthStore()


        // Use ref to create a reactive reference for the loading state
        const isLoading = ref(false)

        // Use mapActions or dispatch to dispatch an action
        const logout = async () => {
            isLoading.value = true
            await authStore.logout()
            isLoading.value = false
        }

        return {
            isLoading,
            authStore,
            logout
        }
    }
}
</script>
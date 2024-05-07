<template>
    <div class="home">
        <h1 v-if="!isLoggedIn">Welcome to the Home Page! Guest</h1>
        <h1 v-else>Welcome to the Home Page! User</h1>
    </div>
</template>


<script>
import { useAuthStore } from '@/store/auth'

export default {
    setup() {
        const authStore = useAuthStore()

        const isLoggedIn = authStore.isLoggedIn

        return {
            isLoggedIn
        }
    }
}
</script>>

<template>
    <div v-loading="isLoading">
        <el-header>
            <router-link to="/">Home</router-link>
            <router-link v-if="!isLoggedIn" to="/login">Login</router-link>
            <el-dropdown v-else trigger="click">
                <span class="el-dropdown-link">
                    User<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu>
                    <template #dropdown>
                        <!-- <el-dropdown-item @click="logout">Logout</el-dropdown-item> -->
                    </template>
                </el-dropdown-menu>
            </el-dropdown>
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
import { computed, ref } from 'vue'
import { useAuthStore } from '@/store/auth'


export default {
    setup() {
        const authStore = useAuthStore()


        // Use computed to create a reactive dependency on the Vuex state
        const isLoggedIn = computed(() => authStore.isLoggedIn)

        // Use ref to create a reactive reference for the loading state
        const isLoading = ref(false)

        // Use mapActions or dispatch to dispatch an action
        const logout = async () => {
            isLoading.value = true
            await authStore.logout()
            isLoading.value = false
        }

        return {
            isLoggedIn,
            isLoading,
            authStore,
            logout
        }
    }
}
</script>
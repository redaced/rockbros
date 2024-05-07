import { defineStore } from 'pinia'

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        isLoggedIn: false,
        user: {},
    }),
    actions: {
        login(user) {
            // Perform login logic here, then set isLoggedIn to true
            this.isLoggedIn = true
            this.user = user

        },
        logout() {
            // Perform logout logic here, then set isLoggedIn to false
            this.isLoggedIn = false
            this.user = undefined
        },
    },
    persist: {
        storage: sessionStorage,
    }
})
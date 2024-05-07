import { defineStore } from 'pinia'

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        isLoggedIn: false,
    }),
    actions: {
        login() {
            // Perform login logic here, then set isLoggedIn to true
            this.isLoggedIn = true
        },
        logout() {
            // Perform logout logic here, then set isLoggedIn to false
            this.isLoggedIn = false
        },
    },
    persist: {
        storage: sessionStorage,
    }
})
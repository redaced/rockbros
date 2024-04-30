import { createStore } from 'vuex'

export default createStore({
    state: {
        isLoggedIn: false,
    },
    mutations: {
        setLoggedIn(state, value) {
            state.isLoggedIn = value
        },
    },
    actions: {
        login({ commit }) {
            // Perform login logic here, then set isLoggedIn to true
            commit('setLoggedIn', true)
        },
        logout({ commit }) {
            // Perform logout logic here, then set isLoggedIn to false
            commit('setLoggedIn', false)
        },
    },
})
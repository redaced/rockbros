<template>
    <el-form :model="form" ref="form" rules="rules">
        <el-form-item label="Username" prop="username">
            <el-input v-model="form.username" />
        </el-form-item>
        <el-form-item label="Password" prop="password">
            <el-input type="password" v-model="form.password" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm">Login</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
import { ref } from 'vue'
import { useStore } from 'vuex'
import axios from '@/api'
import router from '@/router'

export default {
    setup() {
        const store = useStore()

        const form = ref({
            username: '',
            password: ''
        })

        const submitForm = async () => {
            let data = JSON.stringify({
                "username": form.value.username,
                "password": form.value.password
            })

            const response = await axios.post('/login', data)

            if (response.data) {
                // Use the store here
                // For example, you can commit a mutation
                store.commit('setLoggedIn', true)

                // Navigate to another route after successful login
                router.push('/')
            } else {
                console.log('error submit!!')
                return false
            }
        }

        return {
            form,
            submitForm
        }
    }
}
</script>
<template>
    <el-form :model="form" ref="form" :rules="rules">
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
import { useAuthStore } from '@/store/auth'
import axios from '@/api'
import router from '@/router'

export default {
    setup() {
        const authStore = useAuthStore()
        const form = ref({
            username: '',
            password: ''
        })

        const rules = ref({
            username: [
                { required: true, message: 'Please input your username', trigger: 'blur' }
            ],
            password: [
                { required: true, message: 'Please input your password', trigger: 'blur' }
            ]
        })

        const submitForm = async () => {
            // const valid = await form.value.validate()
            // if (!valid) return

            let data = JSON.stringify({
                "username": form.value.username,
                "password": form.value.password
            })

            const response = await axios.post('/login', data)

            if (response.data) {
                // Save the JWT to localStorage
                localStorage.setItem('jwt', response.data)

                // Use the store here
                // For example, you can commit a mutation
                authStore.login()

                // Navigate to another route after successful login
                router.push('/')
            } else {
                console.log('error submit!!')
                return false
            }
        }

        return {
            form,
            rules,
            submitForm,
            authStore
        }
    }
}
</script>
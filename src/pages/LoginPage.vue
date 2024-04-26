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
import axios from 'axios'

export default {
    data() {
        return {
            form: {
                username: '',
                password: ''
            },
            rules: {
                username: [
                    { required: true, message: 'Please input your username', trigger: 'blur' }
                ],
                password: [
                    { required: true, message: 'Please input your password', trigger: 'blur' }
                ]
            }
        }
    },
    methods: {
        submitForm() {
            this.$refs.form.validate((valid) => {
                if (valid) {
                    axios.post('http://localhost:8000/login', {
                        username: this.form.username,
                        password: this.form.password
                    })
                        .then(response => {
                            console.log(response.data)
                            // Handle the successful login here, e.g. by storing the returned token and redirecting to another page
                        })
                        .catch(error => {
                            console.error(error)
                            // Handle the failed login here, e.g. by showing an error message to the user
                        })
                } else {
                    console.log('error submit!!')
                    return false
                }
            })
        }
    }
}
</script>
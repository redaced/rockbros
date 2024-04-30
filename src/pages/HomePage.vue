<template>
    <div class="home">
        <h1 v-if="!isLoggedIn">Welcome to the Home Page! Guest</h1>
        <h1 v-else>Welcome to the Home Page! User</h1>
        <!-- Add your content here -->
    </div>
</template>


<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
    setup() {
        const store = useStore()

        const isLoggedIn = computed(() => store.state.isLoggedIn)

        return {
            isLoggedIn
        }
    }
}
</script>>

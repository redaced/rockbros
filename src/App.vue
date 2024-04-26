<template>
  <component :is="layout">
    <router-view />
  </component>
</template>

<script>
import MainLayout from './layouts/MainLayout.vue'
import EmptyLayout from './layouts/EmptyLayout.vue'

export default {
  components: {
    MainLayout,
    EmptyLayout
  },
  computed: {
    layout() {
      return this.$route.meta.layout || 'empty-layout'
    }
  }
}
</script>
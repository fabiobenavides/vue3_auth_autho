<template>
  <router-view></router-view>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import emitter from "@/plugins/mitt"

import auth from "@/modules/auth"

const router = useRouter();

emitter.on("register", evt => {
  router.push({name:"register"})
})

emitter.on("login", evt => {
  router.push({name:"login"})
})

emitter.on("logout", evt => {
  auth.logout()
})

</script>

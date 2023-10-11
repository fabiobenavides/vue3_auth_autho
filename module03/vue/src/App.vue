
<template>
  <router-view></router-view>

  <v-dialog v-model="registerVisible"
            persistent
            width="40%">
    <register @close="onRegisterClose">
    </register>
  </v-dialog>
  <v-dialog v-model="loginVisible"
            persistent
            width="40%">
    <login @close="onLoginClose">
    </login>
  </v-dialog>
</template>

<script setup>
import { ref } from 'vue'
import emitter from "@/plugins/mitt";

import login  from '@/components/login.vue'
import register  from '@/components/register.vue'

const registerVisible = ref(false)
const loginVisible = ref(false)

emitter.on("register", evt => {
  registerVisible.value = true
})

emitter.on("login", evt => {
  loginVisible.value = true
})

emitter.on("logout", evt => {

})

const onRegisterClose = () => {
  registerVisible.value = false
}

const onLoginClose = () => {
  loginVisible.value = false
}

</script>

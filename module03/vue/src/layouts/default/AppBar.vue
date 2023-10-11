<template>
  <v-app-bar flat>
    <v-app-bar-title>
      <v-icon icon="mdi-circle-slice-6"></v-icon>
      PS Blogs
    </v-app-bar-title>

    <v-spacer> </v-spacer>

    <v-btn v-if="!loggedIn"
           icon
           @click="onRegister">
      <v-icon>
        mdi-account-plus
      </v-icon>
    </v-btn>
    <v-btn v-if="!loggedIn"
           icon
           @click="onLogin">
      <v-icon>
        mdi-login
      </v-icon>
    </v-btn>
    <div v-if="loggedIn">
      {{ displayName }}
      <v-btn icon
             @click="onLogout">
        <v-icon>
          mdi-logout
        </v-icon>
      </v-btn>
    </div>
  </v-app-bar>
</template>

<script setup>

import { computed } from 'vue'
import emitter from "@/plugins/mitt";

const loggedIn = computed(() => {

})

const displayName = computed(() => {

})

const onRegister = () => {
  emitter.emit("register")
}

const onLogin = () => {
  emitter.emit("login")
}

const onLogout = () => {
  emitter.emit("logout")
}

</script>

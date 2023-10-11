<template>
  <v-app-bar flat>
    <v-app-bar-title>
      <v-icon icon="mdi-circle-slice-6"></v-icon>
      PS Blogs
    </v-app-bar-title>

    <v-spacer> </v-spacer>

    <v-btn v-if="!isAuthenticated"
           icon
           @click="onLogin">
      <v-icon>
        mdi-login
      </v-icon>
    </v-btn>
    <div v-if="isAuthenticated">
      {{ currentUser.name }}
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

import { useAuth0 } from '@auth0/auth0-vue'
const auth0 = useAuth0()

const isAuthenticated = computed(() => {
  return auth0.isAuthenticated.value
})

const currentUser = computed(() => {
  if (auth0.user)
		return auth0.user.value

	return {}
})


const onLogin = () => {
  auth0.loginWithRedirect()
}

const onLogout = () => {
	auth0.logout({logoutParams: {returnTo: window.location.origin}})
}

</script>








<template>
  <v-container>
    <v-card>
      <v-card-title>
        <div class="text-h4">
          Register
        </div>
      </v-card-title>
      <v-card-text>
        <div class="text-h6">
          Full name:
        </div>
        <v-text-field v-model="fullName"></v-text-field>
        <div class="text-h6">
          User name:
        </div>
        <v-text-field v-model="userName"></v-text-field>
        <div class="text-h6">
          Password
        </div>
        <v-text-field v-model="password" type="password"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-btn variant="outlined" @click="onCancel">
          Cancel
        </v-btn>
        <v-btn variant="outlined" @click="onSubmit">
          Register
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['close']);

const fullName = ref('')
const userName = ref('')
const password = ref('')

const onSubmit = () => {

}

const onCancel = () => {
  emit("close")
}

</script>

<template>
  <v-container>
    <v-card>
      <v-card-title>
        <div class="text-h4">
          View Post
        </div>
      </v-card-title>
      <v-card-text>
        <div class="text-h6">
          Title
        </div>
        <div>
          {{ post.title }}
        </div>
        <div class="text-h6">
          Content
        </div>
        <div>
          {{ post.content }}
        </div>
      </v-card-text>
      <v-card-actions>
        <v-btn variant="outlined" @click="cancelPost()">
          Cancel
        </v-btn>
        <v-btn variant="outlined" @click="submitPost()">
          Submit
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import _ from 'lodash';

const emit = defineEmits(['close']);

const props = defineProps({
  post: {
    type: Object,
    required: true,
  }
})

const submitPost = () => {
  emit("close")
}

const cancelPost= () => {
  emit("close")
}

</script>

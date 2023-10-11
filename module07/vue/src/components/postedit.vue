<template>
  <v-container>
    <v-card>
      <v-card-title>
        <div class="text-h4">
          Edit Post
        </div>
      </v-card-title>
      <v-card-text>
        <div class="text-h6">
          Title
        </div>
        <v-text-field v-model="post.data.title"></v-text-field>
        <div class="text-h6">
          Content
        </div>
        <v-textarea v-model="post.data.content"></v-textarea>
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
import { useRouter } from 'vue-router'
import _ from 'lodash';

import { useAuth0 } from '@auth0/auth0-vue'
const auth0 = useAuth0()

import { usePostsStore } from "@/stores/posts"
const postsStore = usePostsStore()

const router = useRouter()

const props = defineProps({
  id: {
    type: Number,
    default: 0
  }
})

const post = reactive({
  data: {}
})

const submitPost = async () => {
	const token = await auth0.getAccessTokenSilently()

  postsStore.updatePost(props.id, post.data, token)
    .then(() => {
      router.push({ name: "home" })
    })
}

const cancelPost = () => {
  router.push({ name: "home" })
}

const loadPost = () => {
  const content = _.find(postsStore.posts, { id: props.id })
  if (content) {
    post.data = content
  }
}

loadPost()

</script>







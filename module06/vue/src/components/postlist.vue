<template>
  <v-container>
    <v-card>
      <v-card-text>
        <v-table>
          <thead>
            <tr>
              <th class="text-left" width="5%">
                ID
              </th>
              <th class="text-left">
                Title
              </th>
              <th class="text-left" width="20%">
                Author
              </th>
              <th class="text-right" width="5%">
              </th>
              <th class="text-right" width="5%">
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="post in computedPosts" :key="post.id">
              <td class="text-left">
                {{ post.id }}
              </td>
              <td class="text-left">
                {{ post.title }}
              </td>
              <td class="text-left">
                {{ post.author }}
              </td>
              <td class="text-right">
                <v-btn variant="outlined" @click="viewPost(post)">
                  View
                </v-btn>
              </td>
              <td class="text-right">
                <v-btn variant="outlined" @click="editPost(post)">
                  Edit
                </v-btn>
              </td>
            </tr>
          </tbody>
        </v-table>
      </v-card-text>
      <v-card-actions>
        <v-btn variant="outlined" @click="createPost()">
          Create
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import _ from 'lodash';

import { usePostsStore } from "@/stores/posts"

const router = useRouter()

const postsStore = usePostsStore()

const computedPosts = computed(() => {
  let posts = postsStore.posts;

  posts = _.orderBy(posts, ["id", "timestamp"]);

  return posts
})

const viewPost = (post) => {
  console.log(post.id)
  router.push({ name: "view", params: { id: post.id } })
}

const editPost = (post) => {
  router.push({ name: "edit", params: { id: post.id } })
}

const createPost = () => {
  router.push({ name: "create" })
}

const refreshData = () => {
  postsStore.getPosts();
}

refreshData()
</script>

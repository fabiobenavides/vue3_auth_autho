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
                <v-btn variant="outlined"
                  @click="viewPost(post)">
                  View
                </v-btn>
              </td>
              <td class="text-right">
                <v-btn variant="outlined"
                  @click="editPost(post)">
                  Edit
                </v-btn>
              </td>
            </tr>
          </tbody>
        </v-table>
      </v-card-text>
      <v-card-actions>
        <v-btn variant="outlined"
          @click="createPost()">
          Create
        </v-btn>
      </v-card-actions>
    </v-card>
    <v-dialog v-model="editDialog.visible"
              persistent
              width="80%">
      <postedit v-if="editDialog.visible"
                :post="editDialog.post"
                @close="closeEdit">
      </postedit>
    </v-dialog>
    <v-dialog v-model="viewDialog.visible"
              persistent
              width="80%">
      <postview v-if="viewDialog.visible"
                :post="viewDialog.post"
                @close="closeView">
      </postview>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import _ from 'lodash';

import postedit from "@/components/postedit"
import postview from "@/components/postview"

import { usePostsStore } from "@/stores/posts"

const postsStore = usePostsStore()

const computedPosts = computed(() => {
  let posts = postsStore.posts;

  posts = _.orderBy(posts, ["id","timestamp"]);

  return posts
})

const viewDialog = ref({
  visible: false,
  post: null,
})

const viewPost = (post) => {
  viewDialog.value.visible = true
  viewDialog.value.post = post
}

const closeView = () => {
  viewDialog.value.visible = false
  viewDialog.value.post = null
}

const editDialog = ref({
  visible: false,
  post: null,
})

const editPost = (post) => {
  editDialog.value.visible = true
  editDialog.value.post = post
}

const closeEdit = () => {
  editDialog.value.visible = false
  editDialog.value.post = null
}

const createPost = () => {
  editDialog.value.visible = true
  editDialog.value.post = {}
}

const refreshData = () => {
  postsStore.getPosts();
}

refreshData()
</script>

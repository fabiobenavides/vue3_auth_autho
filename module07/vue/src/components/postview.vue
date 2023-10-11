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
					{{ post.data.title }}
				</div>
				<div class="text-h6">
					Content
				</div>
				<div>
					{{ post.data.content }}
				</div>
			</v-card-text>
			<v-card-actions>
				<v-btn variant="outlined" @click="cancelPost()">
					Cancel
				</v-btn>
				<v-btn v-if="isAuthenticated" variant="outlined" @click="editPost()">
					Edit
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
const postsStore = usePostsStore()

import { useAuth0 } from '@auth0/auth0-vue'
const auth0 = useAuth0()

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

const isAuthenticated = computed(() => {
	return auth0.isAuthenticated.value
})

const editPost = () => {
	router.push({ name: "edit", params: { id: post.id } })
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








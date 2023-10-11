import {defineStore} from 'pinia';
import {reactive} from 'vue';
import _ from 'lodash';
import postsAPI from '@/api/posts';

import auth from '@/modules/auth'

const posts = reactive([]);

const getPosts = () => new Promise((resolve, reject) => {
	postsAPI.get()
		.then(result => {
			if (!result) {
				return;
			}

			_.forEach(result, item => {
				const index = _.findIndex(posts, {id: item.id});
				if (index == -1) {
					posts.push(item);
				} else {
					posts.splice(index, 1, item);
				}
			});

			resolve();
		})
		.catch(error => {
			reject(error);
		});
});

const submitPost = (post) => new Promise((resolve, reject) => {
    const token = auth.getToken()

	postsAPI.submit(post, token)
		.then(result => {
			if (!result) {
				return;
			}

				const index = _.findIndex(posts, {id: result.id});
				if (index == -1) {
					posts.push(result);
				} else {
					posts.splice(index, 1, result);
				}

			resolve();
		})
		.catch(error => {
			reject(error);
		});
});

export const usePostsStore = defineStore('posts', () => ({posts, getPosts, submitPost}));

import {defineStore} from 'pinia';
import {reactive} from 'vue';
import _ from 'lodash';
import postsAPI from '@/api/posts';

export const usePostsStore = defineStore('posts', () => {
	const posts = reactive([]);

	const getPosts = () => new Promise((resolve, reject) => {
		postsAPI.findMany()
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

	return {posts, getPosts};
});

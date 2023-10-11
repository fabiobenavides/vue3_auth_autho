import axios from 'axios';

const RESOURCE_VALUE = 'posts';

const get = () => new Promise((resolve, reject) => {
	const config = {};

	axios.get(RESOURCE_VALUE, config)
		.then(result => {
			resolve(result.data);
		})
		.catch(error => {
			reject(error);
		});
});

const submit = (post, token) => new Promise((resolve, reject) => {
	const config = {
		headers: {
			Token: token,
		},
	};

	axios.post(RESOURCE_VALUE, post, config)
		.then(result => {
			resolve(result.data);
		})
		.catch(error => {
			reject(error);
		});
});

export default {get, submit};

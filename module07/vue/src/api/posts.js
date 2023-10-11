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

const create = (post, token) => new Promise((resolve, reject) => {
	const config = {
		headers: {
			Authorization: 'Bearer ' + token,
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

const update = (id, post, token) => new Promise((resolve, reject) => {
	const config = {
		headers: {
			Authorization: 'Bearer ' + token,
		},
	};

	axios.post(RESOURCE_VALUE + '/' + id, post, config)
		.then(result => {
			resolve(result.data);
		})
		.catch(error => {
			reject(error);
		});
});

export default {get, create, update};

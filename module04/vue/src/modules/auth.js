
import {ref, reactive, computed} from 'vue';
import _ from 'lodash';
import authAPI from '@/api/auth';
import {useSessionStore} from '@/stores/session';

const TOKEN_KEY = 'token';

const getToken = () => localStorage.getItem(TOKEN_KEY);

const login = (userName, password) => new Promise((resolve, reject) => {
	const sessionStore = useSessionStore();

	authAPI.authenticate(userName, password)
		.then(result => {
			localStorage.setItem(TOKEN_KEY, result.token);
			sessionStore.setFullName(result.fullName);
			sessionStore.setLoggedIn(true);
			resolve();
		})
		.catch(error => {
			reject(error);
		});
});

const refresh = () => {
	const token = getToken();
	if (!token) {
		return;
	}

	return new Promise((resolve, reject) => {
		const sessionStore = useSessionStore();

		authAPI.refresh(token)
			.then(result => {
				localStorage.setItem(TOKEN_KEY, result.token);
				sessionStore.setFullName(result.fullName);
				sessionStore.setLoggedIn(true);
				resolve();
			})
			.catch(error => {
				reject(error);
			});
	});
};

const logout = () => new Promise(resolve => {
	const sessionStore = useSessionStore();

	localStorage.removeItem(TOKEN_KEY);
	sessionStore.setFullName('');
	sessionStore.setLoggedIn(false);

	resolve();
});

const register = (fullName, userName, password) => {
	const data = {
		fullName,
		userName,
		password,
	};

	return new Promise((resolve, reject) => {
		authAPI.register(data)
			.then(result => {
				resolve();
			})
			.catch(error => {
				reject(error);
			});
	});
};

export default {login, logout, register, refresh};

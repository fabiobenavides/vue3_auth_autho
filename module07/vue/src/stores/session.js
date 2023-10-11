import {defineStore} from 'pinia';
import {ref} from 'vue';

export const useSessionStore = defineStore('session', () => {
	const fullName = ref('');
	const loggedIn = ref(false);

	const setFullName = value => {
		fullName.value = value;
	};

	const getFullName = () => fullName.value;

	const setLoggedIn = value => {
		loggedIn.value = value;
	};

	const getLoggedIn = () => loggedIn.value;

	return {fullName, loggedIn, setFullName, getFullName, setLoggedIn, getLoggedIn};
});

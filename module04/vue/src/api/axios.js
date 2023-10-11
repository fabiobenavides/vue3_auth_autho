import axios from 'axios';

export default {
	setup(data) {
		axios.defaults.baseURL = import.meta.env.VITE_API_URL;
	}
}

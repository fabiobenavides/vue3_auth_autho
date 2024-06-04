import { defineStore } from "pinia";
import { ref } from 'vue';

export const useSessionStore = defineStore('session', () => {
    const fullName = ref('');
    const loggedIn = ref(false);

    const setFullName = (value) => {
        fullName.value = value;
    }

    const getFullName = () => {
        return fullName.value;
    }

    const setLoggedIn = (value) => {
        loggedIn.value = value;
    }

    const getLoggedIn = () => {
        return loggedIn.value;
    }

    return { setFullName,
        getFullName,
        setLoggedIn,
        getLoggedIn,
        fullName,
        loggedIn
     };
})
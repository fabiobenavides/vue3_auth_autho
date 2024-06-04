import authAPI from '@/api/auth'
import { useSessionStore } from '@/stores/session'
import { reject } from 'core-js/fn/promise';

const TOKEN_KEY = 'my-key-token';

const getToken = () => {
    return localStorage.getItem(TOKEN_KEY);
}

const login = (userName, password) => {
    return new Promise((resolve, reject) => {
        const sessionStore = useSessionStore;

        authAPI.authenticate(userName, password)
            .then(result => {
                localStorage.setItem(TOKEN_KEY, result.token);
                sessionStore.setFullName(result.fulName);
                sessionStore.setLoggedIn(true);
                resolve();
            })
            .catch(error => {
                reject(error)
            })
    })
}
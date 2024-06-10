import axios from 'axios';

const authenticate = (userName, password ) => {
    const creds = btoa(userName + ':' + password);

    const config = {
        headers: {
            Authorization: 'Basic ' + creds
        }
    }

    return new Promise((resolve, reject) => {
        //console.log(`btoa ${config.headers.Authorization}`);
        axios.post('authenticate', null, config)
        .then(response => { 
            resolve(response.data);
        })
        .catch(error => {
            reject(error);
        })
    });
}

const refresh = (token) => {
    const config = { 
        headers: {
            token: token
        }
    }

    return new Promise((resolve, reject) => {
        axios.post('refresh', null, config)
        .then(response => {
            resolve(response.data);
        })
        .catch(error => {
            reject(error);
        })
    });
}

const register = (data) => {
    const config = {}

    return new Promise((resolve, reject) => {
        axios.post('register', data, config)
        .then(response => {
            resolve(response.data);
        })
        .catch(error => {
            reject(error);
        })
    });
}

export default { authenticate, refresh, register }
import { createAuth0 } from '@auth0/auth0-vue'

const auth = createAuth0({
    domain: import.meta.env.VITE_AUTH_DOMAIN,
    clientId: import.meta.env.VITE_AUTH_CLIENTID,
    authorizationParams: {
        redirect_uri: window.location.origin,
        audience: 'http://localhost:8080',
        scope: 'profile email'
    }
})

export default auth

import App from './App.vue'

import { createApp } from 'vue'

import { registerPlugins } from '@/plugins'

import axios from '@/api/axios'
axios.setup()

const app = createApp(App)

registerPlugins(app)

app.mount('#app')

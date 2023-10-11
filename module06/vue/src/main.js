import App from './App.vue'

import { createApp } from 'vue'

import { registerPlugins } from '@/plugins'
import  auth  from '@/modules/auth'

import axios from '@/api/axios'
axios.setup()

const app = createApp(App)

registerPlugins(app)

auth.refresh()

app.mount('#app')

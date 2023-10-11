// Composables
import {createRouter, createWebHistory} from 'vue-router';

import { useSessionStore } from '@/stores/session'

const routes = [
	{
		path: '/',
		component: () => import('@/layouts/default/Default.vue'),
		children: [
			{
				path: '',
				name: 'home',
				component: () => import('@/components/postlist.vue'),
			},
			{
				path: '/create',
				name: 'create',
				component: () => import('@/components/postcreate.vue'),
                meta: { requiresAuth: true },
			},
			{
				path: '/view/:id',
				name: 'view',
				component: () => import('@/components/postview.vue'),
				props: route => ({
					...route.params,
					id: Number.parseInt(route.params.id, 10) || undefined,
				}),
			},
			{
				path: '/edit/:id',
				name: 'edit',
				component: () => import('@/components/postedit.vue'),
                meta: { requiresAuth: true },
				props: route => ({
					...route.params,
					id: Number.parseInt(route.params.id, 10) || undefined,
				}),
			},
		],
	},
	{
		path: '/login',
		name: 'login',
		component: () => import('@/components/login.vue'),
	},
	{
		path: '/register',
		name: 'register',
		component: () => import('@/components/register.vue'),
	},
];

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes,
});

router.beforeEach((to, from) => {
    if (to.meta.requiresAuth) {
        const sessionStore = useSessionStore()

        if (!sessionStore.getLoggedIn()) {
            return {
                name: 'login'
            }
        }
    }
})


export default router;







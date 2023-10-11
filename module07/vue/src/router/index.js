// Composables
import {createRouter, createWebHistory} from 'vue-router';

import { createAuthGuard } from '@auth0/auth0-vue'
const authGuard = createAuthGuard()

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
                beforeEnter: authGuard,
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
                beforeEnter: authGuard,
				props: route => ({
					...route.params,
					id: Number.parseInt(route.params.id, 10) || undefined,
				}),
			},
		],
	},
];

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes,
});


export default router;







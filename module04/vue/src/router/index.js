// Composables
import {createRouter, createWebHistory} from 'vue-router';

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

export default router;

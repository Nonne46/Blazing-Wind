import { createWebHistory, createRouter } from "vue-router";
import Posts from "./components/Posts.vue";
const routes = [
    {
        path: "/",
        name: "home",
        component: Posts,
    },
    {
        path: "/posts",
        name: "posts",
        component: Posts,
    },
];
const router = createRouter({
    history: createWebHistory(),
    routes,
});
export default router;

import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "./assets/tailwind.css";
import "flowbite";
import Markdown from "vue3-markdown-it";

createApp(App).use(router).use(Markdown).mount("#app");

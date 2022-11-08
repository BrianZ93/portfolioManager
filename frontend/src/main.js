import { createApp } from 'vue';
import App from './App.vue';
import { createPinia } from 'pinia';
import money from 'v-money3';

createApp(App).use(createPinia()).use(money).mount('#app');

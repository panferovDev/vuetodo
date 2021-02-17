import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import Axios from 'axios';
import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.css';

import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Checkbox from 'primevue/checkbox';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';

Axios.defaults.baseURL = 'http://127.0.0.1:3000/'

const app = createApp(App);
app.use(PrimeVue);
app.use(ToastService);
app.component('Button', Button);
app.component('InputText', InputText);
app.component('Checkbox', Checkbox);
app.component('Toast', Toast);
app.mount('#app')

import vuetify from 'vuetify';
import axios  from 'axios'
import VueClip from 'vue-clip'
window.Vue = require('vue');


Vue.component('adminLogin', require('./page/adminLogin'));
Vue.use(vuetify);
Vue.use(VueClip);
Vue.use(axios);

const app = new Vue({
    el: '#app',
});

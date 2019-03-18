import vuetify from 'vuetify';
import axios  from 'axios'
import VueClip from 'vue-clip'
window.Vue = require('vue');


Vue.component('index', require('./page/index'));
Vue.use(vuetify);
Vue.use(VueClip);
Vue.use(axios);

const app = new Vue({
    el: '#app',
});

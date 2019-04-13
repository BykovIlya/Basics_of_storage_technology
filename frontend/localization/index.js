import Vue from 'vue'
import VueI18n from 'vue-i18n';
Vue.use(VueI18n);

import VueCookie from 'vue-cookie'
Vue.use(VueCookie);

const messages = {
    ru: {
        message: {

            directors: {
                title: "Директоры",
            },

            movies: {
                title: "Фильмы",
            },

            studios: {
                title: "Студии",
            },

            boxoffice: {
                title: "Касса",
            }

        }
    }
};

Vue.config.lang = VueCookie.get('locale') || 'ru';

const i18n = new VueI18n({
    locale:  Vue.config.lang, // set locale
    messages, // set locale messages
});

export default i18n;

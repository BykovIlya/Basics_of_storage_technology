import Vue from 'vue'
import VueI18n from 'vue-i18n';
Vue.use(VueI18n);

import VueCookie from 'vue-cookie'
Vue.use(VueCookie);

const messages = {
    ru: {
        message: {
            common:{
                emptyTable:"В таблице нет записей",
            },
            actions: {
                remove:"Удалить",
                cancel:"Отмена",
                download:"Скачать",
                upload:"Загрузить",
                export:"Выгрузить",
                corrected:"Исправлено",
                ok:"Да",
            },
            modal:{
                  removeTitle: "Удаление",
                  removeDesc: "Действительно удалить",
                  addTitle: "Добавить",
                  editTitle: "Изменить",
            },
            directors: {
                title: "Директоры",
                add: "Добавить директора",
                gender: {
                    man: "муж.",
                    woman: "жен.",
                },
                table: {
                    name: "ФИО",
                    age: "Возраст",
                    gender: "Пол",
                    email: "Почта",
                    actions: "Действия"
                },
                modal: {
                    name: "ФИО:",
                    age: "Возраст:",
                    gender: "Пол:",
                    email: "Почта:",
                    placeholder: {
                        name: "введите ФИО",
                        age: "введите возраст",
                        email: "введите электронную почту",
                    }
                },
            },

            movies: {
                title: "Фильмы",
                add: "Добавить фильм",
                table: {
                    title: "Название",
                    director: "Директор",
                    year: "Год",
                    length: "Длительность",
                    studio: "Студия",
                    actions: "Действия"
                },
                modal: {
                    title: "Название: ",
                    director: "Директор: ",
                    year: "Год: ",
                    length: "Длительность: ",
                    studio: "Студия: ",
                    placeholder: {
                        title: "введите название фильма",
                        director: "введите ФИО директора ",
                        year: "введите год выпуска ",
                        length: "введите длительность в мин. ",
                        studio: "введите название студии",
                    }
                }
            },

            studios: {
                title: "Студии",
                add: "Добавить студию",
                table: {
                    name: "Название",
                    year: "Год основания",
                    all_films: "Общее количество фильмов",
                    actions: "Действия"
                },
                modal: {
                    name: "Название:",
                    year: "Год основания:",
                    all_films: "Общее количество фильмов:",
                    placeholder: {
                        name: "введи название студии",
                        year: "введи год основания",
                        all_films: "введи общее количество фильмов",
                    }
                }
            },

            boxoffice: {
                title: "Касса",
                add: "Добавить инфо о сборах",
                table: {
                    movie: "Фильм",
                    domestic_sales: "Внутренние сборы",
                    international_sales: "Международные сборы",
                    actions: "Действия"
                },
                modal: {
                    movie: "Фильм: ",
                    domestic_sales: "Внутренние сборы:",
                    international_sales: "Международные сборы:",
                    placeholder: {
                        movie: "введите название фильма",
                        domestic_sales: "введите внутренние сборы",
                        international_sales: "введите международные сборы",
                    },
                }
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

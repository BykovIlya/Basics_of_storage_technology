<template>
    <div>
        <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
            <b-navbar-brand>{{ $t("message.analytics.title") }}</b-navbar-brand>
        </b-navbar>
        <div>
            <b-card no-body>
                <b-tabs pills card vertical type="dark">
                    <b-tab title="Запрос 1" active><b-card-text>Запрос 1. Вывести год выпуска и студию фильмов, чья общая выручка превосходит указанное значение</b-card-text>

                        <b-card>
                            <b-form inline
                                    id="fieldset-horizontal"
                                    label-cols-sm="4"
                                    label-cols-lg="3"
                                    label-for="input-horizontal"
                                    @submit.stop.prevent="handleSubmit('req1',req1)"
                            >
                                <label>Введите общую выручку :   </label>
                                <b-form-input type="text"
                                              placeholder="например 1000000000"
                                              v-model="req1"></b-form-input>
                                <b-button variant="dark" class="m-1" @click="handleSubmit('req1',req1)">Результат</b-button>

                            </b-form>
                            <b-table id="req1"
                                     striped
                                     show-empty
                                     :items="itemsReq1"
                                     :fields="fieldsReq1"
                                     :current-page="currentPageReq1"
                                     :per-page="perPageReq1"
                                     :total-rows="totalRowsReq1"
                                     :busy.sync="isBusyReq1"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 2"><b-card-text>Запрос 2. Показать директоров, у которых более одного фильма в таблице Фильмы</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit2('req2')">Результат</b-button>
                            <b-table id="req2"
                                     striped
                                     show-empty
                                     :items="itemsReq2"
                                     :fields="fieldsReq2"
                                     :current-page="currentPageReq2"
                                     :per-page="perPageReq2"
                                     :total-rows="totalRowsReq2"
                                     :busy.sync="isBusyReq2"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 3"><b-card-text>Запрос 3. Показать общую прибыль студий от фильмов из таблицы Фильмы. Сортировать по убыванию</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit3('req3')">Результат</b-button>
                            <b-table id="req3"
                                     striped
                                     show-empty
                                     :items="itemsReq3"
                                     :fields="fieldsReq3"
                                     :current-page="currentPageReq3"
                                     :per-page="perPageReq3"
                                     :total-rows="totalRowsReq3"
                                     :busy.sync="isBusyReq3"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 4"><b-card-text>Запрос 4. Показать директоров и их возраст, работающих со студией Universal Studios. Сортировать по старшинству</b-card-text>
                        <b-button variant="dark" class="m-1" @click="handleSubmit4('req4')">Результат</b-button>
                        <b-table id="req4"
                                 striped
                                 show-empty
                                 :items="itemsReq4"
                                 :fields="fieldsReq4"
                                 :current-page="currentPageReq4"
                                 :per-page="perPageReq4"
                                 :total-rows="totalRowsReq4"
                                 :busy.sync="isBusyReq4"
                                 ref="table"
                        >
                        </b-table>
                    </b-tab>
                    <b-tab title="Запрос 5"><b-card-text>Запрос 5. Показать директоров-мужчин, снявших менее 2х фильмов и пользующихся почтой gmail</b-card-text>
                        <b-button variant="dark" class="m-1" @click="handleSubmit5('req5')">Результат</b-button>
                        <b-table id="req5"
                                 striped
                                 show-empty
                                 :items="itemsReq5"
                                 :fields="fieldsReq5"
                                 :current-page="currentPageReq5"
                                 :per-page="perPageReq5"
                                 :total-rows="totalRowsReq5"
                                 :busy.sync="isBusyReq5"
                                 ref="table"
                        >
                        </b-table>
                    </b-tab>
                    <b-tab title="Запрос 6"><b-card-text>Запрос 6. Показать самую молодую студию, снявшую более одного фильма из таблицы Фильмы</b-card-text>
                        <b-button variant="dark" class="m-1" @click="handleSubmit6('req6')">Результат</b-button>
                        <b-table id="req6"
                                 striped
                                 show-empty
                                 :items="itemsReq6"
                                 :fields="fieldsReq6"
                                 :current-page="currentPageReq6"
                                 :per-page="perPageReq6"
                                 :total-rows="totalRowsReq6"
                                 :busy.sync="isBusyReq6"
                                 ref="table"
                        >
                        </b-table>
                    </b-tab>
                    <b-tab title="Запрос 7"><b-card-text>Запрос 7. Показать процент фильмов студий, показанных в таблице, от общего числа фильмов студий для каждой студии</b-card-text>
                        <b-button variant="dark" class="m-1" @click="handleSubmit7('req7')">Результат</b-button>
                        <b-table id="req7"
                                 striped
                                 show-empty
                                 :items="itemsReq7"
                                 :fields="fieldsReq7"
                                 :current-page="currentPageReq7"
                                 :per-page="perPageReq7"
                                 :total-rows="totalRowsReq7"
                                 :busy.sync="isBusyReq7"
                                 ref="table"
                        >
                        </b-table>
                    </b-tab>
                    <b-tab title="Запрос 8"><b-card-text>Запрос 8. Показать общую длительность фильмов, снятых студиями, открытыми после указанного года</b-card-text>
                        <b-card>
                            <b-form inline
                                    id="fieldset-horizontal2"
                                    label-cols-sm="4"
                                    label-cols-lg="3"
                                    label-for="input-horizontal2"
                                    @submit.stop.prevent="handleSubmit8('req8',req8)"
                            >
                                <label>Введите год :   </label>
                                <b-form-input type="text"
                                              placeholder="20 век"
                                              v-model="req8"></b-form-input>
                                <b-button variant="dark" class="m-1" @click="handleSubmit8('req8',req8)">Результат</b-button>

                            </b-form>
                            <b-table id="req8"
                                     striped
                                     show-empty
                                     :items="itemsReq8"
                                     :fields="fieldsReq8"
                                     :current-page="currentPageReq8"
                                     :per-page="perPageReq8"
                                     :total-rows="totalRowsReq8"
                                     :busy.sync="isBusyReq8"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 9"><b-card-text>Запрос 9. Показать средний год выпуска фильмов, длительность которых выше указанной</b-card-text>
                        <b-card>
                            <b-form inline
                                    id="fieldset-horizontal3"
                                    label-cols-sm="4"
                                    label-cols-lg="3"
                                    label-for="input-horizontal3"
                                    @submit.stop.prevent="handleSubmit9('req9',req9)"
                            >
                                <label>Введите длительность фильма :   </label>
                                <b-form-input type="text"
                                              placeholder=" от 90 мин"
                                              v-model="req9"></b-form-input>
                                <b-button variant="dark" class="m-1" @click="handleSubmit9('req9',req9)">Результат</b-button>

                            </b-form>
                            <b-table id="req9"
                                     striped
                                     show-empty
                                     :items="itemsReq9"
                                     :fields="fieldsReq9"
                                     :current-page="currentPageReq9"
                                     :per-page="perPageReq9"
                                     :total-rows="totalRowsReq9"
                                     :busy.sync="isBusyReq9"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 10"><b-card-text>Запрос 10. Выбрать возраст второго по старшинству директора</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit10('req10')">Результат</b-button>
                            <b-table id="req10"
                                     striped
                                     show-empty
                                     :items="itemsReq10"
                                     :fields="fieldsReq10"
                                     :current-page="currentPageReq10"
                                     :per-page="perPageReq10"
                                     :total-rows="totalRowsReq10"
                                     :busy.sync="isBusyReq10"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 11"><b-card-text>Запрос 11. Показать самый дорогой фильм Кристофера Нолана</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit11('req11')">Результат</b-button>
                            <b-table id="req11"
                                     striped
                                     show-empty
                                     :items="itemsReq11"
                                     :fields="fieldsReq11"
                                     :current-page="currentPageReq11"
                                     :per-page="perPageReq11"
                                     :total-rows="totalRowsReq11"
                                     :busy.sync="isBusyReq11"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 12"><b-card-text>Запрос 12. Показать студии, которые имеют фильмы длиной более 120 минут. Показать количество таких фильмов у каждой студии</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit12('req12')">Результат</b-button>
                            <b-table id="req12"
                                     striped
                                     show-empty
                                     :items="itemsReq12"
                                     :fields="fieldsReq12"
                                     :current-page="currentPageReq12"
                                     :per-page="perPageReq12"
                                     :total-rows="totalRowsReq12"
                                     :busy.sync="isBusyReq12"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 13"><b-card-text>Запрос 13. Посчитать процентное соотношение студий в таблице Фильмы</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit13('req13')">Результат</b-button>
                            <b-table id="req13"
                                     striped
                                     show-empty
                                     :items="itemsReq13"
                                     :fields="fieldsReq13"
                                     :current-page="currentPageReq13"
                                     :per-page="perPageReq13"
                                     :total-rows="totalRowsReq13"
                                     :busy.sync="isBusyReq13"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 14"><b-card-text>Запрос 14. Посчитать количество вхождений буквы "е" в столбце Название таблицы Студии</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit14('req14')">Результат</b-button>
                            <b-table id="req14"
                                     striped
                                     show-empty
                                     :items="itemsReq14"
                                     :fields="fieldsReq14"
                                     :current-page="currentPageReq14"
                                     :per-page="perPageReq14"
                                     :total-rows="totalRowsReq14"
                                     :busy.sync="isBusyReq14"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                    <b-tab title="Запрос 15"><b-card-text>Запрос 15. Показать все фильмы с 1990 по 1999 с длительностью от 100 до 125 минут</b-card-text>
                        <b-card>
                            <b-button variant="dark" class="m-1" @click="handleSubmit15('req15')">Результат</b-button>
                            <b-table id="req15"
                                     striped
                                     show-empty
                                     :items="itemsReq15"
                                     :fields="fieldsReq15"
                                     :current-page="currentPageReq15"
                                     :per-page="perPageReq15"
                                     :total-rows="totalRowsReq15"
                                     :busy.sync="isBusyReq15"
                                     ref="table"
                            >
                            </b-table>
                        </b-card>
                    </b-tab>
                </b-tabs>
            </b-card>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Analytics",
        data () {
            return {
                formURL: 'analytics',
                background: 'dark',
                req1: '',
                currentPageReq1: 1,
                perPageReq1: 1000,
                isBusyReq1: false,
                totalRowsReq1: 0,
                itemsReq1:[],
                fieldsReq1: [
                    {
                      label: 'Фильм',
                      key: 'movie',
                    },
                    {
                        label: 'Год',
                        key: 'year',
                    },
                    {
                        label: 'Cтудия',
                        key: 'studio',
                    },
                    {
                        label: 'Общая выручка',
                        key: 'total_sales',
                    },
                ],
                req2: '',
                currentPageReq2: 1,
                perPageReq2: 1000,
                isBusyReq2: false,
                totalRowsReq2: 0,
                itemsReq2:[],
                fieldsReq2: [
                    {
                        label: 'Директор',
                        key: 'director',
                    },
                    {
                        label: 'Количество фильмов',
                        key: 'films',
                    },
                ],

                req3: '',
                currentPageReq3: 1,
                perPageReq3: 1000,
                isBusyReq3: false,
                totalRowsReq3: 0,
                itemsReq3:[],
                fieldsReq3: [
                    {
                        label: 'Студия',
                        key: 'studio',
                    },
                    {
                        label: 'Общий доход',
                        key: 'total',
                    },
                ],

                req4: '',
                currentPageReq4: 1,
                perPageReq4: 1000,
                isBusyReq4: false,
                totalRowsReq4: 0,
                itemsReq4:[],
                fieldsReq4: [
                    {
                        label: 'Директор',
                        key: 'director',
                    },
                    {
                        label: 'Возраст',
                        key: 'age',
                    },
                ],

                req5: '',
                currentPageReq5: 1,
                perPageReq5: 1000,
                isBusyReq5: false,
                totalRowsReq5: 0,
                itemsReq5:[],
                fieldsReq5: [
                    {
                        label: 'ФИО',
                        key: 'name',
                    },
                    {
                        label: 'Пол',
                        key: 'gender',
                    },
                    {
                        label: 'Почта',
                        key: 'email',
                    },
                    {
                        label: 'Фильмов снято',
                        key: 'films',
                    },
                ],

                req6: '',
                currentPageReq6: 1,
                perPageReq6: 1000,
                isBusyReq6: false,
                totalRowsReq6: 0,
                itemsReq6:[],
                fieldsReq6: [
                    {
                        label: 'Cтудия',
                        key: 'studio',
                    },
                    {
                        label: 'Год основания',
                        key: 'year',
                    },
                    {
                        label: 'Фильмов снято',
                        key: 'films',
                    },
                ],

                req7: '',
                currentPageReq7: 1,
                perPageReq7: 1000,
                isBusyReq7: false,
                totalRowsReq7: 0,
                itemsReq7:[],
                fieldsReq7: [
                    {
                        label: 'Cтудия',
                        key: 'studio',
                    },
                    {
                        label: '%',
                        key: 'percent',
                    }
                ],

                req8: '',
                currentPageReq8: 1,
                perPageReq8: 1000,
                isBusyReq8: false,
                totalRowsReq8: 0,
                itemsReq8:[],
                fieldsReq8: [
                    {
                        label: 'Cтудия',
                        key: 'studio',
                    },
                    {
                        label: 'Общая длительность фильмов',
                        key: 'time',
                    }
                ],

                req9: '',
                currentPageReq9: 1,
                perPageReq9: 1000,
                isBusyReq9: false,
                totalRowsReq9: 0,
                itemsReq9:[],
                fieldsReq9: [
                    {
                        label: 'Средний год выпуска',
                        key: 'answer',
                    },
                ],

                req10: '',
                currentPageReq10: 1,
                perPageReq10: 1000,
                isBusyReq10: false,
                totalRowsReq10: 0,
                itemsReq10:[],
                fieldsReq10: [
                    {
                        label: 'Возраст почти самого старого директора',
                        key: 'answer',
                    },
                ],

                req11: '',
                currentPageReq11: 1,
                perPageReq11: 1000,
                isBusyReq11: false,
                totalRowsReq11: 0,
                itemsReq11:[],
                fieldsReq11: [
                    {
                        label: 'Фильм',
                        key: 'movie',
                    },
                    {
                        label: 'Общие сборы',
                        key: 'total',
                    }
                ],

                req12: '',
                currentPageReq12: 1,
                perPageReq12: 1000,
                isBusyReq12: false,
                totalRowsReq12: 0,
                itemsReq12:[],
                fieldsReq12: [
                    {
                        label: 'Cтудия',
                        key: 'studio',
                    },
                    {
                        label: 'Фильмы длинее 120 минут',
                        key: 'films',
                    }
                ],

                req13: '',
                currentPageReq13: 1,
                perPageReq13: 1000,
                isBusyReq13: false,
                totalRowsReq13: 0,
                itemsReq13:[],
                fieldsReq13: [
                    {
                        label: 'Студия',
                        key: 'studio',
                    },
                    {
                        label: 'Процент фильмов',
                        key: 'percent',
                    },
                ],

                req14: '',
                currentPageReq14: 1,
                perPageReq14: 1000,
                isBusyReq14: false,
                totalRowsReq14: 0,
                itemsReq14:[],
                fieldsReq14: [
                    {
                        label: 'Букв "е" всего',
                        key: 'answer',
                    },
                ],

                req15: '',
                currentPageReq15: 1,
                perPageReq15: 1000,
                isBusyReq15: false,
                totalRowsReq15: 0,
                itemsReq15:[],
                fieldsReq15: [
                    {
                        label: 'Фильм',
                        key: 'movie',
                    },
                    {
                        label: 'Год',
                        key: 'year',
                    },
                    {
                        label: 'Длительность',
                        key: 'length',
                    },
                ],
            }
        },
        created() {
        },
        methods:{
            handleSubmit(reqnum,req) {
                let mreqnum = new String(reqnum);
                let mreq = new String(req);
                let url = this.formURL + "/" + mreqnum + "/" + mreq;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems(mreqnum,mreq)
                })
            },
            handleSubmit2(reqnum) {
                let mreqnum = new String(reqnum)
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems2(mreqnum)
                })
            },
            handleSubmit3(reqnum) {
                let mreqnum = new String(reqnum)
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems3(mreqnum)
                })
            },
            handleSubmit4(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems4(mreqnum)
                })
            },
            handleSubmit5(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems5(mreqnum)
                })
            },
            handleSubmit6(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems6(mreqnum)
                })
            },
            handleSubmit7(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems7(mreqnum)
                })
            },
            handleSubmit8(reqnum,req) {
                let mreqnum = new String(reqnum);
                let mreq = new String(req);
                let url = this.formURL + "/" + mreqnum + "/" + mreq;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems8(mreqnum,mreq)
                })
            },
            handleSubmit9(reqnum,req) {
                let mreqnum = new String(reqnum);
                let mreq = new String(req);
                let url = this.formURL + "/" + mreqnum + "/" + mreq;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems9(mreqnum,mreq)
                })
            },
            handleSubmit10(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems10(mreqnum)
                })
            },
            handleSubmit11(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems11(mreqnum)
                })
            },
            handleSubmit12(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems12(mreqnum)
                })
            },
            handleSubmit13(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems13(mreqnum)
                })
            },
            handleSubmit14(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems14(mreqnum)
                })
            },
            handleSubmit15(reqnum) {
                let mreqnum = new String(reqnum);
                let url = this.formURL + "/" + mreqnum;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.getItems15(mreqnum)
                })
            },
            getItems(reqnum,req){
                if (this.req1.length ==0){
                    this.itemsReq1 = [];
                    return;
                }
                let url = this.formURL + "/" + reqnum + "/" + req;
                this.isBusyReq1 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq1 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq1 = false;
                    this.itemsReq1 = [];
                    return []
                },error =>{
                    this.isBusyReq1 = false;
                    console.log("ERROR",error);
                });
            },
            getItems2(reqnum){
                /*if (this.req2.length == 0){
                    this.itemsReq2 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq2 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq2 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq2 = false;
                    this.itemsReq2 = [];
                    return []
                },error =>{
                    this.isBusyReq2 = false;
                    console.log("ERROR",error);
                });
            },
            getItems3(reqnum){
                /*if (this.req3.length == 0){
                    this.itemsReq3 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq3 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq3 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq3 = false;
                    this.itemsReq3 = [];
                    return []
                },error =>{
                    this.isBusyReq3 = false;
                    console.log("ERROR",error);
                });
            },
            getItems4(reqnum){
                /*if (this.req4.length == 0){
                    this.itemsReq4 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq4 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq4 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq4 = false;
                    this.itemsReq4 = [];
                    return []
                },error =>{
                    this.isBusyReq4 = false;
                    console.log("ERROR",error);
                });
            },
            getItems5(reqnum){
                /*if (this.req5.length == 0){
                    this.itemsReq5 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq5 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq5 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq5 = false;
                    this.itemsReq5 = [];
                    return []
                },error =>{
                    this.isBusyReq5 = false;
                    console.log("ERROR",error);
                });
            },
            getItems6(reqnum){
                /*if (this.req6.length == 0){
                    this.itemsReq6 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq6 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq6 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq6 = false;
                    this.itemsReq6 = [];
                    return []
                },error =>{
                    this.isBusyReq6 = false;
                    console.log("ERROR",error);
                });
            },
            getItems7(reqnum){
                /*if (this.req7.length == 0){
                    this.itemsReq7 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq7 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq7 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq7 = false;
                    this.itemsReq7 = [];
                    return []
                },error =>{
                    this.isBusyReq7 = false;
                    console.log("ERROR",error);
                });
            },

            getItems8(reqnum, req){
                if (this.req8.length == 0){
                    this.itemsReq8 = [];
                    return;
                }
                let url = this.formURL + "/" + reqnum + "/" + req;
                this.isBusyReq8 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq8 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq8 = false;
                    this.itemsReq8 = [];
                    return []
                },error =>{
                    this.isBusyReq8 = false;
                    console.log("ERROR",error);
                });
            },

            getItems9(reqnum, req){
                if (this.req9.length == 0){
                    this.itemsReq9 = [];
                    return;
                }
                let url = this.formURL + "/" + reqnum + "/" + req;
                this.isBusyReq9 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq9 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq9 = false;
                    this.itemsReq9 = [];
                    return []
                },error =>{
                    this.isBusyReq9 = false;
                    console.log("ERROR",error);
                });
            },

            getItems10(reqnum){
                /*if (this.req10.length == 0){
                    this.itemsReq10 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq10 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq10 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq10 = false;
                    this.itemsReq10 = [];
                    return []
                },error =>{
                    this.isBusyReq10 = false;
                    console.log("ERROR",error);
                });
            },

            getItems11(reqnum){
                /*if (this.req11.length == 0){
                    this.itemsReq11 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq11 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq11 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq11 = false;
                    this.itemsReq11 = [];
                    return []
                },error =>{
                    this.isBusyReq11 = false;
                    console.log("ERROR",error);
                });
            },

            getItems12(reqnum){
                /*if (this.req12.length == 0){
                    this.itemsReq12 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq12 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq12 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq12 = false;
                    this.itemsReq12 = [];
                    return []
                },error =>{
                    this.isBusyReq12 = false;
                    console.log("ERROR",error);
                });
            },

            getItems13(reqnum){
                /*if (this.req13.length == 0){
                    this.itemsReq13 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq13 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq13 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq13 = false;
                    this.itemsReq13 = [];
                    return []
                },error =>{
                    this.isBusyReq13 = false;
                    console.log("ERROR",error);
                });
            },

            getItems14(reqnum){
                /*if (this.req14.length == 0){
                    this.itemsReq14 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq14 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq14 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq14 = false;
                    this.itemsReq14 = [];
                    return []
                },error =>{
                    this.isBusyReq10 = false;
                    console.log("ERROR",error);
                });
            },

            getItems15(reqnum){
                /*if (this.req15.length == 0){
                    this.itemsReq15 = [];
                    return;
                }*/
                let url = this.formURL + "/" + reqnum;
                this.isBusyReq15 = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    if (result.status === 200 || result.status === 304 ){
                        if(result.body.length > 0) {
                            this.itemsReq15 = result.body;
                            return result.body
                        }
                    }
                    this.isBusyReq15 = false;
                    this.itemsReq15 = [];
                    return []
                },error =>{
                    this.isBusyReq15 = false;
                    console.log("ERROR",error);
                });
            },

            delete: function (url, data, callback) {
                return this.$http.delete(url,data,null).then(result => {
                    callback(result.body);
                },error =>{
                    console.log("ERROR",error);
                    if (error.status === 422){
                        callback(error.body);
                    }
                    return;
                });
            },
            post: function (url, data, callback) {
                console.log(data);
                return this.$http.post(url,data,null).then(result => {
                    callback(result);
                },error =>{
                    callback(error);
                    return;
                });
            },
            put: function (url, data, callback) {
                return this.$http.put(url,data,null).then(result => {
                    callback(result);
                },error =>{
                    callback(error);
                    return;
                });
            },
        },
    }
</script>

<style scoped>

</style>
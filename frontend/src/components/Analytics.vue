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
                    <b-tab title="Запрос 3"><b-card-text>Запрос 3</b-card-text></b-tab>
                    <b-tab title="Запрос 4"><b-card-text>Запрос 4</b-card-text></b-tab>
                    <b-tab title="Запрос 5"><b-card-text>Запрос 5</b-card-text></b-tab>
                    <b-tab title="Запрос 6"><b-card-text>Запрос 6</b-card-text></b-tab>
                    <b-tab title="Запрос 7"><b-card-text>Запрос 7</b-card-text></b-tab>
                    <b-tab title="Запрос 8"><b-card-text>Запрос 8</b-card-text></b-tab>
                    <b-tab title="Запрос 9"><b-card-text>Запрос 9</b-card-text></b-tab>
                    <b-tab title="Запрос 10"><b-card-text>Запрос 10</b-card-text></b-tab>
                    <b-tab title="Запрос 11"><b-card-text>Запрос 11</b-card-text></b-tab>
                    <b-tab title="Запрос 12"><b-card-text>Запрос 12</b-card-text></b-tab>
                    <b-tab title="Запрос 13"><b-card-text>Запрос 13</b-card-text></b-tab>
                    <b-tab title="Запрос 14"><b-card-text>Запрос 14</b-card-text></b-tab>
                    <b-tab title="Запрос 15"><b-card-text>Запрос 15</b-card-text></b-tab>
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
            }
        },
        created() {
        },
        methods:{
            handleSubmit(reqnum,req) {
                let mreqnum = new String(reqnum)
                let mreq = new String(req)
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
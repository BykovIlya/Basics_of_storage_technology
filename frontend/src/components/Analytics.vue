<template>
    <div>
        <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
            <b-navbar-brand>{{ $t("message.analytics.title") }}</b-navbar-brand>
        </b-navbar>
        <div>
            <b-card no-body>
                <b-tabs pills card vertical type="dark">
                    <b-tab title="Запрос 1" active><b-card-text>Запрос 1</b-card-text>
                        <b-btn v-b-modal.modalreq1>Введите id покупателя</b-btn>
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
                        <b-modal id="modalreq1"
                                 ref="modal"
                                 title="Enter the total sales"
                                 @ok="handleOkreq1"
                                 @shown="clearNamereq1">
                            <form @submit.stop.prevent="handleSubmitreq1">
                                <b-form-input type="text"
                                              placeholder="Enter the total sales"
                                              v-model="req1"></b-form-input>
                            </form>
                        </b-modal>
                    </b-tab>
                    <b-tab title="Запрос 2"><b-card-text>Запрос 2</b-card-text></b-tab>
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
            }
        },
        methods:{
            clearNamereq1 () {
                this.req1 = ''
            },
            handleOkreq1 (evt) {
                evt.preventDefault();
                if (!this.req1) {
                    alert('Please enter the total sales')
                } else {
                    this.handleSubmitreq1()
                }
            },
            handleSubmitreq1() {
                let req1 = new String(this.req1)
                let url = this.formURL + "/req1/" + req1;
                this.$http.get(url).then(result => {
                    console.log(result);
                    this.$refs.modal.hide();
                    this.getItems(req1)
                })
            },
            getItems(req1){
                if (this.req1.length ==0){
                    this.itemsReq1 = [];
                    return;
                }
                let url = this.formURL + "/req1/" + req1;
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
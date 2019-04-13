<template>
    <div>
        <!-- Modal Component -->
        <b-modal id="modal-center"
                 size="lg"
                 class="modal-fullscreen1"
                 ref="modal"
                 @ok="createProduct"
                 :title=title()
                 :ok-disabled="isOk"
                 centered>
            <b-form @submit.stop.prevent="handleSubmit">
                <b-form-group id="group1"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.boxoffice.modal.movie')"
                              label-for="input1">
                    <b-form-input id="input1"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="movieState"
                                  aria-describedby="inputNameFeedback"
                                  :placeholder="$t('message.boxoffice.modal.placeholder.movie')"
                                  v-model="newItem.movie"></b-form-input>
                </b-form-group>
                <b-form-invalid-feedback id="inputNameFeedback">
                    Введите не менее пяти символов
                </b-form-invalid-feedback>

                <b-form-group id="group2"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.boxoffice.modal.domestic_sales')"
                              label-for="input3-2">
                    <b-form-input id="input2"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="domestic_salesState"
                                  aria-describedby="inputPasswordFeedback"
                                  :placeholder="$t('message.boxoffice.modal.placeholder.domestic_sales')"
                                  v-model="newItem.domestic_sales"></b-form-input>
                </b-form-group>

                <b-form-invalid-feedback id="inputPasswordFeedback">
                </b-form-invalid-feedback>

                <b-form-group id="group4"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.boxoffice.modal.international_sales')"
                              label-for="input4">
                    <b-form-input id="input4"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="international_salesState"
                                  aria-describedby="inputEmailFeedback"
                                  :placeholder="$t('message.boxoffice.modal.placeholder.international_sales')"
                                  v-model="newItem.international_sales"></b-form-input>
                    <!--  <b-form-invalid-feedback id="inputLiveFeedbackEmail">
                        {{ $t("message.managers.modal.errorWithEmail") }}
                      </b-form-invalid-feedback> -->
                </b-form-group>

            </b-form>
        </b-modal>

        <!-- Modal Delete Component -->
        <b-modal id="modal-delete"
                 ref="modalDelete"
                 class="modal-fullscreen1"
                 :title="$t('message.modal.removeTitle')"
                 @ok="removeItem(deleteRow)"
                 :ok-title="$t('message.actions.remove')"
                 :cancel-title="$t('message.actions.cancel')"
                 centered>
            <p>{{ $t("message.modal.removeDesc") }} "{{deleteRow.movie}}"?</p>
        </b-modal>

    </div>
</template>

<script>
    export default {
        name: "Box",
        computed: {
            movieState() {
                if (this.isDisabled) {return null;}
                return this.newItem.movie.length > 0
            },
            domestic_salesState() {
                if (this.isDisabled) {return null;}
                return this.newItem.domestic_sales > 0
            },
            international_salesState() {
                if (this.isDisabled) {return null;}
                return this.newItem.international_sales > 0
            },
            isOk() {
                if (this.isDisabled) {return true;}
                if (!this.isDisabled_2) {
                    return !(this.movieState && this.domestic_salesState &&
                        this.international_salesState);}
                return !(this.movieState && this.domestic_salesState &&
                    this.international_salesState);
            },
        },
        data() {
            return {
                formUrl: 'boxoffice',
                isDisabled:true,
                isDisabled_2: true,
                isDisabled_ok: true,
                newItem: {
                    id: null,
                    movie: '',
                    domestic_sales: 0,
                    international_sales: 0,
                },
                newMan: {
                    movie: '',
                    domestic_sales: 0,
                    international_sales: 0,
                },
                deleteRow: {
                    id: null,
                    movie: '',
                    domestic_sales: 0,
                    international_sales: 0,
                },
            }
        },
        created() {
        },
        methods: {
            showModal(row,isView, isNew) {
                if (row.item) {
                    this.setNewItem(row, isNew);
                } else {
                    this.clearNewItem(isNew)
                }
                this.$refs.modal.show();
                this.isDisabled = isView;
                this.isDisabled_2 = isNew;
            },
            showModalDelete(row) {
                if (row.item) {
                    this.deleteRow = row.item
                }
                this.$refs.modalDelete.show()
            },
            setNewItem(row, isNew) {
                this.newItem.id = row.item.id;
                this.newItem.movie = row.item.movie;
                this.newItem.domestic_sales = parseInt(row.item.domestic_sales);
                this.newItem.international_sales = parseInt(row.item.international_sales);
            },
            clearNewItem(isNew) {
                this.newItem.id = null;
                this.newItem.movie = '';
                this.newItem.domestic_sales = 0;
                this.newItem.international_sales = 0;
            },
            hideModal() {
                if (this.$refs.fileinput){
                    this.$refs.fileinput.reset();
                }
                this.fileProducts = null;
                this.$refs.modal.hide();
            },
            title() {
                if (this.isDisabled_2) {
                    return this.$t('message.modal.editTitle')
                }
                return this.$t('message.modal.addTitle')
            },
            createProduct(evt) {
                evt.preventDefault();
                this.handleSubmit();
            },
            delete: function (url, data, callback) {
                return this.$http.delete(url, data, null).then(result => {
                    callback(result.body);
                }, error => {
                    console.log("ERROR", error);
                    if (error.status === 422) {
                        callback(error.body);
                    }
                });
            },
            removeItem(row) {
                let url = this.formUrl + '/' + row.id;
                this.delete(url, null, function (result) {
                    if (result.status === 422) {
                        //this.formErrors = result.errors
                    } else {
                        this.$emit('refresh');
                        this.hideModalDelete()
                    }
                }.bind(this));
            },
            hideModalDelete() {
                this.$refs.modalDelete.hide()
            },
            post: function (url, data, callback) {
                console.log(data);
                return this.$http.post(url, data, null).then(result => {
                    callback(result);
                }, error => {
                    callback(error);
                });
            },
            put: function (url, data, callback) {
                return this.$http.put(url, data, null).then(result => {
                    callback(result);
                }, error => {
                    callback(error);
                });
            },
            handleSubmit() {
                let self;
                let url = this.formUrl;
                if (this.newItem.id > 0) {
                    console.log(this.newItem);
                    self = this;
                    url = this.formUrl + '/' + this.newItem.id;
                    console.log(this.newItem.id);
                    this.newItem.domestic_sales = parseInt(this.newItem.domestic_sales);
                    this.newItem.international_sales = parseInt(this.newItem.international_sales);
                    self.put(url, self.newItem, function (result) {
                        console.log(result);
                        if (result.status === 200) {

                            self.newItem.id = result.body.id;

                            self.hideModal();
                            self.$emit('refresh');
                        }
                    },error => {
                        console.log("ERROR", error);
                    });

                } else {

                    this.newMan.movie = this.newItem.movie;
                    this.newMan.domestic_sales = parseInt(this.newItem.domestic_sales);
                    this.newMan.international_sales = parseInt(this.newItem.international_sales);
                    let self = this;
                    self.post(url, this.newMan, function (result) {
                        self.hideModal();
                        self.$emit('refresh');
                    },error => {
                        console.log("ERROR", error);

                    });
                }
            },

        }
    }
</script>

<style scoped>

</style>

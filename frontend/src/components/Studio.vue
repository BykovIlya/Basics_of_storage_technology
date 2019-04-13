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
                              :label="$t('message.studios.modal.name')"
                              label-for="input1">
                    <b-form-input id="input1"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="nameState"
                                  aria-describedby="inputNameFeedback"
                                  :placeholder="$t('message.studios.modal.placeholder.name')"
                                  v-model="newItem.name"></b-form-input>
                </b-form-group>
                <b-form-invalid-feedback id="inputNameFeedback">
                    Введите не менее пяти символов
                </b-form-invalid-feedback>

                <b-form-group id="group2"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.studios.modal.year')"
                              label-for="input3-2">
                    <b-form-input id="input2"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="yearState"
                                  aria-describedby="inputPasswordFeedback"
                                  :placeholder="$t('message.studios.modal.placeholder.year')"
                                  v-model="newItem.year"></b-form-input>
                </b-form-group>

                <b-form-invalid-feedback id="inputPasswordFeedback">
                </b-form-invalid-feedback>

                <b-form-group id="group4"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.studios.modal.all_films')"
                              label-for="input4">
                    <b-form-input id="input4"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="all_filmsState"
                                  aria-describedby="inputEmailFeedback"
                                  :placeholder="$t('message.studios.modal.placeholder.all_films')"
                                  v-model="newItem.all_films"></b-form-input>
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
            <p>{{ $t("message.modal.removeDesc") }} "{{deleteRow.name}}"?</p>
        </b-modal>

    </div>
</template>

<script>
    export default {
        name: "Box",
        computed: {
            nameState() {
                if (this.isDisabled) {return null;}
                return this.newItem.name.length > 0
            },
            yearState() {
                if (this.isDisabled) {return null;}
                return this.newItem.year > 1800 && this.newItem.year < 2020
            },
            all_filmsState() {
                if (this.isDisabled) {return null;}
                return this.newItem.all_films > 0
            },
            isOk() {
                if (this.isDisabled) {return true;}
                if (!this.isDisabled_2) {
                    return !(this.nameState && this.yearState &&
                        this.all_filmsState);}
                return !(this.nameState && this.yearState &&
                    this.all_filmsState);
            },
        },
        data() {
            return {
                formUrl: 'studios',
                isDisabled:true,
                isDisabled_2: true,
                isDisabled_ok: true,
                newItem: {
                    id: null,
                    name: '',
                    year: 0,
                    all_films:0,
                },
                newMan: {
                    name: '',
                    year: 0,
                    all_films: 0,
                },
                deleteRow: {
                    id: null,
                    name: '',
                    year: 0,
                    all_films: 0,
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
                this.newItem.name = row.item.name;
                this.newItem.year = parseInt(row.item.year);
                this.newItem.all_films = parseInt(row.item.all_films);
            },
            clearNewItem(isNew) {
                this.newItem.id = null;
                this.newItem.name = '';
                this.newItem.year = 0;
                this.newItem.all_films = 0;
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
                    this.newItem.year = parseInt(this.newItem.year);
                    this.newItem.all_films = parseInt(this.newItem.all_films);
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

                    this.newMan.name = this.newItem.name;
                    this.newMan.year = parseInt(this.newItem.year);
                    this.newMan.all_films = parseInt(this.newItem.all_films);
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

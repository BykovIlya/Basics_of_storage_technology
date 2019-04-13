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
                              :label="$t('message.movies.modal.title')"
                              label-for="input1">
                    <b-form-input id="input1"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="titleState"
                                  aria-describedby="inputNameFeedback"
                                  :placeholder="$t('message.movies.modal.placeholder.title')"
                                  v-model="newItem.title"></b-form-input>
                </b-form-group>
                <b-form-invalid-feedback id="inputNameFeedback">
                    Введите не менее пяти символов
                </b-form-invalid-feedback>

                <b-form-group id="group2"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.movies.modal.director')"
                              label-for="input3-2">
                    <b-form-input id="input2"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="directorState"
                                  aria-describedby="inputPasswordFeedback"
                                  :placeholder="$t('message.movies.modal.placeholder.director')"
                                  v-model="newItem.director"></b-form-input>
                </b-form-group>

                <b-form-invalid-feedback id="inputPasswordFeedback">
                </b-form-invalid-feedback>

                <b-form-group id="group4"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.movies.modal.studio')"
                              label-for="input4">
                    <b-form-input id="input4"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="studioState"
                                  aria-describedby="inputEmailFeedback"
                                  :placeholder="$t('message.movies.modal.placeholder.studio')"
                                  v-model="newItem.studio"></b-form-input>
                    <!--  <b-form-invalid-feedback id="inputLiveFeedbackEmail">
                        {{ $t("message.managers.modal.errorWithEmail") }}
                      </b-form-invalid-feedback> -->
                </b-form-group>

                <b-form-group id="group5"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.movies.modal.year')"
                              label-for="input5">
                    <b-form-input id="input5"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="yearState"
                                  aria-describedby="inputEmailFeedback"
                                  :placeholder="$t('message.movies.modal.placeholder.year')"
                                  v-model="newItem.year"></b-form-input>
                    <!--  <b-form-invalid-feedback id="inputLiveFeedbackEmail">
                        {{ $t("message.managers.modal.errorWithEmail") }}
                      </b-form-invalid-feedback> -->
                </b-form-group>

                <b-form-group id="group6"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.movies.modal.length')"
                              label-for="input6">
                    <b-form-input id="input6"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="lengthState"
                                  aria-describedby="inputEmailFeedback"
                                  :placeholder="$t('message.movies.modal.placeholder.length')"
                                  v-model="newItem.length"></b-form-input>
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
            <p>{{ $t("message.modal.removeDesc") }} "{{deleteRow.title}}"?</p>
        </b-modal>

    </div>
</template>

<script>
    export default {
        name: "Directors",
        computed: {
            titleState() {
                if (this.isDisabled) {return null;}
                return this.newItem.title.length > 0
            },
            directorState() {
                if (this.isDisabled) {return null;}
                return this.newItem.director.length > 0
            },
            studioState() {
                if (this.isDisabled) {return null;}
                return this.newItem.studio.length > 0
            },
            lengthState() {
                if (this.isDisabled) {return null;}
                return this.newItem.length > 0 && this.newItem.length < 100000
            },
            yearState() {
                if (this.isDisabled) {return null;}
                return this.newItem.year > 1900 && this.newItem.year < 2020
            },
            isOk() {
                if (this.isDisabled) {return true;}
                if (!this.isDisabled_2) {
                    return !(this.titleState && this.directorState &&
                        this.studioState && this.lengthState && this.yearState);}
                return !(this.titleState && this.directorState &&
                    this.studioState && this.lengthState && this.yearState);
            },
        },
        data() {
            return {
                formUrl: 'movies',
                isDisabled:true,
                isDisabled_2: true,
                isDisabled_ok: true,
                newItem: {
                    id: null,
                    title: '',
                    director: '',
                    year: 0,
                    length: 0,
                    studio:'',
                },
                newMan: {
                    title: '',
                    director: '',
                    year: 0,
                    length: 0,
                    studio:'',
                },
                deleteRow: {
                    id: null,
                    title: '',
                    director: '',
                    year: 0,
                    length: 0,
                    studio:'',
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
                this.newItem.title = row.item.title;
                this.newItem.director = row.item.director;
                this.newItem.studio = row.item.studio;
                this.newItem.year = parseInt(row.item.year);
                this.newItem.length = parseInt(row.item.length);
            },
            clearNewItem(isNew) {
                this.newItem.id = null;
                this.newItem.title = '';
                this.newItem.director = '';
                this.newItem.studio = '';
                this.newItem.year = 0;
                this.newItem.length = 0;
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
                    this.newItem.length = parseInt(this.newItem.length);
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

                    this.newMan.title = this.newItem.title;
                    this.newMan.director = this.newItem.director;
                    this.newMan.studio = this.newItem.studio;
                    this.newMan.year = parseInt(this.newItem.year);
                    this.newMan.length = parseInt(this.newItem.length);
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

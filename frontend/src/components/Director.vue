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
                              :label="$t('message.directors.modal.name')"
                              label-for="input1">
                    <b-form-input id="input1"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="nameState"
                                  aria-describedby="inputNameFeedback"
                                  :placeholder="$t('message.directors.modal.placeholder.name')"
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
                              :label="$t('message.directors.modal.age')"
                              label-for="input3-2">
                    <b-form-input id="input2"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="ageState"
                                  aria-describedby="inputPasswordFeedback"
                                  :placeholder="$t('message.directors.modal.placeholder.age')"
                                  v-model="newItem.age"></b-form-input>
                </b-form-group>

                <b-form-invalid-feedback id="inputPasswordFeedback">
                    введите корректный возраст
                </b-form-invalid-feedback>

                <b-form-group id="group4"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.directors.modal.email')"
                              label-for="input4">
                    <b-form-input id="input4"
                                  type="text"
                                  :disabled="isDisabled"
                                  :state="emailState"
                                  aria-describedby="inputEmailFeedback"
                                  :placeholder="$t('message.directors.modal.placeholder.email')"
                                  v-model="newItem.email"></b-form-input>
                    <!--  <b-form-invalid-feedback id="inputLiveFeedbackEmail">
                        {{ $t("message.managers.modal.errorWithEmail") }}
                      </b-form-invalid-feedback> -->
                </b-form-group>

                <b-form-group id="groupOption"
                              horizontal
                              :label-cols='2'
                              label-size='sm'
                              label='Small'
                              :label="$t('message.directors.modal.gender')"
                              label-for="inputOption">
                    <b-form-select id="inputOption"
                                   :disabled="isDisabled"
                                   v-model="newItem.gender"
                                   :required="true">
                        <template slot="first">
                            <option :value="null" disabled>-- Выберите пол --</option>
                        </template>
                        <option v-for="option in options" v-bind:value="option.value">{{option.text}}</option>
                    </b-form-select>
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
        name: "Directors",
        computed: {
            nameState() {
                if (this.isDisabled) {return null;}
                return this.newItem.name.length > 0
            },
            emailState() {
                if (this.isDisabled) {return null;}
                return this.reg.test(this.newItem.email)
            },
            ageState() {
                if (this.isDisabled) {return null;}
                return this.newItem.age > 0 && this.newItem.age < 100
            },
            isOk() {
                if (this.isDisabled) {return true;}
                if (!this.isDisabled_2) {
                    return !(this.nameState && this.emailState &&
                        this.ageState);}
                return !(this.nameState && this.emailState &&
                    this.ageState);
            },
        },
        data() {
            return {
                formUrl: 'directors',
                isDisabled:true,
                isDisabled_2: true,
                isDisabled_ok: true,
                reg: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,24}))$/,
                options : [
                    {text: this.$t('message.directors.gender.man'), value: "муж."},
                    {text: this.$t('message.directors.gender.woman'), value: "жен."},
                ],
                newItem: {
                    id: null,
                    name: '',
                    email: '',
                    age: 0,
                    gender:'',
                },
                newMan: {
                    name: '',
                    email: '',
                    age: 0,
                    gender:'',
                },
                deleteRow: {
                    id: null,
                    name: '',
                    age: 0,
                    gender:'',
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
                this.newItem.email = row.item.email;
                this.newItem.name = row.item.name;
                this.newItem.gender = row.item.gender;
                this.newItem.age = parseInt(row.item.age);
            },
            clearNewItem(isNew) {
                this.newItem.id = null;
                this.newItem.email = '';
                this.newItem.name = '';
                this.newItem.gender = '';
                this.newItem.age = 0;
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
                this.handleSubmit()
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
                    this.newItem.age = parseInt(this.newItem.age)
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

                    this.newMan.email = this.newItem.email;
                    this.newMan.name = this.newItem.name;
                    this.newMan.gender = this.newItem.gender;
                    this.newMan.age = parseInt(this.newItem.age);
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

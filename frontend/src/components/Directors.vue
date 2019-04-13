<template>
    <div>
        <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
            <b-navbar-brand>{{ $t("message.directors.title") }}</b-navbar-brand>
            <b-navbar-nav class = "ml-auto">
                <b-button variant="success" @click="showModal(false, false)">{{ $t("message.directors.add") }}</b-button>
            </b-navbar-nav>
        </b-navbar>
        <b-container id="table-container" fluid>
            <b-table id="directors"
                     ref="mytable"
                     striped
                     small
                     thead-class="fixed-top sticky-top top-fixed-thead"
                     responsive
                     show-empty
                     :busy.sync="isBusy"
                     :items="items"
                     :fields="fields"
                     :current-page="currentPage"
                     :per-page="perPage"
                     :empty-text="$t('message.common.emptyTable')"
                     :sort-by.sync="sortBy"
                     :sort-desc.sync="sortDesc"
                     :filter="filter"
                     @filtered="onFiltered"
            >
                <template slot="action" slot-scope="row">
                    <b-button size="sm" variant="success" @click.stop="showModal(row,true,true)">
                        <font-awesome-icon icon="eye" />
                    </b-button>
                    <b-btn  size="sm" variant="primary" @click="showModal(row,false,true)">
                        <font-awesome-icon icon="edit" />
                    </b-btn>
                    <b-btn  size="sm" variant="danger" @click="showModalDelete(row)">
                        <font-awesome-icon icon="trash" />
                    </b-btn>
                </template>

            </b-table>
        </b-container>

        <director ref="director" v-on:refresh="refreshTable"></director>

    </div>
</template>

<script>
    import Director from "./Director";
    export default {
        components: {
            Director,
        },
        name: "Directors",
        data() {
            return {
                background: 'dark',
                formUrl: "directors",
                sortBy: 'name',
                sortDesc: false,
                filter: null,
                items:[],
                fields: [
                    {
                        key: 'id',
                        sortable: true,
                    },
                    {
                        key: 'name',
                        label: this.$t('message.directors.table.name'),
                        sortable: true,
                    },
                    {
                        key: 'age',
                        label: this.$t('message.directors.table.age'),
                        sortable: true,
                    },
                    {
                        key: 'gender',
                        label: this.$t('message.directors.table.gender'),
                        sortable: true
                    },
                    {
                        key: 'email',
                        label: this.$t('message.directors.table.email'),
                        sortable: true
                    },
                    {
                        key: 'action',
                        label: this.$t('message.directors.table.actions'),
                        class:'actions-width'
                    },
                ],
                currentPage: 1,
                perPage: 1000,
                isBusy: false,
                totalRows: 0,
            }
        },
        created() {
            this.getItems();
        },
        methods: {
            getItems() {
                let url = this.formUrl;
                return this.$http.get(url).then(result => {
                    this.isBusy = false;
                    if (result.status === 403) {
                        this.showForbiddenAlert = true;
                    } else if (result.status === 200 || result.status === 304) {
                        this.showForbiddenAlert = false;
                        if (result.body != null) {
                            //this.totalRows = result.body.total;
                            this.items = result.body;
                            this.totalRows = this.items.length;

                            for (let i = 0; i < this.items.length; i++) {
                                var item = this.items[i];
                                //this.updateTotalItem(item)
                            }

                        } else {
                            console.log("body null");
                            this.items = [];
                        }
                    }
                }, error => {
                    this.isBusy = false;
                    console.log("ERROR", error);
                });
            },
            onFiltered(filteredItems) {
                this.totalRows = filteredItems.length;
                this.currentPage = 1
            },
            refreshTable() {
                this.getItems();
                this.$refs.mytable.refresh()
            },
            showModal(row, isRead, isNew) {
                this.$refs.director.showModal(row, isRead, isNew)
            },
            showModalDelete(row) {
                this.$refs.director.showModalDelete(row)
            },
        }
    }
</script>

<style scoped>

</style>
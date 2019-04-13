<template>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
        <b-navbar-brand>{{ $t("message.movies.title") }}</b-navbar-brand>
    </b-navbar>
</template>

<script>
    export default {
        name: "Movies",
        data() {
            return {
                background: 'dark',
            }
        }
    }
</script>

<style scoped>

</style>
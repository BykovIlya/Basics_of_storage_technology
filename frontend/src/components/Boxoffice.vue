<template>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
        <b-navbar-brand>{{ $t("message.boxoffice.title") }}</b-navbar-brand>
    </b-navbar>
</template>

<script>
    export default {
        name: "Boxoffice",
        data() {
            return {
                background: 'dark',
            }
        }
    }
</script>

<style scoped>

</style>
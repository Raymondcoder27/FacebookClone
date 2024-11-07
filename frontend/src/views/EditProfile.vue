<script setup>
    import MainNavLayout from '@/Layouts/MainNavLayout.vue';

    import DeleteUserForm from './Partials/DeleteUserForm.vue';
    import UpdatePasswordForm from './Partials/UpdatePasswordForm.vue';
    import UpdateProfileInformation from './Partials/UpdateProfileInformation.vue';

    defineProps({
        mustVerifyEmail: Boolean,
        status, String,
    })
</script>

<template>
    <MainNavLayout >
        <template #header />

    </MainNavLayout>
</template>
<script setup>
    import Close from 'vue-material-design-icons/Close.vue';

    import { useGeneralStore } from '@/stores/general';
    import { storeToRefs } from 'pinia';
    const useGeneral = useGeneralStore();
    const { isImageDisplay } = storeToRefs(useGeneral);
</script>

<template>
    
</template>
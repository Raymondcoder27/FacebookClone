<script setup>
    import Close from 'vue-material-design-icons/Close.vue';

    import { useGeneralStore } from '@/stores/general';
    import { storeToRefs } from 'pinia';
    const useGeneral = useGeneralStore();
    const { isImageDisplay } = storeToRefs(useGeneral);
</script>

<template>
    <div class="fixed z-50 top-0 left-0 w-full h-full bg-black bg-opacity-80">
        
    </div>
</template>
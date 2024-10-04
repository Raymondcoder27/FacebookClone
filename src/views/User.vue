<script setup>
    import MainNavLayout from '@/Layouts/MainNavLayout.vue';
    import CreatePostBox from '@/Layouts/CreatePostBox.vue';
    import Post from '@/components/Post.vue'

    import Camera from 'vue-material-design-icons/Camera.vue';
    import Pen from 'vue-material-design-icons/Pen.vue';

    import { useGeneralStore } from '@/stores/general';
    import { storeToRefs } from 'pinia';
    const useGeneral = useGeneralStore();
    const {isCropperModal, isImageDisplay} = storeToRefs(useGeneral);
</script>

<template>
    
</template>
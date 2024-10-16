<script setup>
    import {ref, reactive} from 'vue';
    import {useRouter} from 'vue-router';

    import VideoImage from 'vue-material-design-icons/VideoImage.vue';
    import Image from 'vue-material-design-icons/Image.vue';
    import EmoticonOutline from 'vue-material-design-icons/EmoticonOutline.vue';
    import Close from 'vue-material-design-icons/Close.vue';
    import ChevronDown from 'vue-material-design-icons/ChevronDown.vue';
    import Earth from 'vue-material-design-icons/Earth.vue';
    import DotsHorizontal from 'vue-material-design-icons/DotsHorizontal.vue';

    import {useGeneralStore} from '@/stores/general';
    import {storeToRefs} from 'pinia';
    const useGeneral = useGeneralStore();
    const {isPostOverlay} = storeToRefs(useGeneral);

    const emit = defineEmits(['showModal']);
</script>

<template>
    
</template>
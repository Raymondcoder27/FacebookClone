<script setup>
    import {ref, reactive} from 'vue';
    import {useRoute, useRouter} from 'vue-router';

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

    let ImageDisplay = ref(null);

    const form = reactive({
        text: null,
        image: null,
    })
    let error = ref(null);

    const getUploadedImage = (e) => {
        uploadedImage.value = URL.createObjectURL(file);
        form.image = e.target.files[0];
    }
</script>

<template>
    
</template>
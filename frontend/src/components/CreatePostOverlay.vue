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

    const clearImage = () => {
        ImageDisplay.value = null;
        form.image = null;
    }
</script>

<template>
    <div id="CreatePostOverly" class="fixed z-50 top-0 left-0 w-full h-full bg-white bg-opacity-70">
        <div class="grid h-screen place-items-center p-4">
            <div class="bg-white w-full mx-auto shadow-xl rounded-xl max-w-[600px]">
                <div class="flex items-center relative mx-1 my-3.5">
                    <div class="font-extrabold w-full text-center text-[22px]">
                        Create Post
                    </div>
                    <div
                    @click="isPostOverlay = false" 
                    class="absolute right-3 p-1.5 rounded-full bg-gray-200 hover:bg-gray-300 cursor-pointer">
                    <Close :size="28" fillColor="#5E6771" />
                    </div>
                </div>

                <div class="border-t border-t-gray-300">
                    <div class="p-4">
                        <div class="flex items-center">
                            <img src="https://picsum.photos/id/87/300/320" alt="" class="rounded-full ml-1 min-w-[45px] max-h-[45px]">
                            <div class="ml-4">
                                <div class="font-extrabold">Raymond Mwebe</div>
                                <div class="flex items-center justify-between bg-gray-200 px-2 rounded-lg w-[100px] p-0.5">
                                    <Earth :size="18" />
                                    <span class="font-bold pl-1.5 text-[13px]">Public</span>
                                    <ChevronDown :size="18" class="pr-10 pl-1" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
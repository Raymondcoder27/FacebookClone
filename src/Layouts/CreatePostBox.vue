<script setup>
import { toRefs } from 'vue';
import VideoImage from "vue-material-design-icons/VideoImage.vue";
import Image from "vue-material-design-icons/Image.vue"
import EmoticonOutline from "vue-material-design-icons/EmoticonOutline.vue"

import { useGeneralStore } from '@/stores/general';
import { storeToRefs } from 'pinia';
const useGeneral = useGeneralStore();
const {isPostOverlay} = storeToRefs(useGeneral)

const props = defineProps({
    image: String,
    placeholder: String,
})

const {image, placeholder} = toRefs(props)
</script>

<template>
    <div id="CreatePostBox" class="w-full bg-white rounded-lg px-3 mt-4 shadow-md">
        <div class="flex items-center py-3 border-b">
            <router-link class="mr-2">
                <img :src="image" alt="" class="rounded-ful ml-1 min-w-[36px] max-h-[36px]">
            </router-link>
            <div
            @click="isPostOverlay = true"
             class="flex items-center justify-start p-2 rounded-full w-full cursor-pointer bg-[#EFF2F5]">
                <div class="text-left pl-2">{{ placeholder }}</div>
            </div>
        </div>
    </div>
</template>
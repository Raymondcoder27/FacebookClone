<script setup>
import { toRefs } from "vue";
import VideoImage from "vue-material-design-icons/VideoImage.vue";
import Image from "vue-material-design-icons/Image.vue";
import EmoticonOutline from "vue-material-design-icons/EmoticonOutline.vue";

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";



import { defineEmits } from 'vue';
import CreatePostOverlay from "@/components/CreatePostOverlay.vue";

const emit = defineEmits(['postAdded']);

// function emitPostCreated() {
//   emit('postAdded');
//   isPostOverlay.value = false; // Close the overlay
// }

function emitPostCreated() {
  isPostOverlay.value = false; // Close the overlay
  emit("postAdded");           // Notify parent
}




const useGeneral = useGeneralStore();
const { isPostOverlay } = storeToRefs(useGeneral);

const props = defineProps({
  image: String,
  placeholder: String,
});

const { image, placeholder } = toRefs(props);
</script>

<template>
  <CreatePostOverlay v-if="isPostOverlay" @postAdded="emitPostCreated" />
  <div
    id="CreatePostBox"
    class="w-full bg-white rounded-lg px-3 mt-4 shadow-md"
  >
    <div class="flex items-center py-3 border-b">
      <router-link class="mr-2">
        <img
          :src="image"
          alt=""
          class="rounded-full ml-1 min-w-[36px] max-h-[36px]"
        />
      </router-link>
      <div
        @click="isPostOverlay = true"
        class="flex items-center justify-start p-2 rounded-full w-full cursor-pointer bg-[#EFF2F5]"
      >
        <div class="text-left pl-2">{{ placeholder }}</div>
      </div>
    </div>

    <div class="flex items-center py-3 border-b">
      <button
        class="flex items-center justify-center p-1 w-full rounded-lg hover:bg-[#F2F2F2] mx-l cursor-pointer"
      >
        <VideoImage :size="35" fillColor="#F12848" />
        <div class="font-bold text-gray-400">Live Video</div>
      </button>
      <button
        class="flex items-center justify-center p-1 w-full rounded-lg hover:bg-[#F2F2F2] mx-l cursor-pointer"
      >
        <Image :size="35" fillColor="#43BE62" />
        <div class="font-bold text-gray-400">Photo/Video</div>
      </button>
      <button
        class="flex items-center justify-center p-1 w-full rounded-lg hover:bg-[#F2F2F2] mx-l cursor-pointer"
      >
        <EmoticonOutline :size="35" fillColor="#F8B927" />
        <div class="font-bold text-gray-400">Feeling/Activity</div>
      </button>
    </div>
  </div>
</template>
<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import PrimaryButton from "@/components/PrimaryButton.vue";
import GuestLayout from "@/Layouts/GuestLayout.vue";
// import TextInput from "@"

import Camera from "vue-material-design-icons/Camera.vue";
import Pen from "vue-material-design-icons/Pen.vue";

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
const useGeneral = useGeneralStore();
const { isCropperModal, isImageDisplay } = storeToRefs(useGeneral);

defineProps({ posts: Object, user: Object });
</script>


<template>
  <Head title="User" />
  <MainNavLayout>
    <div class="w-full pb-20 min-h-[100vh] bg-[#F1F2F5]">
      <div class="w-full bg-white">
        <div class="mx-auto pb-1 pt-[56px] max-w-[1100px]">


        </div>
      </div>

      <div
        class="flex flex-col md:flex-row w-full justify-between md:px-0 max-w-[1100px] h-[calc(100%-56px)] px-2 mx-auto"
      >

      </div>
    </div>
  </MainNavLayout>
</template>




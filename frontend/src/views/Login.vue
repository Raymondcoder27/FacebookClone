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
  <Head title="Login" />
  <!-- <GuestLayout> -->
  <div class="w-full min-h-[100vh] bg-white text-center">
    <!-- <div class="w-full bg-red-300">
        <div class="mx-auto pb-1 pt-[56px] max-w-[1100px] bg-green-400"></div>
      </div> -->

    <!-- <form action="" @submit.prevent="submit">
      <div>
        <TextInput
          id="email"
          type="email"
          class="mt-1 block w-full"
          v-model="form.email"
          required
          autofocus
          autocomplete="username"
          placeholder="Email"
        />
        Hello World!!

        <InputError class="mt-2" :message="form.errors.email" />
      </div>
    </form> -->
    <div class="bg-white border border-black shadow-black shadow-md text-lg text-black justify-center flex w-fit items-center">Hello World</div>

  </div>
  <!-- </GuestLayout> -->
</template>
<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import PrimaryButton from "@/components/PrimaryButton.vue";
import GuestLayout from "@/Layouts/GuestLayout.vue";
import TextInput from "@/components/TextInput.vue";
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
  <!-- <img src="/public/icons/FacebookLogo.png" class="w-[200px]" alt=""> -->

  <GuestLayout>
  <div class="w-full min-h-[100vh] bg-white text-center justify-center items-center mb-5">
    <!-- <div class="w-full bg-red-300">
        <div class="mx-auto pb-1 pt-[56px] max-w-[1100px] bg-green-400"></div>
      </div> -->
    <!-- <img src="/public/icons/FacebookLogo.png" class="text-sm" alt=""> -->
    <form action="" @submit.prevent="submit" class="mb-5">
    <div>
        <TextInput
          id="email"
          type="email"
          class="mt-1 block w-full"
          required
          autofocus
          autocomplete="username"
          placeholder="Email"
        />
        Hello World!!

        <!-- <InputError class="mt-2" :message="form.errors.email" /> -->
      </div>




      <div class="mt-4">
            <TextInput
              id="password"
              type="password"
              class="mt-1 block w-full"
              required
              autofocs
              autocomplete="username"
              placeholder="Email"
            />
          </div>

          <div class="flex items-center justify-center pt-4">
            <PrimaryButton class="w-full">
              Login
            </PrimaryButton>
          </div>
    </form>
  </div>
  </GuestLayout>
</template>
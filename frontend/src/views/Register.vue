<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import PrimaryButton from "@/components/PrimaryButton.vue";
import GuestLayout from "@/Layouts/GuestLayout.vue";
import TextInput from "@/components/TextInput.vue";
import api from "@/config/api";
// import TextInput from "@"

import Camera from "vue-material-design-icons/Camera.vue";
import Pen from "vue-material-design-icons/Pen.vue";

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
import router from "@/router";
const useGeneral = useGeneralStore();
const { isCropperModal, isImageDisplay } = storeToRefs(useGeneral);
const canResetPassword = true;

defineProps({ posts: Object, user: Object });

const register = async () => {
  console.log("Registering...");

  try {
    await api.post("/register", {
      // name: username.value,
      email: email.value,
      password: password.value,
    });
    router.push("/login");
  } catch (error) {
    console.log(error);
    // error.value = error
  }
};
</script>

<template>
  <!-- <Head title="Register" /> -->
  <div v-if="error" class="text-red-500">{{ error }}</div>

  <GuestLayout>
    <div
      class="w-full max-h-[80vh] bg-white text-center justify-center items-center mb-5"
    >
      <!-- <img src="/public/icons/FacebookLogo.png" class="text-sm" alt=""> -->
      <form action="" @submit.prevent="submit" class="mb-5 text-black">
        <div>
          <TextInput
            id="username"
            type="name"
            class="mt-1 block w-full"
            required
            autofocus
            autocomplete="username"
            placeholder="Name"
          />
        </div>

        <div class="mt-4">
          <TextInput
            id="email"
            type="email"
            class="mt-1 block w-full"
            required
            autofocus
            autocomplete="username"
            placeholder="Email"
          />
          <!-- Hello World!! -->
          <!-- <InputError class="mt-2" :message="form.errors.email" /> -->
        </div>

        <div class="mt-4">
          <TextInput
            id="password"
            type="password"
            class="mt-1 block w-full"
            required
            autofocus
            autocomplete="password"
            placeholder="Password"
          />
        </div>

        <div class="mt-4">
          <TextInput
            id="password"
            type="password"
            class="mt-1 block w-full"
            required
            autofocus
            autocomplete="username"
            placeholder="Confirm Password"
          />
        </div>

        <div class="flex items-center justify-center pt-4">
          <PrimaryButton class="w-full text-sm" @click="register()">
            Register
          </PrimaryButton>
        </div>

        <div class="flex items-center justify-center my-3">
          <RouterLink
            v-if="canResetPassword"
            to="/login"
            class="hover:underline font-bold text-blue-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Already Registered?
          </RouterLink>
        </div>

        <!-- <div class="border-b border-b-gray-300"></div>
          <div class="flex items-center justify-center pt-4 pb-2">
            <RouterLink
              class="px-5 py-3 text-white bg-[#37A621] text-[20px] font-bold rounded-lg text-sm"
            >
              Log in.
            </RouterLink>
          </div> -->
      </form>
    </div>
  </GuestLayout>
</template>
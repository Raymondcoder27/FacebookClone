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
    <Layout>
      <div class="w-full min-h-[100vh] bg-[#F1F2F5]">
        <div class="w-full bg-white">
          <div class="mx-auto pb-1 pt-[56px] max-w-[1100px]">
  
  
          </div>
        </div>
  
        <div
          class="flex flex-col md:flex-row w-full justify-between md:px-0 max-w-[1100px] h-[calc(100%-56px)] px-2 mx-auto"
        >
  


        <GuestLayout>
    <!-- <div class="w-full min-h-[100vh] bg-[#F1F2F5] z-50 h-screen"> -->
    <div class="w-full pb-20 min-h-[100vh] bg-[#F1F2F5]">
      <div class="w-full bg-white">
        <div class="mb-4 font-medium text-sm text-green-600">
          {{ status }}
        </div>

        <form action="" @submit.prevent="submit">
          <div>
            <!-- <InputLabel for="email" value="Email"> -->
            <!-- <TextInput
                    id="email"
                    type="email"
                    class="mt-1 block w-full"
                    v-model="form.email"
                    required
                    autofocs
                    autocomplete="username"
                    placeholder="Email"/> -->
            <TextInput
              id="email"
              type="email"
              class="mt-1 block w-full"
              v-model="form.email"
              required
              autofocs
              autocomplete="username"
              placeholder="Email"
            />
            <InputError class="mt-2" :message="form.errors.email" />
          </div>

          <div class="mt-4">
            <TextInput
              id="password"
              type="password"
              class="mt-1 block w-full"
              v-model="form.password"
              required
              autofocs
              autocomplete="username"
              placeholder="Email"
            />
            <InputError class="mt-2" :message="form.errors.password" />
          </div>

          <div class="flex items-center justify-center pt-4">
            <PrimaryButton class="w-full" :disabled="form.processing">
              Login
            </PrimaryButton>
          </div>

          <div class="flex items-center justify-center mt-4">
            <!-- <RouterLink
                v-if="canReset"
                :href=""
                class="hover:underline font-bold text-blue-500 focus:outline-none focus:ring-2">
                    Forgot your password?
                </RouterLink> -->
            <!-- <RouterLink
                v-if="canReset"
                class="hover:underline font-bold text-blue-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                    Forgot your password?
                </RouterLink> -->
            <!-- <PrimaryButton class="ml-4" :class="{ 'opacity-25': form.processing }" :disabled="form.processing">
                    Login
                </PrimaryButton> -->
          </div>
          <div class="border-b border-b-gray-200"></div>
          <div class="flex items-center justify-center pt-6 pb-2">
            <RouterLink
              class="px-5 py-3 text-white bg-[#42B792] hover:bg-[#37A621] text-[20px] font-bold rounded-lg"
            >
              Create a new account.
            </RouterLink>
          </div>
          <PrimaryButton class="ml-4" :disabled="form.processing">
            Login
          </PrimaryButton>
        </form>
      </div>
    </div>
  </GuestLayout>
        </div>
      </div>
    </Layout>
  </template>
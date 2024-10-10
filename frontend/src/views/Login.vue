<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import PrimaryButton from "@/components/PrimaryButton.vue";
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
  <!-- <MainNavLayout> -->
    <div class="w-full min-h-[100vh] bg-[#F1F2F5] z-50 h-screen">

        <div class="mb-4 font-medium text-sm text-green-600">
            {{ status }}
        </div>

        <form action="" @submit.prevent="submit">
            <div>
                <!-- <InputLabel for="email" value="Email"> -->
                    <TextInput
                    id="email"
                    type="email"
                    class="mt-1 block w-full"
                    v-model="form.email"
                    required
                    autofocs
                    autocomplete="username"
                    placeholder="Email"/>
                    <InputError class="mt-2" :message="form.errors.email"/>
            </div>

            <div class="mt-4">
                <TextInput
                    id="password"
                    type="password"
                    class="mt-1 block w-full"
                    v-model="form.email"
                    required
                    autofocs
                    autocomplete="username"
                    placeholder="Email"/>
                    <InputError class="mt-2" :message="form.errors.password"/>
            </div>

            <div class="block mt-4">
                <label for="" class="flex items-center">
                    <Checkbox name="remember" v-model:checked="form.remember"/>
                    <span class="text-sm ml-2 text-gray-600">Remember me</span>
                </label>
            </div>

            <div class="flex items-center justify-center mt-4">
                <RouterLink>
                    Forgot your password?
                </RouterLink>

                <PrimaryButton class="ml-4" :class="{'opacity-25':form.processing}" :disabled="form.processing">
                    Login
                </PrimaryButton>
            </div>
        </form>
    </div>
  <!-- </MainNavLayout> -->
</template>
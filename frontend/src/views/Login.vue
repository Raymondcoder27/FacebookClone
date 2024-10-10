<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";

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
                <InputLabel for="email" value="Email">
                    <TextInput
                    id="email"
                    type="email"
                    class="mt-1 block w-full"
                    v-model="form.email"
                    required
                    autofocs
                    autocomplete="username">
                </InputLabel>
            </div>
        </form>
    </div>
  <!-- </MainNavLayout> -->
</template>
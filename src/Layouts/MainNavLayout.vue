<script setup>
import { ref } from "vue";
import Magnify from "vue-material-design-icons/Magnify.vue";
import Home from "vue-material-design-icons/Home.vue";
import HomeOutline from "vue-material-design-icons/HomeOutline.vue";
import TelevisionPlay from "vue-material-design-icons/TelevisionPlay.vue";
import StorefrontOutline from "vue-material-design-icons/StorefrontOutline.vue";
import AccountGroup from "vue-material-design-icons/AccountGroup.vue";
import ControllerClassicOutline from "vue-material-design-icons/ControllerClassicOutline.vue";
import DotsGrid from "vue-material-design-icons/DotsGrid.vue";
import FacebookMessenger from "vue-material-design-icons/FacebookMessenger.vue";
import Bell from "vue-material-design-icons/Bell.vue";
import Logout from "vue-material-design-icons/Logout.vue";
// Removed unused imports for brevity

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
const useGeneral = useGeneralStore();
const { isPostOverlay, isCropperModal, isImageDisplay } =
  storeToRefs(useGeneral);

let showMenu = ref(false);
</script>

<template>
  <div
    id="MainNav"
    class="flex w-full fixed z-50 items-center justify-between bg-white shadow-xl border-b h-[56px]"
  >
    <div id="NavLeft" class="flex items-center justify-start w-[260px]">
      <router-link to="/" class="pl-3 min-w-[55px]">FB</router-link>
      <div
        class="flex items-center justify-center p-1 rounded-full ml-2 bg-[#EFF2F5] h-[40px]"
      >
        <Magnify class="p-1" :size="22" fillColor="#64676B" />
        <input
          type="text"
          placeholder="Search Facebook"
          class="bg-[#EFF2F5] placeholder-[#64676B] lg:block hidden border-none p-0 ring-0 focus:ring-0"
        />
      </div>
    </div>

    <div
      id="NavCenter"
      class="hidden lg:flex items-center ml-5 justify-center w-8/12 max-w-[600px]"
    >
      <router-link to="/">
        <div
          :class="$route.path === '/' ? 'mt-1.5' : ''"
          class="flex items-center justify-center p-1 w-full rounded-lg h-[48px] cursor-pointer"
        >
          <div>
            <Home
              v-if="$route.path === '/'"
              class="mx-auto"
              :size="27"
              fillColor="#1A73E3"
            />
            <HomeOutline
              v-else
              class="mx-auto"
              :size="32"
              fillColor="#64676B"
            />
          </div>
        </div>

        <div
        v-if="$route.path ==='/'"
         class="border-b-4 border-b-blue-400 rounded-md">

        </div>
      </router-link>
    </div>
  </div>
  <slot />
</template>
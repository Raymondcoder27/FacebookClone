<script setup>
import { ref } from "vue";
import CreatePostOverlay from "@/components/CreatePostOverlay.vue";

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

import CropperModal from "@/components/CropperModal.vue";
import ImageDisplay from "@/components/ImageDisplay.vue";

const handleLogout = () => {

}


import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
const useGeneral = useGeneralStore();
const { isPostOverlay, isCropperModal, isImageDisplay } = storeToRefs(useGeneral);

let showMenu = ref(false);

const user = async () => {
    console.log("retrieving user...");
    error.value = null; // Reset error state before attempting login
    try {
       const response = await api.get('/validate');

       if (response.status === 200) {
        const userDetails = response.data

        console.log(userDetails)
       }

    } catch (err) {
        // Set error message if login fails
        error.value = "Login failed. Please check your credentials.";
    }
};
</script>

<template>

  <div
    id="MainNav"
    class="flex w-full fixed z-50 items-center justify-between bg-white shadow-xl border-b h-[56px]"
  >
  {{response}}

    <div id="NavLeft" class="flex items-center justify-start w-[260px]">
      <router-link to="/home" class="pl-3 min-w-[55px]">
        <img src="/public/icons/FacebookLogoCircle.png" alt="">
      </router-link>
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
      <router-link to="/home" class="w-full">
        <div
          :class="$route.path === '/' ? 'mt-1.5' : ''"
          class="flex items-center justify-center p-1 w-full rounded-lg h-[48px] cursor-pointer hover:bg-[#F2F2F2]"
        >
          <div>
            <Home
              v-if="$route.path === '/home'"
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
      <button class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]">
        <TelevisionPlay class="mx-auto" :size="27" fillColor="#64676B" />
      </button>
      <button class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]">
        <StorefrontOutline class="mx-auto" :size="27" fillColor="#64676B" />
      </button>
      <button class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]">
        <span class="rounded-full p-0.5 border-[2px] border-[#64676B]">
        <AccountGroup class="mx-auto" :size="20" fillColor="#64676B" />
        </span>
      </button>
      <button class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]">
        <ControllerClassicOutline class="mx-auto" :size="32" fillColor="#64676B" />
      </button>
    </div>

    <div class="flex items-center justify-end mr-4 w-2/12">
        <button class="rounded-full p-2 hover:bg-gray-300 mx-1 cursor-pointer bg-[#E3E6EA]">
            <DotsGrid :size="23" fillColor="#050505"/>
        </button>
        <button class="rounded-full p-2 hover:bg-gray-300 mx-1 cursor-pointer bg-[#E3E6EA]">
            <FacebookMessenger :size="23" fillColor="#050505"/>
        </button>
        <button class="rounded-full p-2 hover:bg-gray-300 mx-1 cursor-pointer bg-[#E3E6EA]">
            <Bell :size="23" fillColor="#050505"/>
        </button>
        <div class="flex items-center justify-center relative">
          <button @click="showMenu = !showMenu">
            <img src="https://picsum.photos/id/87/300/320" alt="" class="rounded-full ml-1 cursor-pointer min-w-[40px] max-h-[40px]">
          </button>
          <div
          v-if="showMenu"
           class="absolute bg-white shadow-xl top-10 rounded-lg p-1 border mt-1 w-[330px] right-0">
            <router-link to="/"
            @click="showMenu = !showMenu">
            <div class="flex items-center gap-3 rounded-lg hover:bg-gray-200 p-2">
              <img src="https://picsum.photos/id/87/300/320" alt="" class="rounded-full ml-1 cursor-pointer min-w-[35px] max-h-[35px]">
              <span>Raymond Mwebe</span>
            </div>
            </router-link>
            <button class="w-full flex items-center gap-3 hover:bg-gray-200 px-2 py-2.5 rounded-lg" as="button" @click="handleLogout" method="post">
              <Logout class="pl-2" :size="30"/>
              <span>Logout</span>
            </button>
            <div class="text-xs font-semi-bold p-2 pt-3 border-t mt-2">
              Privacy . Terms . Advertising . AdChoices . Cookies . Meta 
            </div>
          </div>
        </div>
    </div>
  </div>
  <slot />
  <CreatePostOverlay
   v-if="isPostOverlay"
  @showModal="isPostOverlay = false" />


  <CropperModal
   v-if="isCropperModal"
  @showModal="isCropperModal = false" />

  <ImageDisplay
  v-if="isImageDisplay"
  />
</template>
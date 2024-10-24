<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue"

import Magnify from "vue-material-design-icons/Magnify.vue";
import TelevisionPlay from "vue-material-design-icons/TelevisionPlay.vue";
import StorefrontOutline from "vue-material-design-icons/StorefrontOutline.vue";
import AccountGroup from "vue-material-design-icons/AccountGroup.vue";
import AccountMultiple from "vue-material-design-icons/AccountMultiple.vue";
import Flag from "vue-material-design-icons/Flag.vue";
import ClockTimeTwoOutline from "vue-material-design-icons/ClockTimeTwoOutline.vue";
import Restore from "vue-material-design-icons/Restore.vue";
import VideoImage from "vue-material-design-icons/VideoImage.vue";
import DotsHorizontal from "vue-material-design-icons/DotsHorizontal.vue";

//import router from "@/router";
import {ref, onMounted} from 'vue'
import api from "@/config/api";
import { useAuthStore } from "@/stores/auth";
const userDetails = ref(null);
const authStore = useAuthStore();
// const getUserDetails = async () => {
//   const token = authStore.token;
//   if (!token) {
//     console.error("No token found");
//     return;
//   }
  
//   try {
//     // const response = await api.get("/user", {
//     const response = await api.get("/validate", {
//       headers: {
//         Authorization: `${token}`,
//       },
//     });
//     userDetails.value = response.data;
//   } catch (error) {
//     console.error("Failed to fetch user details", error);
//   }
// };



// onMounted(() => {
//   getUserDetails();
// });

onMounted(async () => {
  await getUserDetails();
});
  const getUserDetails = async () => {
  const token = authStore.token;
  if (!token) {
    console.error("No token found");
    return;
  }
  
  try {
    // const response = await api.get("/user", {
    const response = await api.get("/validate", {
      headers: {
        Authorization: `${token}`,
      },
    });
    userDetails.value = response.data;
  } catch (error) {
    console.error("Failed to fetch user details", error);
  }
};

</script>

<template>
  <MainNavLayout> 
    <div class="fixed w-full bg-[#F1F2F5] h-[100%]">
        <div class="grid grid-rows-3 grid-flow-col w-full mt-[56px] max-w-[1600px] h-[calc(100%-56px)] mx-auto px-4">
            <div id="LeftSection" class="lg:block hidden xl:w-[345px]">
                <div class="pt-4 pr-4 max-w-[320px]">
                    <router-link to="/user" class="flex items-center justify-start w-full cursor-pointer p-2 hover:bg-[#E5E6E9] rounded-md">
                        <img src="https://picsum.photos/id/87/300/320" class="rounded-full ml-l min-w-[38px] max-h-[38px]">
                        <!-- <div class="text-[15px] text-gray-800 font-extrabold pl-3">Raymond Mwebe</div> -->
                        <!-- <div class="text-[15px] text-gray-800 font-extrabold pl-3">{{userDetails.name}}</div> -->
                        <!-- <div v-if="userDetails" class="text-[15px] text-gray-800 font-extrabold pl-3">{{userDetails.name}}</div> -->

                    </router-link>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <AccountMultiple :size="32" fillColor="#5BD7C6" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Friends</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <Flag :size="32" fillColor="#F2682C" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Pages</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <ClockTimeTwoOutline :size="32" fillColor="#21AAFA" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Most Recent</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <AccountGroup :size="32" fillColor="#20A9FD" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Groups</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <StorefrontOutline :size="32" fillColor="#48C0D8" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Marketplace</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <TelevisionPlay :size="32" fillColor="#9739CF" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Watch</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <Restore :size="32" fillColor="#5BD7C6" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Memories</div>
                    </button>
                </div>
            </div>

            <div id="PostsSection" class="row-span-6 max-w-[600px] lg:mx-0 mx-auto overflow-auto">
                <CreatePostBox 
                image="https://picsum.photos/id/140/300/320"
                placeholder="What's on your mind Raymond?!" />
                <Post />
                <Post />
                <Post />
                <Post />
                <Post />
                <Post />
            </div>

            <div id="RightSection" class="pl-4 md:block hidden">
                <div class="mx-auto pt-4 max-w-[340px] min-w-[250px]">
                    <div class="flex items-center justify-between border-b border-b-gray-300">
                        <div class="font-semibold">Contacts</div>
                        <div class="flex items-center">
                            <div class="p-2 hover:bg-gray-300 rounded-full cursor-pointer">
                                <VideoImage :size="23" fillColor="#050505" />
                            </div>
                            <div class="p-2 hover:bg-gray-300 rounded-full cursor-pointer">
                                <Magnify :size="23" fillColor="#050505" />
                            </div>
                            <div class="p-2 hover:bg-gray-300 rounded-full cursor-pointer">
                                <DotsHorizontal :size="23" fillColor="#050505" />
                            </div>
                        </div>
                    </div>
                    <div class="h-[calc(100vh-115px)] overflow-auto pt-2">
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/140/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">John Doe</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/141/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Mpagi Ronald</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/142/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Nakimera Racheal</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/143/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Kironde Masembe</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/144/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Wanyana Leticia</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/145/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Mwebe Raphael</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/146/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Joel Mendez</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/147/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Cole Palmer</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/127/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Elon Musk</div>
                        </div>
                        <div class="flex items-center justify-start cursor-pointer hover:bg-[#E5E6E9] rounded-md py-2">
                            <img src="https://picsum.photos/id/117/300/320" alt="" class="rounded-full ml-1 min-w-[38px] max-h-[38px]">
                            <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Gwen Shotwell</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </MainNavLayout>
</template>
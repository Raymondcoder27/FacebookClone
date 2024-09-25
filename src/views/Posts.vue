<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";

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
</script>

<template>
  <MainNavLayout> 
    <div class="fixed w-full bg-[#F1F2F5] h-[100%]">
        <div class="grid grid-rows-3 grid-flow-col w-full mt-[56px] max-w-[1600px] h-[calc(100%-56px)] mx-auto px-4">
            <div id="LeftSection" class="lg:block hidden xl:w-[345px]">
                <div class="pt-4 pr-4 max-w-[320px]">
                    <router-link to="/" class="flex items-center justify-start w-full cursor-pointer p-2 hover:bg-[#E5E6E9] rounded-md">
                        <img src="https://picsum.photos/id/87/300/320" class="rounded-full ml-l min-w-[38px] max-h-[38px]">
                        <div class="text-[15px] text-gray-800 font-extrabold pl-3">Raymond Mwebe</div>
                    </router-link>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <AccountMultiple :size="40" fillColor="#5BD7C6" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Friends</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <Flag :size="40" fillColor="#F2682C" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Pages</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <ClockTimeTwoOutline :size="40" fillColor="#21AAFA" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Most Recent</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <AccountGroup :size="40" fillColor="#20A9FD" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Groups</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <StorefrontOutline :size="40" fillColor="#48C0D8" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Marketplace</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <TelevisionPlay :size="40" fillColor="#9739CF" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Watch</div>
                    </button>
                    <button class="flex items-center justify-start w-full cursor-pointer px-2 py-1.5 rounded-md hover:bg-[#E5E6E9]">
                        <Restore :size="40" fillColor="#5BD7C6" />
                        <div class="text-gray-800 font-extrabold pl-3 text-[15px]">Memories</div>
                    </button>
                </div>
            </div>

            <div id="PostsSection" class="row-span-6 max-w-[600px] lg:mx-0 mx-auto overflow-auto">
                Posts Section
            </div>

            <div id="RightSection" class="pl-4 md-block hidden">
                <div class="mx-auto pt-4 max-w-[340px] min-w-[250px]">
                    <div class="flex items-center justify-between border-b border-b-gray-300">
                        
                    </div>
                </div>
            </div>
        </div>
    </div>
  </MainNavLayout>
</template>
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

                </div>
            </div>
        </div>
    </div>
  </MainNavLayout>
</template>
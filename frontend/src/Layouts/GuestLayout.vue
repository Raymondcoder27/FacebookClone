<script setup>
import { RouterLink } from "vue-router";
import ApplicationLogo from "@/components/ApplicationLogo.vue";
</script>

<template>
  <div
    class="h-[100vh] flex flex-col items-center pt-[100px] sm:pt-0 bg-gray-100 z-40"
  >
    <div class="mt-12">
      <RouterLink to="/">
        <!-- <ApplicationLogo class="w-20 h-20 fill-current text-gray-500" /> -->
        <img src="/public/icons/FacebookLogo.png" class="w-[180px] mt-3" alt="" />
      </RouterLink>
    </div>

    <div
      class="w-full sm:max-w-md mt-5 px-6 py-4 bg-white shadow-md overflow-hidden sm:rounded-lg mb-5"
    >
      <slot />
    </div>
  </div>
</template>
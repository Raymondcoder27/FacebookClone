<script setup>
import { RouterLink } from 'vue-router';
import ApplicationLogo from '@/components/ApplicationLogo.vue';

</script>

<template>
    <div class="h-[100vh] flex flex-col pt-[100px] sm:pt-0 bg-gray-100 z-40">
        <div>
            <RouterLink>
                <ApplicationLogo class="w-20 h-20 fill-current text-gray-500" />
            </RouterLink>
        </div>

        <div class="w-full sm:max-w-md mt-6 px-6 py-4 bg-white shadow-md overflow-hidden sm:rounded-lg">
            <slot />
        </div>
    </div>
</template>
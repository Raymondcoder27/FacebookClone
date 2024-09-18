<script setup>
    import {ref} from 'vue'
    import Magnify from 'vue-material-design-icons/Magnify.vue'
    import Home from 'vue-material-design-icons/Home.vue'
    import HomeOutline from 'vue-material-design-icons/HomeOutline.vue'
    import TelevisionPlay from 'vue-material-design-icons/TelevisionPlay.vue'
    import StorefrontOutline from 'vue-material-design-icons/StorefrontOutline.vue'
    import AccountGroup from 'vue-material-design-icons/AccountGroup.vue'
    import ControllerClassicOutline from 'vue-material-design-icons/ControllerClassicOutline.vue'
    import DotsGrid from 'vue-material-design-icons/DotsGrid.vue'
    import FacebookMessenger from 'vue-material-design-icons/FacebookMessenger.vue'
    import Bell from 'vue-material-design-icons/Bell.vue'
    import Logout from 'vue-material-design-icons/Logout.vue'
    // import {Link} from '@inertiajs/vue3'

    import {useGeneralStore} from '@/stores/general'
    import { storeToRefs } from 'pinia'
    const useGeneral = useGeneralStore()
    const {isPostOverlay, isCropperModal, isImageDisplay} = storeToRefs(useGeneral)

    // const user = usePage().props.auth.user

    let showMenu = ref(false)
</script>

<template>
    <div id="MainNav" class="flex w-full fixed z-50 items-center justify-between bg-white shadow-xl border-b h-[56px]">
        <div id="NavLeft" class="flex items-center justify-start w-[260px]">
            <router-link href="/" class="pl-3 min-w-[55px]">FB</router-link>
            <div class="flex items-center justify-center p-1 rounded-full ml-2 bg-[#EFF2F5] h-[40px]">
                <Magnify class="p-1" :size="22" fillColor="#64676B"/>
            </div>  
        </div>
    </div>
    <slot/>
</template>
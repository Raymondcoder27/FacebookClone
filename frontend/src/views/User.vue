<script setup>
    import MainNavLayout from '@/Layouts/MainNavLayout.vue';
    import CreatePostBox from '@/Layouts/CreatePostBox.vue';
    import Post from '@/components/Post.vue'

    import Camera from 'vue-material-design-icons/Camera.vue';
    import Pen from 'vue-material-design-icons/Pen.vue';

    import { useGeneralStore } from '@/stores/general';
    import { storeToRefs } from 'pinia';
    const useGeneral = useGeneralStore();
    const {isCropperModal, isImageDisplay} = storeToRefs(useGeneral);

    defineProps({posts:Object, user:Object})
</script>

<template>
    <Head title="User" />
    <MainNavLayout>
        <div class="w-full pb-20 min-h-[100vh] bg-[#F1F2F5]">
            <div class="w-full bg-white">
                <div class="mx-auto pb-1 pt-[56px] max-w-[1100px]">
                    <img src="https://picsum.photos/id/45/2000/320" alt="" class="rounded-b-xl">

                    <div id="ProfileInfo" class="flex md:flex-row flex-col items-center justify-between px-4">
                        <div class="flex md:flex-row flex-col gap-4 md:-mt-6 -mt-16 items-center">
                            <div class="relative">
                                <img src="https://picsum.photos/id/200/500/500" alt="" class="rounded-full h-[165px] w-[165px] border-4 border-white">
                                <button
                                @click="isCropperModal = true" 
                                class="absolute right-0 bg-gray-200 hover:bg-gray-300 rounded-full p-1.5 cursor-pointer top-[100px]">
                                <Camera :size="25" />
                            </button>
                            </div>
                            <div class="md:mt-4 text-center md:text-left -mt-3">
                                <div class="text-md font-extrabold pt-1">
                                    Raymond Mwebe
                                </div>
                                <div class="font-bold text-gray-600">
                                    234 friends
                                </div>
                                <div class="flex md:justify-start justify-center md:-mt-1">
                                    <img src="https://picsum.photos/id/141/500/500" alt="" class="rounded-full -ml-1 boarder-white border-2 z-[10] w-[40px]">
                                    <img src="https://picsum.photos/id/142/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[9] w-[40px]">
                                    <img src="https://picsum.photos/id/143/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[8] w-[40px]">
                                    <img src="https://picsum.photos/id/144/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[7] w-[40px]">
                                    <img src="https://picsum.photos/id/145/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[6] w-[40px]">
                                    <img src="https://picsum.photos/id/146/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[5] w-[40px]">
                                    <img src="https://picsum.photos/id/147/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[4] w-[40px]">
                                    <img src="https://picsum.photos/id/149/500/500" alt="" class="rounded-full -ml-3 boarder-white border-2 z-[3] w-[40px]">
                                </div>
                            </div>
                        </div>

                        <RouterLink to="/"
                        class="flex justify-center w-7/12 md:w-[160px] items-baseline my-4 bg-gray-200 rounded-lg cursor-pointer">
                        <button class="flex items-center px-5 py-2 font-bold">
                            <Pen class="-mr-2 mr-1" :size="20" />Edit Profile
                        </button>
                    </RouterLink>
                    </div>

                    <div class="flex border-t h-[50px] py-[4px]">
                        <button
                        class="w-[85px]">
                            <div class="flex items-center justify-center text-blue-500 w-full font-bold rounded-lg font-bold text-[15px] h-[45px]">
                            </div>
                        </button>
                        <button class="flex items-center justify-center p-1 hover:bg-[#F2F2F2] w-[85px] text-[15px] ">

                        </button>
                    </div>
                </div>
            </div>
        </div>
    </MainNavLayout>
</template>
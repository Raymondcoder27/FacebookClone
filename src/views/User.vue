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
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </MainNavLayout>
</template>
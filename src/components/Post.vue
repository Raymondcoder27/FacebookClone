<script setup lang="ts">
    import {toRefs, reactive} from 'vue'
    import { useGeneralStore } from '@/stores/general';

    import AccountMultiple from 'vue-material-design-icons/AccountMultiple.vue';
    import ThumbUp from 'vue-material-design-icons/ThumbUp.vue';
    import Delete from 'vue-material-design-icons/Delete.vue';
    import Check from 'vue-material-design-icons/Check.vue';

    import { storeToRefs } from 'pinia';
    const useGeneral = useGeneralStore();

    const {isImageDisplay} = storeToRefs(useGeneral)

    const form = reactive({comment: null})

</script>

<template>
    <div id="Post" class="w-full bg-white rounded-lg my-4 shadow-md">
        <div class="flex items-center py-3 px-3">
            <button class="mr-2">
                <img  alt="" class="rounded-full ml-1 min-w-[42px] max-h-[42px]" src="https://picsum.photos/id/140/300/320">
            </button>
            <div class="flex items-center justify-between p-2 rounded-full w-full">
                <div>
                    <div class="font-extrabold text-[15px]">Raymond Mwebe Dev</div>
                    <div class="flex items-center text-xs text-gray-600">
                      14h  <AccountMultiple class="ml-1" :size="15" fillColor="#64676B" />
                    </div>
                </div>
                <div class="flex items-center">
                    <button class="rounded-full cursor-pointer p-1.5 hover:bg-[#F2F2F2]">
                        <Delete fillColor="#64676B" />
                    </button>
                </div>
            </div>
        </div>
        <div class="px-5 pb-2 font-semi-bold text-[17px]">
            Lorem, ipsum dolor sit amet consectetur adipisicing elit. 
        </div>
        <img src="https://picsum.photos/id/189/800/800" alt="" class="mx-auto cursor-pointer">
        <div id="Likes" class="px-5">
            <div class="flex items-center justify-between py-3 border-b">
                <ThumbUp fillColor="#1D72E2" />
                <div class="text-sm text-gray-600 font-semi-bold">5 comments</div>
            </div>
        </div>

        <div id="Comments" class="px-3">
            <div id="CreateComment" class="flex items-center justify-between py-2">
                <div class="flex items-center w-full">
                    <RouterLink>
                        <img src="" alt="" class="rounded-full ml-1 min-w-[36px] max-h-[36px]">
                    </RouterLink>
                </div>
            </div>
        </div>
    </div>
</template>
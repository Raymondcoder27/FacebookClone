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
                    <RouterLink to="/">
                        <img src="https://picsum.photos/id/199/800/800" alt="" class="rounded-full ml-1 min-w-[36px] max-h-[36px]">
                    </RouterLink>
                    <div class="flex items-center justify-center p-2 rounded-full w-full bg-[#EFF2F5]">
                        <input
                        v-model="form.comment" 
                        type="text" 
                        class="w-full mx-1 border-none p-0 text-sm ring-0 focus:outline-none bg-[#EFF2F5] placeholder-[#64676B]"
                        placeholder="Write a public comment...">
                        <button
                    type="button"
                    class="flex items-center text-sm pl-2 pr-3.5 rounded-full bg-blue-500 hover:bg-blue-800 text-white font-bold">
                        <Check fillColor="#FFF" :size="20" /> Send
                    </button>
                    </div>
                </div>
            </div>

            <div id="Comments" class="max-h-[120px] overflow-auto pb-2 px-4">
                <div class="flex items-center justify-between pb-2">
                    <div class="flex items-center w-full mb-1">
                        <RouterLink to="/">
                            <img src="https://picsum.photos/id/299/800/800" alt="" class="rounded-full min-w-[36px] max-h-[36px]">
                        </RouterLink>
                        <div class="flex items-center w-full">
                            <div class="flex items-center text-xs p-2 rounded-lg w-full bg-[#EFF2F5]">
                                This is a comment
                            </div>
                            <button class="rounded-full ml-2 cursor-pointer">
                                <Delete fillColor="#64676B" />
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
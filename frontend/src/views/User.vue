<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import Camera from "vue-material-design-icons/Camera.vue";
import Pen from "vue-material-design-icons/Pen.vue";
import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import api from "@/config/api";

const posts = ref([])

const useGeneral = useGeneralStore();
const { isCropperModal, isImageDisplay } = storeToRefs(useGeneral);

const userDetails = ref(null);
defineProps({ posts: Object, user: Object });
const authStore = useAuthStore();

onMounted(async () => {
  await getUserDetails();
  // await api.get("/posts");
  await getPosts();
});

const getPosts = async() => {
  try{
    const response = await api.get("/posts")
    posts.value = response.data
  }catch(error){
    console.error('error fetching posts', error)
  }
}
const getUserDetails = async () => {
  const token = authStore.token;
  if (!token) {
    console.error("No token found");
    return;
  }

  try {
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
  <Head title="User" />
  <MainNavLayout>
    <div class="w-full pb-20 min-h-[100vh] bg-[#F1F2F5]">
      <div class="w-full bg-white">
        <div class="mx-auto pb-1 pt-[56px] max-w-[1100px]">
          <img
            src="https://picsum.photos/id/45/2000/320"
            alt=""
            class="rounded-b-xl"
          />
          <!-- <img
            :src="userDetails?.image"
            alt=""
            class="rounded-b-xl"
          /> -->
          <div
            id="ProfileInfo"
            class="flex md:flex-row flex-col items-center justify-between px-4"
          >
            <div
              class="flex md:flex-row flex-col gap-4 md:-mt-6 -mt-16 items-center"
            >
              <div class="relative">
                {{ userDetails }}
                <img
                  :src="userDetails?.image"
                  alt=""
                  class="rounded-full h-[165px] w-[165px] border-4 border-white"
                />
                <button
                  @click="isCropperModal = true"
                  class="absolute right-0 bg-gray-200 hover:bg-gray-300 rounded-full p-1.5 cursor-pointer top-[100px]"
                >
                  <Camera :size="25" />
                </button>
              </div>
              <div class="md:mt-4 text-center md:text-left -mt-3">
                <!-- <div class="text-md font-extrabold pt-1">Raymond Mwebe</div> -->
                <div v-if="userDetails" class="text-md font-extrabold pt-1">
                  {{ userDetails?.name }}
                </div>
                <div class="font-bold text-gray-600">234 friends</div>
                <div class="flex md:justify-start justify-center md:-mt-1">
                  <img
                    src="https://picsum.photos/id/141/500/500"
                    alt=""
                    class="rounded-full -ml-1 boarder-white border-2 z-[10] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/142/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[9] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/143/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[8] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/144/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[7] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/145/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[6] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/146/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[5] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/147/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[4] w-[40px]"
                  />
                  <img
                    src="https://picsum.photos/id/149/500/500"
                    alt=""
                    class="rounded-full -ml-3 boarder-white border-2 z-[3] w-[40px]"
                  />
                </div>
              </div>
            </div>

            <RouterLink
              to="/user"
              class="flex justify-center w-7/12 md:w-[160px] items-baseline my-4 bg-gray-200 rounded-lg cursor-pointer"
            >
              <button class="flex items-center px-5 py-2 font-bold">
                <Pen class="-mr-2 mr-1" :size="20" />Edit Profile
              </button>
            </RouterLink>
          </div>

          <div class="flex border-t h-[50px] py-[4px]">
            <button class="w-[85px]">
              <div
                class="flex items-center justify-center text-blue-500 w-full font-bold rounded-lg text-[15px] h-[45px] cursor-pointer"
              >
                Posts
              </div>
              <div class="border-b-4 border-blue-400 rounded-md"></div>
            </button>
            <button
              class="flex items-center justify-center p-1 hover:bg-[#F2F2F2] w-[85px] text-[15px] text-[15px] h-[48px] font-bold rounded-lg mx-l cursor-pointer"
            >
              About
            </button>
            <button
              class="flex items-center justify-center p-1 hover:bg-[#F2F2F2] w-[85px] text-[15px] text-[15px] h-[48px] font-bold rounded-lg mx-l cursor-pointer"
            >
              Friends
            </button>
            <button
              class="flex items-center justify-center p-1 hover:bg-[#F2F2F2] w-[85px] text-[15px] text-[15px] h-[48px] font-bold rounded-lg mx-l cursor-pointer"
            >
              Videos
            </button>
            <button
              class="flex items-center justify-center p-1 hover:bg-[#F2F2F2] w-[85px] text-[15px] text-[15px] h-[48px] font-bold rounded-lg mx-l cursor-pointer"
            >
              Photos
            </button>
            <button
              class="flex items-center justify-center p-1 hover:bg-[#F2F2F2] w-[85px] text-[15px] text-[15px] h-[48px] font-bold rounded-lg mx-l cursor-pointer"
            >
              Check-Ins
            </button>
          </div>
        </div>
      </div>

      <div
        class="flex flex-col md:flex-row w-full justify-between md:px-0 max-w-[1100px] h-[calc(100%-56px)] px-2 mx-auto"
      >
        <div id="LeftSection" class="w-full mt-4 mr-4 md:w-5/12">
          <div class="bg-white p-3 rounded-lg shadow-lg">
            <!-- <div class="font-extrabold pb-2 text-xl"> -->
            <div class="font-extrabold pb-2 text-xl">Intro</div>
            <div class="pb-5">
              <button class="w-full bg-gray-200 rounded-lg p-2 font-bold">
                Add bio
              </button>
            </div>
            <div class="pb-5">
              <button class="w-full bg-gray-200 rounded-lg p-2 font-bold">
                Edit details
              </button>
            </div>
            <div class="pb-5">
              <button class="w-full bg-gray-200 rounded-lg p-2 font-bold">
                Add hobbies
              </button>
            </div>
            <div class="pb-5">
              <button class="w-full bg-gray-200 rounded-lg p-2 font-bold">
                Add features
              </button>
            </div>
            <!-- </div> -->
          </div>

          <div class="bg-white p-3 mt-4 rounded-lg shadow-lg">
            <div class="font-extrabold pb-2 text-xl">Photos</div>
            <div class="flex flex-wrap items-center justify-start w-full">
              <span v-for="photo in posts" :key="photo" class="w-1/3">
                <!-- <img
                  @click="
                    isImageDisplay = 'https://picsum.photos/id/78/300/300'
                  "
                  src="https://picsum.photos/id/78/300/300"
                  class="aspect-square object-cover p-1 rounded-lg cursor-pointer"
                /> -->
                <img
                  v-if="photo.image"
                  @click="isImageDisplay = 'userDetails?.image'"
                  :src="photo?.image"
                  class="aspect-square object-cover p-1 rounded-lg cursor-pointer"
                />
                {{ userDetails }}
              </span>
            </div>
          </div>
        </div>

        <div id="PostsSection" class="w-full md:w-7/12 overflow-auto">
          <!-- <CreatePostBox
            image="https://picsum.photos/id/140/300/320"
            :placeholder="'What\'s on your mind ' + userDetails?.name + '?!'"
          /> -->
          <CreatePostBox
            :image="userDetails?.image"
            :placeholder="'What\'s on your mind ' + userDetails?.name + '?!'"
          />

          <div v-for="post in posts" :key="post">
          <Post :user="post.user" :post="post.text" :comments="post.comments"/>
          </div>
          <!-- <Post />
          <Post />
          <Post />
          <Post />
          <Post /> -->
        </div>
      </div>
    </div>
  </MainNavLayout>
</template>
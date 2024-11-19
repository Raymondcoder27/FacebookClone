<script setup lang="ts">
import { toRefs, reactive } from "vue";
import { useGeneralStore } from "@/stores/general";

import AccountMultiple from "vue-material-design-icons/AccountMultiple.vue";
import ThumbUp from "vue-material-design-icons/ThumbUp.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import Check from "vue-material-design-icons/Check.vue";

import { storeToRefs } from "pinia";
import { useAuthStore } from "@/stores/auth";
import api from "@/config/api";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";

const emit = defineEmits();

const useGeneral = useGeneralStore();

const { isImageDisplay } = storeToRefs(useGeneral);

const form = reactive({ comment: null });

const userDetails = ref(null);
const authStore = useAuthStore();
const router = useRouter();

// const posts = ref([]);

// const getPosts = async () => {
//   try {
//     const response = await api.get("/posts");
//     posts.value = response.data.data;
//   } catch (error) {
//     console.error("Failed to fetch posts", error);
//   }
// };

const getUserDetails = async () => {
  const token = authStore.token;
  if (!token) {
    console.error("No token found");
    return;
  }

  try {
    const response = await api.get("/validate", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    userDetails.value = response.data; // Storing user details
  } catch (error) {
    console.error("Failed to fetch user details", error);
  }
};

// const props = defineProps({
//   user: Object,
//   post: Object,
//   comments: Object,
// });

const props = defineProps({
  user: Object,
  post: Object,
  comments: {
    type: Array,
    default: () => [] // Initialize as an empty array if not provided
  },
});


const { post, user, comments } = toRefs(props);

const CreateComment = async () => {
  try {
    const response = await api.post("/create-comment", {
      post_id: post.value.id,
      text: form.comment,
    });
    comments.value = response.data.comment;
    alert(JSON.stringify(response.data.comment.text));

    console.log('API Response:', response.data); // Check response structure
  } catch (error) {
    console.error("error creating tweet:", error);
  }
};

const deleteComment = async (id) => {
  try {
    await api.delete("/delete-comment/" + id);
  } catch (error) {
    console.error("error deleting comment", error);
  }
};

const deletePost = async (id) => {
  try {
    await api.delete("/delete-post/" + id);
    emit("postDeleted");
  } catch (error) {
    console.error("error deleting post", error);
  }
};

const logout = async () => {
  await authStore.logout();
  router.push("/login");
};
// console.log("Fetched posts:", posts.value);

// Ensure this runs when the component mounts
onMounted(async () => {
  await getUserDetails();
  // await getPosts();
});
</script>

<template>
  <div id="Post" class="w-full bg-white rounded-lg my-4 shadow-md">
    <div class="flex items-center py-3 px-3">
      <button class="mr-2">
        <!-- <img  alt="" class="rounded-full ml-1 min-w-[42px] max-h-[42px]" src="https://picsum.photos/id/140/300/320"> -->
        <img
          alt=""
          class="rounded-full ml-1 min-w-[42px] max-h-[42px]"
          :src="userDetails?.image"
        />
        <!-- <img src="./public/user-placeholder.png" alt="" class="rounded-full ml-1 min-w-[42px] max-h-[42px]"> -->
      </button>

    
      <div class="flex items-center justify-between p-2 rounded-full w-full">
        <!-- {{posts}} -->
        <div>
          <!-- <div class="font-extrabold text-[15px]">Raymond Mwebe Dev</div> -->
          <div class="font-extrabold text-[15px]">{{ userDetails?.name }}</div>
          <div class="flex items-center text-xs text-gray-600">
            {{ post.created_at }} h
            <!-- 14h -->
            <AccountMultiple class="ml-1" :size="15" fillColor="#64676B" />
          </div>
        </div>
        <div class="flex items-center">
          <button
            @click="deletePost(post.id)"
            class="rounded-full cursor-pointer p-1.5 hover:bg-[#F2F2F2]"
          >
            <Delete fillColor="#64676B" />
            <!-- Delete Post -->
          </button>
        </div>
      </div>
    </div>
    <div class="px-5 pb-2 font-semi-bold text-[17px]">
      {{ post.text }}
      <!-- Lorem ipsum dolor sit. -->
    </div>
    <!-- <img
      @click="isImageDisplay = 'https://picsum.photos/id/189/800/800'"
      :src="post?.image ? `data:image/jpeg;base64,${post.image}` : 'default-placeholder-image.png'"
      alt=""
      class="mx-auto cursor-pointer"
    /> -->
    <img
      @click="isImageDisplay = 'https://picsum.photos/id/188/800/900'"
      src="https://picsum.photos/id/188/800/900"
      alt=""
      class="mx-auto cursor-pointer"
    />
    <div id="Likes" class="px-5">
      <div class="flex items-center justify-between py-3 border-b">
        <!-- <ThumbUp fillColor="#1D72E2" /> -->
        <ThumbUp fillColor="#babfc5" class="cursor-pointer" />
        <div class="text-sm text-gray-600 font-semi-bold">0 comments</div>
      </div>
    </div>

    <div id="Comments" class="px-3">
      <div id="CreateComment" class="flex items-center justify-between py-2">
        <div class="flex items-center w-full">
          <RouterLink to="/">
            <img
              :src="userDetails?.image"
              alt=""
              class="rounded-full ml-1 min-w-[36px] max-h-[36px]"
            />
          </RouterLink>
          <div
            class="flex items-center justify-center p-2 rounded-full w-full bg-[#EFF2F5]"
          >
            <input
              v-model="form.comment"
              type="text"
              class="w-full mx-1 border-none p-0 text-sm ring-0 focus:outline-none bg-[#EFF2F5] placeholder-[#64676B]"
              placeholder="Write a public comment..."
            />
            <button
              type="button"
              @click="CreateComment"
              class="flex items-center text-sm pl-2 pr-3.5 rounded-full bg-blue-500 hover:bg-blue-800 text-white font-bold"
            >
              <Check fillColor="#FFF" :size="20" /> Send
            </button>
          </div>
        </div>
      </div>

      <div
        v-if="comments"
        id="Comments"
        class="max-h-[120px] overflow-auto pb-2 px-4"
      >
        <div
          class="flex items-center justify-between pb-2"
          v-for="comment in comments"
          :key="comment.ID"
        >
          <div class="flex items-center w-full mb-1">
            <RouterLink to="/">
              <img
                @click="isImageDisplay = comment.user.image"
                :src="userDetails?.image"
                alt=""
                class="rounded-full min-w-[36px] max-h-[36px]"
              />
            </RouterLink>
            <div class="flex items-center w-full">
              <div
                class="flex items-center text-xs p-2 rounded-lg w-full bg-[#EFF2F5]"
              >
                <!-- This is a comment -->
                {{ comment.text }}
              </div>
              <!-- <button class="rounded-full ml-2 cursor-pointer">
                <Delete fillColor="#64676B" />
              </button> -->
              <button
                v-if="userDetails?.id === comment.user.id"
                @click="deleteComment(comment.id)"
                class="rounded-full p-1.5 ml-2 cursor-pointer hover:bg-[#F2F2F2]"
              >
                <Delete fillColor="#64676B" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
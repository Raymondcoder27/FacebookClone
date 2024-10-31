<script setup lang="ts">
// import { toRefs, reactive } from "vue";
// import { useGeneralStore } from "@/stores/general";

// import AccountMultiple from "vue-material-design-icons/AccountMultiple.vue";
// import ThumbUp from "vue-material-design-icons/ThumbUp.vue";
// import Delete from "vue-material-design-icons/Delete.vue";
// import Check from "vue-material-design-icons/Check.vue";

// import { storeToRefs } from "pinia";
// import { useAuthStore } from "@/stores/auth";
// import api from "@/config/api";
// import { onMounted, ref } from "vue";
// import { useRouter } from "vue-router";

// const useGeneral = useGeneralStore();

// const { isImageDisplay } = storeToRefs(useGeneral);

// const form = reactive({ comment: null });

// const userDetails = ref(null);
// const authStore = useAuthStore();
// const router = useRouter();

// const posts = ref([]);

// const getPosts = async () => {
//   try {
//     const response = await api.get("/posts");
//     posts.value = response.data;
//   } catch (error) {
//     console.error("Failed to fetch posts", error);
//   }
// };

// const getUserDetails = async () => {
//   const token = authStore.token;
//   if (!token) {
//     console.error("No token found");
//     return;
//   }

//   try {
//     const response = await api.get("/validate", {
//       headers: {
//         Authorization: `Bearer ${token}`,
//       },
//     });
//     userDetails.value = response.data; // Storing user details
//   } catch (error) {
//     console.error("Failed to fetch user details", error);
//   }
// };

// const props = defineProps({
//   user: Object,
//   post: Object,
//   comments: Object,
// });

// const { post, user, comments } = toRefs(props);

// const CreateComment = async () => {
//   try {
//     await api.post("/create-comment", {
//       post_id: post.value.id,
//       text: form.comment,
//     });
//   } catch (error) {
//     console.error("error creating tweet:", error);
//   }
// };

// const deleteComment = async (id) => {
//   try {
//     await api.delete("/delete-comment/" + id);
//   } catch (error) {
//     console.error("error deleting comment", error);
//   }
// };

// const deletePost = async (id) => {
//   try {
//     await api.delete("/delete-post/" + id);
//   } catch (error) {
//     console.error("error deleting post", error);
//   }
// };

// const logout = async () => {
//   await authStore.logout();
//   router.push("/login");
// };
// console.log("Fetched posts:", posts.value);

// // Ensure this runs when the component mounts
// onMounted(async () => {
//   await getUserDetails();
//   await getPosts();
// });

























import { ref } from 'vue';

const posts = ref([]);

fetchPosts(); // Call your function to fetch posts here

async function fetchPosts() {
    try {
        const response = await fetch('/api/posts');
        const result = await response.json();
        posts.value = result.data; // Correctly assign data here
    } catch (error) {
        console.error(error);
    }
}
console.log("Fetched posts:", posts.value); // Shows data in array format

</script>

<template>
  <div id="Post" class="w-full bg-white rounded-lg my-4 shadow-md">
    <div v-for="post in posts" :key="post.id">
        <p>{{ post.text }}</p>
        <p>Posted by: {{ post.user.name }}</p>
    </div>
  </div>
</template>
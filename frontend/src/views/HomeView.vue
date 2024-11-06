<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';

import api from "@/config/api";

const userDetails = ref(null);
defineProps({ posts: Object, user: Object });
const authStore = useAuthStore();


const posts = ref([])

onMounted(async () => {
  await getUserDetails();
});

  const getUserDetails = async () => {
  const token = authStore.token;
  if (!token) {
    console.error("No token found");
    return;
  }
  
  try {
    // const response = await api.get("/user", {
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



const getPosts = async() => {
  try{
    const response = await api.get("/posts")
    posts.value = response.data.data
  }catch(error){
    console.error('error fetching posts', error)
  }
}

onMounted(() => {
  // getUserDetails();
  getPosts();
  // getComments();
});
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->
  </main>
</template>

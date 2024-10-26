<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';

import api from "@/config/api";

const userDetails = ref(null);
defineProps({ posts: Object, user: Object });
const authStore = useAuthStore();

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
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->
  </main>
</template>

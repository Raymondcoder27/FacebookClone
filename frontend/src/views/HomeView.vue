<script setup>
// import TheWelcome from '../components/TheWelcome.vue'
import { onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore()

onMounted(() => {
  // const userDetails = api.get("/validate");

  const userDetails = authStore.validate();
  // const user = ref(
  // localStorage.getItem("userDetails") ? JSON.parse(localStorage.getItem("userDetails")) : null;
  console.log(userDetails);
});
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->
  </main>
</template>

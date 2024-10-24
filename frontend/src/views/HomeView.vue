<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();
const hasReloaded = ref(false); // Track if the page has been reloaded

onMounted(() => {
  // Fetch user details or perform any necessary setup
  const userDetails = authStore.validate();
  console.log(userDetails);

  // Reload the page only if it hasn't been reloaded before
  if (!hasReloaded.value) {
    hasReloaded.value = true; // Set to true to prevent future reloads
    setTimeout(() => {
      window.location.reload(); // Reload after a delay
    }, 1000);
  }
});
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->
  </main>
</template>

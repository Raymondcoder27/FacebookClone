<template>
  <div>
    <input type="file" @change="getUploadedImage" />
    <img v-if="imageDisplay" :src="imageDisplay" alt="Preview" />
  </div>
</template>
  
  <script setup>
import { ref } from "vue";

const imageDisplay = ref(null);

const getUploadedImage = (e) => {
  if (e.target.files && e.target.files[0]) {
    const file = e.target.files[0];
    imageDisplay.value = URL.createObjectURL(file); // Generate preview URL
    console.log("File selected:", file);
  } else {
    console.error("No file selected");
  }
};
</script>
  
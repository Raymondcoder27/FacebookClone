<template>
  <div>
    <input type="file" @change="getUploadedImage" accept="image/*" />
    <img v-if="imageDisplay" :src="imageDisplay" alt="Preview" />
    <p v-if="uploadStatus">{{ uploadStatus }}</p>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios"; // Axios for HTTP requests

const imageDisplay = ref(null); // To store the preview URL
const selectedFile = ref(null); // To store the file for backend upload
const uploadStatus = ref(""); // To show upload status

// Handle file selection and upload
const getUploadedImage = async (e) => {
  const file = e.target.files[0];
  if (file && file.type.startsWith("image/")) {
    // Create image preview
    if (imageDisplay.value) {
      URL.revokeObjectURL(imageDisplay.value); // Clean up previous preview
    }
    imageDisplay.value = URL.createObjectURL(file);
    selectedFile.value = file; // Store file for upload

    try {
      // Create FormData and append the file
      const formData = new FormData();
      formData.append("image", file);

      // POST to backend
      const response = await api.post("create-post", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      uploadStatus.value = `Upload successful: ${response.data.message}`;
    } catch (error) {
      console.error("Upload failed:", error);
      uploadStatus.value = "Upload failed. Please try again.";
    }
  } else {
    console.error("Invalid file type. Please upload an image.");
    uploadStatus.value = "Invalid file type. Please upload an image.";
  }
};
</script>

<style>
img {
  max-width: 100%;
  height: auto;
  display: block;
  margin-top: 10px;
}
p {
  margin-top: 10px;
  color: #333;
}
</style>

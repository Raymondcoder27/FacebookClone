<script setup>
import { ref } from "vue";
import { Cropper, CircleStencil } from "vue-advanced-cropper";
import "vue-advanced-cropper/dist/style.css";

import Close from "vue-material-design-icons/Close.vue";
import Plus from "vue-material-design-icons/Plus.vue";

const emit = defineEmits(["showModal"]);

let fileInput = ref(null);
let cropper = ref(null);
let uploadedImage = ref(null);
let showCropModal = ref(false); // Flag to control visibility of the crop modal
let croppedImageData = {
  file: null,
  imageUrl: null,
  height: null,
  width: null,
  left: null,
  top: null,
};

const getUploadedImage = (e) => {
  const file = e.target.files[0];
  if (file) {
    uploadedImage.value = URL.createObjectURL(file);
    showCropModal.value = true; // Show the crop modal when an image is uploaded
  }
};

const crop = () => {
  const { coordinates, canvas } = cropper.value.getResult();
  croppedImageData.imageUrl = canvas.toDataURL();

  let data = new FormData();
  data.append('image', fileInput.value.files[0] || '');
  data.append('height', coordinates.height || '');
  data.append('width', coordinates.width || '');
  data.append('left', coordinates.left || '');
  data.append('top', coordinates.top || '');

  // Handle sending to backend here

  // Reset state and close modal
  uploadedImage.value = null; // Reset uploaded image
  showCropModal.value = false; // Close the crop modal
  emit('showModal', false); // Close the main modal
};

const closeModal = () => {
  uploadedImage.value = null; // Reset uploaded image
  showCropModal.value = false; // Close the crop modal
  emit('showModal', false); // Close the main modal
};
</script>

<template>
  <div class="fixed z-50">
    <div class="fixed inset-0 bg-white bg-opacity-60"></div>
    <div class="fixed inset-0 z-10 overflow-y-auto">
      <div class="flex flex-col min-h-full justify-center items-center py-2">
        <div class="transform overflow-hidden rounded-lg bg-white shadow-2xl transition-all max-w-xl">
          <div class="flex items-center py-4 border-b border-b-gray-300">
            <div class="text-[22px] font-extrabold w-full text-center">
              Update Profile Picture
            </div>
            <div
              @click="closeModal"
              class="absolute right-3 rounded-full bg-gray-200 hover:bg-gray-300 cursor-pointer p-1.5"
            >
              <Close :size="28" fillColor="#5E6771" />
            </div>
          </div>

          <div class="flex items-center bg-white px-4 pb-4">
            <div>
              <div class="my-4">
                <label for="image" class="flex items-center justify-center bg-[#E7F3FF] hover:bg-[#DBE7F2] text-[#1977F2] font-bold p-2 rounded-lg w-full cursor-pointer">
                  <Plus :size="20" /> Upload Photo
                </label>
                <input type="file" id="image" ref="fileInput" @change="getUploadedImage" class="hidden" />
              </div>

              <!-- Show the cropper only if uploadedImage is present -->
              <div v-if="showCropModal" class="w-[350px] mx-auto">
                <Cropper
                  class="object-cover"
                  ref="cropper"
                  :stencil-component="CircleStencil"
                  :src="uploadedImage"
                />
              </div>

              <div class="flex gap-4" :class="uploadedImage ? 'pt-4' : ''">
                <button
                  type="button"
                  @click="closeModal"
                  class="w-full justify-center rounded-md py-2 text-gray-600 hover:text-gray-800 font-bold hover:shadow-sm hover:bg-gray-200 focus:outline-none focus:ring-0"
                >
                  Cancel
                </button>
                <button
                  type="button"
                  @click="crop"
                  v-if="showCropModal"
                  class="w-full justify-center rounded-md py-2 text-white bg-blue-600 hover:bg-blue-700 font-bold hover:shadow-sm focus:outline-none focus:ring-0"
                >
                  Crop
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

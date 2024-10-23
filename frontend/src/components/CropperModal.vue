<script setup>
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

import { Cropper, CircleStencil } from "vue-advanced-cropper";
import "vue-advanced-cropper/dist/style.css";

import Close from "vue-material-design-icons/Close.vue";
import Plus from "vue-material-design-icons/Plus.vue";

const emit = defineEmits(["showModal"]);

// import { watch } from 'vue';

// watch(uploadedImage, (newValue) => {
//   if (cropper.value) {
//     cropper.value.refresh();  // Ensure cropper reinitializes
//   }
// });


let fileInput = ref(null);
let cropper = ref(null);
let uploadedImage = ref(null);
let croppedImageData = {
  file: null,
  imageUrl: null,
  height: null,
  width: null,
  left: null,
  top: null,
};
const file = ref(null)

const getUploadedImage = (e) => {
  uploadedImage.value = URL.createObjectURL(e.target.files[0]);
  file.value = e.target.files[0];
};

// const getUploadedImage = (e) => {
//         ImageDisplay.value = URL.createObjectURL(e.target.files[0]);
//         form.image = e.target.files[0];
//     }

// const crop = () => {
//   const { coordinates, canvas } = cropper.value.getResult()
//   croppedImageData.imageUrl = canvas.toDataURL()

//   let data = new FormData();
//   // data.append('image', fileInput.value.files[0] || '')
//     data.append('image', file.value || '')
//     data.append('height', coordinates.height || '')
//     data.append('width', coordinates.width || '')
//     data.append('left', coordinates.left || '')
//     data.append('top', coordinates.top || '')

//     //send to backend
//     emit ('showModal', false)
// }

const crop = () => {
  if (!cropper.value) {
    console.error("No cropper instance found");
    return;
  }
  
  const { coordinates, canvas } = cropper.value.getResult();
  if (!coordinates) {
    console.error("No cropping coordinates found");
    return;
  }

  croppedImageData.imageUrl = canvas.toDataURL();
  let data = new FormData();
  data.append('image', file.value || '');
  data.append('height', coordinates.height || '');
  data.append('width', coordinates.width || '');
  data.append('left', coordinates.left || '');
  data.append('top', coordinates.top || '');
  //send to backend
  emit('showModal', false);
};

</script>

<template>
    <div class="fixed z-50">
      <div class="fixed inset-0 bg-white bg-opacity-60 z-40"></div>
      <div class="fixed inset-0 z-50 overflow-y-auto">
        <div class="flex flex-col min-h-full justify-center items-center py-2">
          <!-- <div class="transform overflow-hidden rounded-lg bg-white shadow-2xl transition-all max-w-xl"> -->
          <div class="transform overflow-hidden rounded-lg bg-white shadow-2xl transition-all max-w-xl">
            <div class="flex items-center py-4 border-b border-b-gray-300">
              <div class="text-[22px] font-extrabold w-full text-center">
                Update Profile picture
              </div>
              <div
              @click="$emit('showModal', false)"
               class="absolute right-3 rounded-full bg-gray-200 hover:bg-gray-300 cursor-pointer p-1.5">
              <Close :size="28" fillColor="#5E6771"/>
              </div>
            </div>

            <div class="flex items-center bg-white px-4 pb-4">
              <div>
                <div class="my-4">
                  <label for="image" class="flex items-center justify-center bg-[#E7F3FF] hover:bg-[#DBE7F2] text-[#1977F2] font-bold p-2 rounded-lg w-full cursor-pointer">
                    <Plus  :size="20"/>Upload Photo
                  </label>
                  <!-- <input type="file" id="image" ref="fileInput" class="hidden" @change="getUploadedImage" > -->
                  <input type="file" id="image" ref="fileInput" class="hidden" @change="getUploadedImage($event)" >
                  <!-- <input id="image" type="file" class="hidden" @input="getUploadedImage($event)"> -->
                </div>

                <div class="w-[350px] mx-auto">
                  <Cropper 
                  class="object-cover"
                  ref="cropper"
                  :stencil-component="CircleStencil"
                  :src="uploadedImage" />
                </div>

                <div class="flex gap-4" :class="uploadedImage ? 'pt-4': ''">
                  <button 
                  type="button"
                  @click="emit('showModal', false)"
                  class="w-full justify-center rounded-md py-2 text-gray-600 hover:text-gray-800 font-bold hover:shadow-sm hover:bg-gray-200 focus:outline-none focus:ring-0">
                Cancel</button>
                <button
                v-if="uploadedImage"
                @click="crop"
                type="button"
                class="w-full rounded-md bg-blue-500 py-2 text-white font-bold shadow-sm hover:bg-blue-600 focus:outline-none focus:ring-0">
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
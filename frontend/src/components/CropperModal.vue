<script setup>
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

import { Cropper, CircleStencil } from "vue-advanced-cropper";
import "vue-advanced-cropper/dist/style.css";

import Close from "vue-material-design-icons/Close.vue";
import Plus from "vue-material-design-icons/Plus.vue";

const emit = defineEmits(["showModal"]);

let fileInput = ref(null);
let cropper = ref(null);
let uploadedImage = ref(null);
let croppedImage = {
  file: null,
  imageUrl: null,
  height: null,
  width: null,
  left: null,
  top: null,
};

const getUploadedImage = (e) => {
  const file = e.target.files[0];
  uploadedImage.value = URL.createObjectURL(file);
};

const crop = () => {
  const {coordinates, canvas} = cropper.value.getResult()
  croppedImageData.imageUrl = canvas.toDataURL()

  let data = new FormData()
    data.append('image', fileInput.value.files[0] || '')
    data.append('height', coordinates.height || '')
    data.append('width', coordinates.width || '')
    data.append('left', coordinates.left || '')
    data.append('top', coordinates.top || '')
}

//send to backend
    emit ("showModal", false)
</script>

<template>
    <div class="fixed z-50">
      <div class="fixed inset-0 bg-white bg-opacity-60"></div>
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="flex flex-col min-h-full justify-center items-center py-2"></div>
      </div>
    </div>
</template>
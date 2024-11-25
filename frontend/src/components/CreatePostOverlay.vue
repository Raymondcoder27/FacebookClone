<script setup>
import { ref, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";

import VideoImage from "vue-material-design-icons/VideoImage.vue";
import Image from "vue-material-design-icons/Image.vue";
import EmoticonOutline from "vue-material-design-icons/EmoticonOutline.vue";
import Close from "vue-material-design-icons/Close.vue";
import ChevronDown from "vue-material-design-icons/ChevronDown.vue";
import Earth from "vue-material-design-icons/Earth.vue";
import DotsHorizontal from "vue-material-design-icons/DotsHorizontal.vue";

import { defineEmits } from "vue";
import { onMounted } from "vue";
import { useAuthStore } from "@/stores/auth";
import api from "@/config/api";

const userDetails = ref(null);
const authStore = useAuthStore();
const router = useRouter();

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
const useGeneral = useGeneralStore();
const { isPostOverlay, postAdded } = storeToRefs(useGeneral);

const emit = defineEmits(["showModal"]);

let imageDisplay = ref("");

// let imageFile = ref(null);

// const posts = ref([]);

const form = reactive({
  text: null,
  image: null,
});
let error = ref(null);

const createPost = async () => {
  error.value = null;
  const formData = new FormData();
  formData.append("text", form.text);
  // Only append the image if it exists
  // if (form.image) {
  // formData.append("image", form.image);
  console.log("Image appended to FormData:", formData.get("image"));
  // } else {
  // console.log("No image found when creating post");
  // }

  try {
    const response = await api.post("/create-post", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${authStore.token}`,
      },
    });

    //emit post created to update pinia store
    emitPostCreated();

    // Close the modal
    useGeneral.isPostOverlay = false;
    // isPostOverlay.value = false;

    // getPosts();
  } catch (error) {
    console.error("Failed to create post", error);
    error.value = "Failed to create post. Please try again.";
  }
};

// const getUploadedImage = (e) => {
//   try {
//     if (e.target.files && e.target.files[0]) {
//       console.log("Selected file:", e.target.files[0]);
//       const selectedFile = e.target.files[0];
//       imageDisplay.value = URL.createObjectURL(selectedFile); // For preview
//       form.image = selectedFile; // Bind the file to the reactive form object
//     } else {
//       console.log("No file selected or file input is empty");
//     }
//   } catch (error) {
//     console.error("Error in getUploadedImage:", error);
//   }
// };

const getUploadedImage = (e) => {
  imageDisplay.value = URL.createObjectURL(e.target.files[0]);
  form.image = e.target.files[0];
};

const clearImage = () => {
  imageDisplay.value = null; // Clear the preview
  form.image = null; // Reset the form's image field
};

const getUserDetails = async () => {
  const token = authStore.token;
  if (!token) {
    console.error("No token found");
    return;
  }

  try {
    const response = await api.get("/validate", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    userDetails.value = response.data; // Storing user details
  } catch (error) {
    console.error("Failed to fetch user details", error);
  }
};

// const createPost = async () => {
//   error.value = null; // Reset error state before attempting to create a post

//   const formData = new FormData();
//   formData.append("text", form.text);
//   formData.append("image", form.image);

//   console.log(form.image);

//   try {
//     const response = await api.post("/create-post", formData, {
//       headers: {
//         "Content-Type": "multipart/form-data",
//         Authorization: `Bearer ${authStore.token}`,
//       },
//     });

//     console.log("Emitting postAdded with new post:", response.data.data);
//     emit("postAdded", response.data.data);
//     emit("showModal", false);
//     // emit("postAdded");
//     getPosts();
//   } catch (error) {
//     console.error("Failed to create post", error);
//     error.value = "Failed to create post. Please try again.";
//   }
// };



// function emitPostCreated() {
//   // useGeneral.postAdded = true;
//   postAdded.value = true;
//   console.log(" post emitted");
// }

function emitPostCreated() {
  // // Emit the new post details to the Pinia store
  // useGeneral.emitPostAdded(response.data.data);
  postAdded.value = true;
  setTimeout(() => {
    // Reset after all listeners have acted on it
    useGeneral.postAdded = false;
  }, 0);
  console.log(" post emitted");
}

// Ensure this runs when the component mounts
onMounted(async () => {
  await getUserDetails();
  // getPosts()
});
</script>

<template>
  <div
    id="CreatePostOverly"
    class="fixed z-50 top-0 left-0 w-full h-full bg-white bg-opacity-70"
  >
    <div class="grid h-screen place-items-center p-4">
      <div class="bg-white w-full mx-auto shadow-xl rounded-xl max-w-[600px]">
        <div class="flex items-center relative mx-1 my-3.5">
          <div class="font-extrabold w-full text-center text-[22px]">
            Create Post
          </div>
          <div
            @click="isPostOverlay = false"
            class="absolute right-3 p-1.5 rounded-full bg-gray-200 hover:bg-gray-300 cursor-pointer"
          >
            <Close :size="28" fillColor="#5E6771" />
          </div>
        </div>

        <div class="border-t border-t-gray-300">
          <div class="p-4">
            <div class="flex items-center">
              <img
                :src="userDetails?.image"
                alt=""
                class="rounded-full ml-1 min-w-[45px] max-h-[45px]"
              />
              <div class="ml-4">
                <div class="font-extrabold">{{ userDetails?.name }}</div>
                <div
                  class="flex items-center justify-between bg-gray-200 px-2 rounded-lg w-[100px] p-0.5"
                >
                  <Earth :size="18" />
                  <span class="font-bold pl-1.5 text-[13px]">Public</span>
                  <ChevronDown :size="18" class="pr-10 pl-1" />
                </div>
              </div>
            </div>
            <div class="max-h-[350px] overflow-auto">
              <textarea
                v-model="form.text"
                name=""
                id=""
                cols="30"
                rows="3"
                placeholder="What's on your mind?"
                class="w-full border-0 mt-4 focus:ring-0 text-[22px]"
              ></textarea>
              <div
                v-if="form.image"
                class="p-2 border border-gray-300 rounded-lg relative"
              >
                <Close
                  :size="24"
                  @click="clearImage()"
                  class="absolute bg-white p-0.5 m-2 right-2 rounded-full border cursor-pointer"
                  fillColor="#5E6771"
                />
                <p>Why isn't the image appearing here?</p>
                <img
                  class="rounded-lg mx-auto"
                  :src="imageDisplay"
                  alt="Image should be here!"
                />
              </div>
            </div>

            <div
              class="border-2 rounded-xl mt-4 shadow-sm flex items-center justify-between"
            >
              <div class="font-extrabold p-4">Add to your post</div>
              <div class="flex items-center">
                <label
                  for="image"
                  class="hover:bg-gray-200 rounded-full p-2 cursor-pointer"
                >
                  <Image :size="27" fillColor="#43BE62" />
                </label>
                <input
                  id="image"
                  type="file"
                  class="hidden"
                  @change="getUploadedImage($event)"
                />
                <button
                  class="hover:bg-gray-200 rounded-full p-2 cursor-pointer"
                >
                  <EmoticonOutline :size="27" fillColor="#F8B927" />
                </button>
                <button
                  class="hover:bg-gray-200 rounded-full p-2 cursor-pointer"
                >
                  <VideoImage :size="27" fillColor="#F12848" />
                </button>
                <button
                  class="hover:bg-gray-200 rounded-full p-2 cursor-pointer"
                >
                  <DotsHorizontal :size="27" fillColor="#050505" />
                </button>
              </div>
            </div>

            <div v-if="error">
              <div
                class="w-full bg-red-100 text-red-700 rounded-lg mt-3 text-center"
              >
                <div class="p-0.5">
                  {{ error }}
                </div>
              </div>
            </div>

            <button
              @click="createPost"
              class="w-full bg-blue-500 hover:bg-blue-600 text-white font-extrabold rounded-lg p-1.5 mt-3"
            >
              Post
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
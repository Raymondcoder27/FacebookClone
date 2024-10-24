<!-- <script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import PrimaryButton from "@/components/PrimaryButton.vue";
import GuestLayout from "@/Layouts/GuestLayout.vue";
import TextInput from "@/components/TextInput.vue";
// import TextInput from "@"
import api from "@/config/api";
import { useAuthStore } from "@/stores/auth";

import Camera from "vue-material-design-icons/Camera.vue";
import Pen from "vue-material-design-icons/Pen.vue";
// import { useRouter } from "vue-router";
import router from "@/router";
import {ref} from "vue"

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";

const authStore = useAuthStore()
const useGeneral = useGeneralStore();
const { isCropperModal, isImageDisplay } = storeToRefs(useGeneral);
const canResetPassword = true

const email = ref('')
const password = ref('')

defineProps({ posts: Object, user: Object });


const error = ref(null);

const submit = async () => {
    console.log("logging in...");

    error.value = null; // Reset error state before attempting login

    try {
        // Call the login action from the auth store
        await authStore.login(email.value, password.value);

        // On successful login, navigate to the home page
        router.push("/home");
    } catch (err) {
        // Set error message if login fails
        error.value = "Login failed. Please check your credentials.";
    }
};
</script> -->



<script setup>
import MainNavLayout from "@/Layouts/MainNavLayout.vue";
import CreatePostBox from "@/Layouts/CreatePostBox.vue";
import Post from "@/components/Post.vue";
import PrimaryButton from "@/components/PrimaryButton.vue";
import GuestLayout from "@/Layouts/GuestLayout.vue";
import TextInput from "@/components/TextInput.vue";
import api from "@/config/api";
import { useAuthStore } from "@/stores/auth";
import { ref } from "vue";
import router from "@/router";
import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";

const authStore = useAuthStore();
const useGeneral = useGeneralStore();
const { isCropperModal, isImageDisplay } = storeToRefs(useGeneral);
const canResetPassword = true;

const email = ref('');
const password = ref('');
const error = ref(null);

const submit = async () => {
    console.log("logging in...");

    error.value = null; // Reset error state before attempting login

    try {
        // Call the login action from the auth store
        await authStore.login(email.value, password.value);

        // Check local storage to see if the page has been reloaded
        const reloadStatus = localStorage.getItem('hasReloaded');

        if (!reloadStatus) {
            // Set local storage flag to prevent future reloads
            localStorage.setItem('hasReloaded', 'true');

            // Reload the page after a delay
            setTimeout(() => {
                window.location.reload(); // Reload after a delay
            }, 1000);
        } else {
            // On successful login, navigate to the home page
            router.push("/home");
        }
    } catch (err) {
        // Set error message if login fails
        error.value = "Login failed. Please check your credentials.";
    }
};

// const submit = async () => {
//   console.log("logging in...");

//   error.value = null; // Reset error state before attempting login

//   try {
//     // Call the login action from the auth store
//     await authStore.login(email.value, password.value);
    
//     // Since isAuthenticated is a computed property, it should automatically reflect changes
//     if (authStore.isAuthenticated) {
//       router.push("/home");
//     } else {
//       error.value = "Login failed. Please check your credentials.";
//     }
//   } catch (err) {
//     error.value = "Login failed. Please check your credentials.";
//   }
// };



</script>


<template>
  <Head title="Log in" />
  <!-- <img src="/public/icons/FacebookLogo.png" class="w-[200px]" alt=""> -->
  <div v-if="error" class="text-red-500">{{ error }}</div>


  <GuestLayout>
    <div
      class="w-full max-h-[50vh] bg-white text-center justify-center items-center mb-5"
    >
      <!-- <div class="w-full bg-red-300">
        <div class="mx-auto pb-1 pt-[56px] max-w-[1100px] bg-green-400"></div>
      </div> -->
      <!-- <img src="/public/icons/FacebookLogo.png" class="text-sm" alt=""> -->
      <form action="" @submit.prevent="submit" class="mb-5 text-black">
        <div>
          <TextInput
          v-model="email"
            id="email"
            type="email"
            class="mt-1 block w-full"
            required
            autofocus
            autocomplete="username"
            placeholder="Email"
          />
          <!-- Hello World!! -->

          <!-- <InputError class="mt-2" :message="form.errors.email" /> -->
        </div>

        <div class="mt-4">
          <TextInput
          v-model="password"
            id="password"
            type="password"
            class="mt-1 block w-full"
            required
            autofocus
            autocomplete="password"
            placeholder="Password"
          />
        </div>

        <div class="flex items-center justify-center pt-4">
          <PrimaryButton class="w-full text-sm" type="submit">
            <!-- <RouterLink
            type="submit"> -->
              Login
            <!-- </RouterLink> -->
           </PrimaryButton>
        </div>

        <div class="flex items-center justify-center my-5">
          <RouterLink
          to="/"
          v-if="canResetPassword"
          class="hover:underline font-bold text-blue-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Forgot Password?
          </RouterLink>
        </div>


        <div class="border-b border-b-gray-300"></div>
          <div class="flex items-center justify-center pt-6 pb-2">
            <RouterLink
            to="/register"
              class="px-5 py-3 text-white bg-[#37A621] text-[20px] font-bold rounded-lg text-sm"
            >
              Create a new account.
            </RouterLink>
          </div>
      </form>
    </div>
  </GuestLayout>
</template>
<script setup>
// import InputError from '@/Components/InputError.vue';
// import InputLabel from '@/Components/InputLabel.vue';
import PrimaryButton from "@/components/PrimaryButton.vue";
import TextInput from "@/components/TextInput.vue";
// import { useForm } from '@inertiajs/vue3';
import { ref } from "vue";

const passwordInput = ref(null);
const currentPasswordInput = ref(null);

// Use reactive object to hold form data
const form = ref({
  current_password: "",
  password: "",
  password_confirmation: "",
});

// const updatePass// Function to handle password update form submission
const updatePassword = async () => {
  try {
    // Send form data to server via axios (or fetch)
    const response = await axios.put("/password/update", form.value);

    // Reset form if the request is successful
    form.value = {
      current_password: "",
      password: "",
      password_confirmation: "",
    };

    // Optionally redirect to another page on success
    // router.push('/somewhere-else');
  } catch (error) {
    // Handle errors
    if (error.response?.data?.errors?.password) {
      passwordInput.value.focus();
    }
    if (error.response?.data?.errors?.current_password) {
      currentPasswordInput.value.focus();
    }
  }
};
</script>

<template>
  <section>
    <header>
      <h2 class="text-lg font-medium text-gray-900">Update Password</h2>
      <p class="mt-1 text-sm text-gray-600">
        Ensure your account is using a long, random password to stay secure.
      </p>
    </header>

    <form @submit.prevent="updatePassword" class="mt-6 space-y-6">
      <div>
        <InputLabel for="current_password" value="Current Password" />
        <TextInput
          id="current_password"
          ref="currentPasswordInput"
          v-model="form.current_password"
          type="password"
          class="mt-1 block w-full"
          autocomplete="current-password"
        />
        <InputError :message="form.errors?.current_password" class="mt-2" />
      </div>

      <div>
        <InputLabel for="password" value="New Password" />
        <TextInput
          id="password"
          ref="passwordInput"
          v-model="form.password"
          type="password"
          class="mt-1 block w-full"
          autocomplete="new-password"
        />
        <InputError :message="form.errors?.password" class="mt-2" />
      </div>

      <div>
        <InputLabel for="password_confirmation" value="Confirm Password" />
        <TextInput
          id="password_confirmation"
          v-model="form.password_confirmation"
          type="password"
          class="mt-1 block w-full"
          autocomplete="new-password"
        />
        <InputError
          :message="form.errors?.password_confirmation"
          class="mt-2"
        />
      </div>

      <div class="flex items-center gap-4">
        <PrimaryButton :disabled="form.processing">Save</PrimaryButton>

        <Transition
          enter-from-class="opacity-0"
          leave-to-class="opacity-0"
          class="transition ease-in-out"
        >
          <p v-if="form.recentlySuccessful" class="text-sm text-gray-600">
            Saved.
          </p>
        </Transition>
      </div>
    </form>
  </section>
</template>
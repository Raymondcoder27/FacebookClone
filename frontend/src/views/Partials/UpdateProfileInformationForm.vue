<script setup>
import InputError from '@/Components/InputError.vue';
import InputLabel from '@/Components/InputLabel.vue';
import PrimaryButton from '@/Components/PrimaryButton.vue';
import TextInput from '@/Components/TextInput.vue';
import { ref } from 'vue';
import axios from 'axios'; // Import axios for API requests

// Define props to receive user data and the email verification status
const props = defineProps({
    mustVerifyEmail: Boolean,
    status: String,
    user: Object, // We assume user data is passed as a prop
});

// Create a ref to hold form data
const form = ref({
    name: props.user.name || '',
    email: props.user.email || '',
});

// Handle form submission via axios
const updateProfile = async () => {
    try {
        // Send PATCH request to update profile information
        const response = await axios.patch('/profile/update', {
            name: form.value.name,
            email: form.value.email,
        });

        // Handle successful response, e.g., show success message
        console.log('Profile updated:', response.data);
    } catch (error) {
        // Handle errors, e.g., set error messages
        console.error('Failed to update profile:', error.response?.data || error);
    }
};
</script>

<template>
    <section>
        <header>
            <h2 class="text-lg font-medium text-gray-900">Profile Information</h2>
            <p class="mt-1 text-sm text-gray-600">
                Update your account's profile information and email address.
            </p>
        </header>

        <!-- Profile update form -->
        <form @submit.prevent="updateProfile" class="mt-6 space-y-6">
            <!-- Name Input -->
            <div>
                <InputLabel for="name" value="Name" />
                <TextInput
                    id="name"
                    type="text"
                    class="mt-1 block w-full"
                    v-model="form.name"
                    required
                    autofocus
                    autocomplete="name"
                />
                <InputError class="mt-2" :message="form.errors?.name" />
            </div>

            <!-- Email Input -->
            <div>
                <InputLabel for="email" value="Email" />
                <TextInput
                    id="email"
                    type="email"
                    class="mt-1 block w-full"
                    v-model="form.email"
                    required
                    autocomplete="email"
                />
                <InputError class="mt-2" :message="form.errors?.email" />
            </div>

            <!-- Email Verification Link -->
            <div v-if="props.mustVerifyEmail && !props.user.email_verified_at">
                <p class="text-sm mt-2 text-gray-800">
                    Your email address is unverified.
                    <button
                        @click="sendVerificationEmail"
                        class="underline text-sm text-gray-600 hover:text-gray-900 rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    >
                        Click here to re-send the verification email.
                    </button>
                </p>
                <div v-show="props.status === 'verification-link-sent'" class="mt-2 font-medium text-sm text-green-600">
                    A new verification link has been sent to your email address.
                </div>
            </div>

            <!-- Save Button and Success Message -->
            <div class="flex items-center gap-4">
                <PrimaryButton :disabled="form.processing">Save</PrimaryButton>
                <Transition enter-from-class="opacity-0" leave-to-class="opacity-0" class="transition ease-in-out">
                    <p v-if="form.recentlySuccessful" class="text-sm text-gray-600">Saved.</p>
                </Transition>
            </div>
        </form>
    </section>
</template>

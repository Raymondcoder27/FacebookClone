<script setup>
import DangerButton from '@/Components/DangerButton.vue';
import InputError from '@/Components/InputError.vue';
import InputLabel from '@/Components/InputLabel.vue';
import Modal from '@/Components/Modal.vue';
import SecondaryButton from '@/Components/SecondaryButton.vue';
import TextInput from '@/Components/TextInput.vue';
// import { useForm } from '@inertiajs/vue3';
import { nextTick, ref } from 'vue';


const confirmingUserDeletion = ref(false);
const passwordInput = ref(null);

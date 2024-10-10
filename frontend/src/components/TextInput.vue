<script setup>
    import { onMounted, ref } from 'vue';
</script>
<template>
    <input
    value="modelValue"
    @input="$emit('update:modelValue', $event.target.value)" type="text" class="border-gray-300 focus:border-indigo-500 focus:boder-indigo-500 rounded-md shadow-md">
</template>
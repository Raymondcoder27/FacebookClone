<script setup>
    import { onMounted, ref } from 'vue';
    defineProps(['modelValue', 'placeholder'])

    defineEmits('update:modelValue')

    const input = ref(null)

    onMounted(()=>{
        if(input.value.hasAttribute('Autofocus')){
            input.value.focus()
        }
    })

    defineExpose({focus: () => input.value.focus()})
</script>
<template>
    <input
    :placeholder="placeholder"
    :value="modelValue"
    ref="input"
    @input="$emit('update:modelValue', $event.target.value)" type="text" 
    class="border-2 border-gray-300 focus:border-indigo-500 focus:boder-indigo-500 rounded-md shadow-md p-3">
</template>
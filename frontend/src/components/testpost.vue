<script setup>
import { ref } from 'vue';

const posts = ref([]);

fetchPosts(); // Call your function to fetch posts here

async function fetchPosts() {
    try {
        const response = await fetch('/api/posts');
        const result = await response.json();
        posts.value = result.data; // Correctly assign data here
    } catch (error) {
        console.error(error);
    }
}
console.log("Fetched posts:", posts.value); // Shows data in array format

</script>
<template>
    <div v-for="post in posts" :key="post.id">
        <p>{{ post.text }}</p>
        <p>Posted by: {{ post.user.name }}</p>
    </div>
</template>

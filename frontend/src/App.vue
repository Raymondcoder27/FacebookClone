<script setup>
import { RouterLink, RouterView } from "vue-router";
import MainNavLayout from "./Layouts/MainNavLayout.vue";
import Posts from "./views/Posts.vue";
import Post from "./components/Post.vue";


//import router from "@/router";
import { ref, onMounted, toRefs } from "vue";
import api from "@/config/api";
import { useAuthStore } from "@/stores/auth";
const authStore = useAuthStore();

const posts = ref([]);

const getPosts = async () => {
  try {
    const response = await api.get("/posts");
    posts.value = response.data.data;
  } catch (error) {
    console.error("Failed to fetch posts", error);
  }
};


// defineProps({
//   // user: Object,
//   posts: Object,
//   // comments: Object,
// });

defineProps({
  posts: Array,
});

const handlePostDeleted = () => {
  getPosts()
}

const handlePostAdded = () => {
  getPosts()
}

// const { post, user, comments } = toRefs(props);

onMounted(() => {
  getPosts();
  // getUserDetails();
});
</script>

<template>
  <RouterView />
  <MainNavLayout>
    <!-- <CreatePostOverlay v-if="isPostOverlay" @showModal="isPostOverlay = false" @postAdded="handlePostAdded" /> -->
  <Posts  @postDeleted="handlePostDeleted"  />
   <User @postDeleted="handlePostDeleted" />
</MainNavLayout>


</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>

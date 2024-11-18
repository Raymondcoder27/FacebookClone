import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
// import { useStorage } from '@pinia/plugin-persistedstate'
import piniaPluginPersistedState from 'pinia-plugin-persistedstate'


export const useGeneralStore = defineStore('general', () => {
  const isPostOverlay = ref(false)
  const isCropperModal = ref(false)
  const isImageDisplay = ref(null)
  const postAdded = ref(null)

  // Action to emit a post
  const emitPostAdded = (newPost) => {
    postAdded.value = newPost; // Update with the new post data

    // Reset the value after emitting so it doesn't stay persistent
    setTimeout(() => {
      postAdded.value = null;
    }, 0);
  };

  return { isPostOverlay, isCropperModal, isImageDisplay, postAdded }
},
  // { persist: true })
  { 
    persist: {
     paths: ['isPostOverlay', 'isCropperModal', 'isImageDisplay'], // Persist only relevant states
     }
  })


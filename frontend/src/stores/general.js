import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
// import { useStorage } from '@pinia/plugin-persistedstate'
import piniaPluginPersistedState from 'pinia-plugin-persistedstate'


export const useGeneralStore = defineStore('general', () => {
  const isPostOverlay = ref(false)
  const isCropperModal = ref(false)
  const isImageDisplay = ref(null)
  const postAdded = ref(false)

  return { isPostOverlay, isCropperModal, isImageDisplay, postAdded }
},
  // { persist: true })
  { 
    persist: {
     paths: ['isPostOverlay', 'isCropperModal', 'isImageDisplay'], // Persist only relevant states
     }
  })


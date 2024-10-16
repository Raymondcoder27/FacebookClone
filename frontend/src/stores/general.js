import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
// import { useStorage } from '@pinia/plugin-persistedstate'
import piniaPluginPersistedState from 'pinia-plugin-persistedstate'


export const useGeneralStore = defineStore('general', () => {
  const isPostOverlay = ref(false)
  const isCropperModal = ref(false)
  const isImageDisplay = ref(null)


  return { isPostOverlay, isCropperModal, isImageDisplay }
},
  { persist: true })

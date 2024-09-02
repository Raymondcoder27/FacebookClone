import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useGeneralStore = defineStore('general', () => {
  const isPostOverlay = ref(false)
  const isCropperModal = ref(false)
  const isImageDisplay = ref(null)
  function increment() {
    count.value++
  }

  return { isPostOverlay, isCropperModal, isImageDisplay, increment }
})

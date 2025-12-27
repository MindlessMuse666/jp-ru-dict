// frontend\src\composables\useHotkeys.js

import { onMounted, onUnmounted } from 'vue'

export function useHotkeys({ onNewWord, onCloseModals }) {
    const handleKeyDown = (event) => {
        // Shift + N для нового слова
        if (event.shiftKey && event.key === 'N') {
            event.preventDefault()
            onNewWord?.()
        }

        // Shift + Escape для закрытия модальных окон
        if (event.shiftKey && event.key === 'Escape') {
            event.preventDefault()
            onCloseModals?.()
        }

        // Escape для закрытия модальных окон (без Shift)
        if (event.key === 'Escape') {
            onCloseModals?.()
        }
    }

    onMounted(() => {
        document.addEventListener('keydown', handleKeyDown)
    })

    onUnmounted(() => {
        document.removeEventListener('keydown', handleKeyDown)
    })
}
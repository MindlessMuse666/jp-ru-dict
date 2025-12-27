import { onMounted, onUnmounted } from 'vue'

export function useHotkeys() {
    const handleKeyDown = (event) => {
        // Ctrl/Cmd + K - фокус на поиск
        if ((event.ctrlKey || event.metaKey) && event.key === 'k') {
            event.preventDefault()
            const searchInput = document.querySelector('input[type="text"]')
            if (searchInput) {
                searchInput.focus()
            }
        }

        // Ctrl/Cmd + N - новое слово
        if ((event.ctrlKey || event.metaKey) && event.key === 'n') {
            event.preventDefault()
            return 'new-word'
        }

        // Escape - закрыть модалки
        if (event.key === 'Escape') {
            return 'close-modals'
        }

        // / (слеш) - фокус на поиск
        if (event.key === '/' && !event.ctrlKey && !event.metaKey) {
            const focusedElement = document.activeElement
            if (focusedElement.tagName !== 'INPUT' && focusedElement.tagName !== 'TEXTAREA') {
                event.preventDefault()
                const searchInput = document.querySelector('input[type="text"]')
                if (searchInput) {
                    searchInput.focus()
                }
            }
        }
    }

    const setupHotkeys = (callbacks = {}) => {
        const enhancedHandler = (event) => {
            const result = handleKeyDown(event)

            if (result === 'new-word' && callbacks.onNewWord) {
                callbacks.onNewWord()
            } else if (result === 'close-modals' && callbacks.onCloseModals) {
                callbacks.onCloseModals()
            }
        }

        onMounted(() => {
            window.addEventListener('keydown', enhancedHandler)
        })

        onUnmounted(() => {
            window.removeEventListener('keydown', enhancedHandler)
        })
    }

    return { setupHotkeys }
}
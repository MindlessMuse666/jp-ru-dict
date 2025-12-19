import { ref } from 'vue'

export function useToast() {
    const toasts = ref([])

    const showToast = (message, type = 'info', duration = 3000) => {
        const id = Date.now()
        const toast = { id, message, type }

        toasts.value.push(toast)

        // Автоматическое скрытие
        setTimeout(() => {
            removeToast(id)
        }, duration)

        return id
    }

    const showSuccess = (message, duration = 3000) => {
        return showToast(message, 'success', duration)
    }

    const showError = (message, duration = 5000) => {
        return showToast(message, 'error', duration)
    }

    const showWarning = (message, duration = 4000) => {
        return showToast(message, 'warning', duration)
    }

    const removeToast = (id) => {
        const index = toasts.value.findIndex(toast => toast.id === id)
        if (index !== -1) {
            toasts.value.splice(index, 1)
        }
    }

    return {
        toasts,
        showToast,
        showSuccess,
        showError,
        showWarning,
        removeToast
    }
}
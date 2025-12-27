import axios from 'axios'
import { useAuthStore } from '@/stores/auth'
import router from '@/router'

const api = axios.create({
    baseURL: import.meta.env.VITE_API_URL || '/api',
    headers: {
        'Content-Type': 'application/json',
    },
})

// Интерцептор для добавления токена к запросам
api.interceptors.request.use(
    (config) => {
        const authStore = useAuthStore()
        const token = authStore.token

        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }

        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

// Интерцептор для обработки ошибок
api.interceptors.response.use(
    (response) => response,
    (error) => {
        const authStore = useAuthStore()
        const { showError } = useToast()

        if (error.response?.status === 401) {
            authStore.clearAuth()
            if (!['Login', 'Register'].includes(router.currentRoute.value.name)) {
                router.push({ name: 'Login', query: { redirect: router.currentRoute.value.fullPath } })
            }
        } else if (error.response?.status === 429) {
            showError('Слишком много запросов. Пожалуйста, подождите.')
        } else if (error.code === 'ERR_NETWORK') {
            showError('Нет соединения с сервером. Проверьте подключение к интернету.')
        }

        return Promise.reject(error)
    }
)

export { api }
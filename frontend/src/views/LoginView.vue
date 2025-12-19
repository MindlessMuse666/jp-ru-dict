<template>
    <div class="max-w-md mx-auto">
        <div class="bg-white rounded-lg shadow-md p-8">
            <div class="text-center mb-8">
                <h1 class="text-2xl font-bold text-gray-800 mb-2">Вход в систему</h1>
                <p class="text-gray-600">Введите свои учетные данные</p>
            </div>

            <form @submit.prevent="handleLogin" class="space-y-6">
                <!-- Email -->
                <div>
                    <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
                        Email
                    </label>
                    <input id="email" v-model="form.email" type="email" required
                        class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent transition-colors"
                        placeholder="user@example.com" />
                </div>

                <!-- Пароль -->
                <div>
                    <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
                        Пароль
                    </label>
                    <input id="password" v-model="form.password" type="password" required
                        class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent transition-colors"
                        placeholder="••••••••" />
                </div>

                <!-- Ошибка -->
                <div v-if="error" class="p-4 bg-red-50 border border-red-200 rounded-lg">
                    <div class="flex items-center">
                        <svg class="w-5 h-5 text-red-500 mr-3" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd"
                                d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                                clip-rule="evenodd" />
                        </svg>
                        <span class="text-red-800">{{ error }}</span>
                    </div>
                </div>

                <!-- Кнопка отправки -->
                <button type="submit" :disabled="loading"
                    class="w-full py-3 px-4 bg-primary hover:bg-primary/90 text-white font-medium rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed">
                    <span v-if="loading">Вход...</span>
                    <span v-else>Войти</span>
                </button>

                <!-- Ссылка на регистрацию -->
                <div class="text-center pt-4 border-t">
                    <p class="text-gray-600">
                        Нет аккаунта?
                        <RouterLink :to="{ name: 'Register' }" class="text-primary hover:text-primary/80 font-medium">
                            Зарегистрироваться
                        </RouterLink>
                    </p>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useToast } from '@/composables/useToast'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const { showSuccess } = useToast()

const loading = ref(false)
const error = ref('')

const form = reactive({
    email: '',
    password: ''
})

const handleLogin = async () => {
    loading.value = true
    error.value = ''

    try {
        const result = await authStore.login(form.email, form.password)

        if (result.success) {
            showSuccess('Вы успешно вошли в систему')

            // Перенаправляем на страницу, с которой пришел пользователь, или на главную
            const redirect = route.query.redirect || '/'
            router.push(redirect)
        } else {
            error.value = result.error
        }
    } catch (err) {
        error.value = 'Произошла ошибка при входе'
    } finally {
        loading.value = false
    }
}
</script>
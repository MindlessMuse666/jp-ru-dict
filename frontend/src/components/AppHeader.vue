<template>
    <header class="bg-white shadow-sm border-b">
        <div class="container mx-auto px-4 py-3">
            <div class="flex items-center justify-between">
                <!-- Логотип и название -->
                <div class="flex items-center space-x-3">
                    <div class="text-2xl font-bold text-primary">
                        辞書
                    </div>
                    <div>
                        <h1 class="text-lg font-semibold text-gray-800">
                            Личный русско-японский словарь
                        </h1>
                        <p v-if="authStore.user" class="text-sm text-gray-600">
                            {{ authStore.user.name }}
                        </p>
                    </div>
                </div>

                <!-- Навигация -->
                <div class="flex items-center space-x-4">
                    <template v-if="authStore.isAuthenticated">
                        <button @click="handleLogout"
                            class="px-4 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors">
                            Выйти
                        </button>
                    </template>
                    <template v-else>
                        <RouterLink :to="{ name: 'Login' }"
                            class="px-4 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors">
                            Вход
                        </RouterLink>
                        <RouterLink :to="{ name: 'Register' }"
                            class="px-4 py-2 text-sm font-medium text-white bg-primary hover:bg-primary/90 rounded-md transition-colors">
                            Регистрация
                        </RouterLink>
                    </template>
                </div>
            </div>
        </div>
    </header>
</template>

<script setup>
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'
import { useToast } from '@/composables/useToast'

const authStore = useAuthStore()
const router = useRouter()
const { showSuccess } = useToast()

const handleLogout = async () => {
    await authStore.logout()
    showSuccess('Вы успешно вышли из системы')
    router.push({ name: 'Login' })
}
</script>
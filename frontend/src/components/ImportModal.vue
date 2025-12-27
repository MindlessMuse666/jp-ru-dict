<template>
    <Transition name="modal">
        <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center">
            <!-- Фон с затемнением -->
            <div class="fixed inset-0 bg-black bg-opacity-50" @click="close"></div>

            <!-- Модальное окно -->
            <div class="relative z-50 bg-white rounded-xl shadow-xl w-full max-w-2xl max-h-[90vh] overflow-hidden mx-4">
                <!-- Заголовок -->
                <div class="px-6 py-4 border-b bg-gray-50">
                    <div class="flex items-center justify-between">
                        <h3 class="text-lg font-semibold text-gray-800">Импорт слов из CSV</h3>
                        <button @click="close" class="text-gray-400 hover:text-gray-600 transition-colors p-1">
                            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M6 18L18 6M6 6l12 12" />
                            </svg>
                        </button>
                    </div>
                </div>

                <!-- Контент -->
                <div class="overflow-y-auto max-h-[calc(90vh-8rem)] p-6">
                    <!-- Формат CSV -->
                    <div class="mb-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
                        <h4 class="font-medium text-blue-800 mb-2">Формат CSV файла:</h4>
                        <p class="text-sm text-blue-700 mb-2">
                            <code class="bg-blue-100 px-2 py-1 rounded">jp;ru;on;kun;ex_jp;ex_ru;tags</code>
                        </p>
                        <p class="text-sm text-blue-700">
                            • Поля разделяются точкой с запятой<br>
                            • Массивы внутри полей разделяются запятой<br>
                            • Пример: <code
                                class="bg-blue-100 px-2 py-1 rounded">机;стол;キ,き;つくえ;机の上に本があります;На столе лежит книга;n5,мебель</code>
                        </p>
                    </div>

                    <!-- Поле для CSV -->
                    <div class="mb-6">
                        <label class="block text-sm font-medium text-gray-700 mb-2">
                            Содержимое CSV
                            <span class="text-red-500">*</span>
                        </label>
                        <textarea v-model="csvContent" rows="10"
                            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent transition-colors font-mono text-sm"
                            placeholder="Вставьте содержимое CSV файла сюда..."></textarea>
                        <div class="mt-2 text-sm text-gray-500">
                            Поддерживается до 1000 строк за один импорт
                        </div>
                    </div>

                    <!-- Ошибки -->
                    <div v-if="importResult && importResult.errors && importResult.errors.length > 0"
                        class="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg">
                        <h4 class="font-medium text-red-800 mb-2">
                            Ошибки импорта ({{ importResult.failed_count }}):
                        </h4>
                        <ul class="text-sm text-red-700 list-disc list-inside space-y-1">
                            <li v-for="(error, index) in importResult.errors" :key="index">
                                {{ error }}
                            </li>
                        </ul>
                    </div>

                    <!-- Результат импорта -->
                    <div v-if="importResult && importResult.imported_count > 0"
                        class="mb-6 p-4 bg-green-50 border border-green-200 rounded-lg">
                        <div class="flex items-center">
                            <svg class="w-5 h-5 text-green-500 mr-3" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd"
                                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                                    clip-rule="evenodd" />
                            </svg>
                            <div>
                                <h4 class="font-medium text-green-800">Импорт успешно завершен!</h4>
                                <p class="text-sm text-green-700">
                                    Импортировано слов: {{ importResult.imported_count }}
                                    <span v-if="importResult.failed_count > 0">
                                        , ошибок: {{ importResult.failed_count }}
                                    </span>
                                </p>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Футер -->
                <div class="px-6 py-4 border-t bg-gray-50">
                    <div class="flex justify-end space-x-3">
                        <button @click="close" type="button"
                            class="px-5 py-2.5 text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors font-medium">
                            Отмена
                        </button>
                        <button @click="handleImport" :disabled="!csvContent.trim() || isImporting" :class="[
                            'px-5 py-2.5 bg-primary hover:bg-primary/90 text-white font-medium rounded-lg transition-colors',
                            (!csvContent.trim() || isImporting) ? 'opacity-50 cursor-not-allowed' : ''
                        ]">
                            <span v-if="isImporting">Импорт...</span>
                            <span v-else>Импортировать</span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </Transition>
</template>

<script setup>
import { ref } from 'vue'
import { useToast } from '@/composables/useToast'
import { api } from '@/api/axios'

const props = defineProps({
    isOpen: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['close', 'imported'])

const { showSuccess, showError } = useToast()

const csvContent = ref('')
const isImporting = ref(false)
const importResult = ref(null)

const handleImport = async () => {
    if (!csvContent.value.trim()) return

    isImporting.value = true
    importResult.value = null

    try {
        const response = await api.post('/words/import', {
            content: csvContent.value
        })

        importResult.value = response.data.data

        if (importResult.value.imported_count > 0) {
            showSuccess(`Успешно импортировано ${importResult.value.imported_count} слов`)
            emit('imported')
        }

        if (importResult.value.failed_count > 0) {
            showError(`Не удалось импортировать ${importResult.value.failed_count} слов`)
        }

    } catch (error) {
        showError('Ошибка при импорте: ' + (error.response?.data?.error || 'неизвестная ошибка'))
    } finally {
        isImporting.value = false
    }
}

const close = () => {
    csvContent.value = ''
    importResult.value = null
    emit('close')
}
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
    transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}
</style>
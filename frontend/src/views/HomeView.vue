<template>
    <div class="max-w-6xl mx-auto">
        <!-- Панель управления -->
        <div class="mb-6">
            <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-4">
                <div>
                    <h2 class="text-2xl font-bold text-gray-800">Мой словарь</h2>
                    <p v-if="wordsStore.words.length" class="text-sm text-gray-600">
                        Всего слов: {{ wordsStore.words.length }}
                    </p>
                </div>

                <div class="flex items-center space-x-4">
                    <!-- Переключатель компактного вида -->
                    <div class="flex items-center">
                        <span class="text-sm text-gray-700 mr-2">Компактный вид:</span>
                        <button @click="isCompactView = !isCompactView" :class="[
                            'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                            isCompactView ? 'bg-primary' : 'bg-gray-200'
                        ]">
                            <span :class="[
                                'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                                isCompactView ? 'translate-x-6' : 'translate-x-1'
                            ]" />
                        </button>
                    </div>

                    <!-- Кнопка добавления -->
                    <button @click="showAddModal = true"
                        class="px-4 py-2 bg-primary hover:bg-primary/90 text-white font-medium rounded-md transition-colors flex items-center">
                        <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                        </svg>
                        Добавить слово
                    </button>
                </div>
            </div>

            <!-- Поисковая панель -->
            <SearchBar @search="handleSearch" @advanced-search="handleAdvancedSearch" />
        </div>

        <!-- Сообщение об ошибке -->
        <div v-if="wordsStore.error" class="mb-4 p-4 bg-red-50 border border-red-200 rounded-lg">
            <div class="flex items-center">
                <svg class="w-5 h-5 text-red-500 mr-3" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                        clip-rule="evenodd" />
                </svg>
                <span class="text-red-800">{{ wordsStore.error }}</span>
            </div>
        </div>

        <!-- Сообщение о пустом списке -->
        <div v-if="!wordsStore.loading && wordsStore.words.length === 0" class="text-center py-12">
            <div class="text-gray-400 mb-4">
                <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1"
                        d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.746 0 3.332.477 4.5 1.253v13C19.832 18.477 18.246 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                </svg>
            </div>
            <h3 class="text-lg font-medium text-gray-700 mb-2">Словарь пуст</h3>
            <p class="text-gray-500 mb-4">Добавьте первое слово, нажав кнопку "Добавить слово"</p>
        </div>

        <!-- Список слов -->
        <div v-else>
            <div class="space-y-4">
                <WordCard v-for="word in wordsStore.words" :key="word.id" :word="word" :is-compact-view="isCompactView"
                    @edit="handleEditWord(word)" />
            </div>

            <!-- Индикатор загрузки -->
            <LoadingSpinner v-if="wordsStore.loading && wordsStore.words.length > 0" />

            <!-- Кнопка "Загрузить еще" для десктопа -->
            <div v-if="!wordsStore.loading && wordsStore.hasMore" class="mt-6 text-center">
                <button @click="loadMoreWords"
                    class="px-6 py-3 bg-gray-100 hover:bg-gray-200 text-gray-800 font-medium rounded-lg transition-colors">
                    Загрузить еще
                </button>
            </div>

            <!-- Сообщение о конце списка -->
            <div v-if="!wordsStore.hasMore && wordsStore.words.length > 0" class="mt-6 pt-6 border-t text-center">
                <p class="text-gray-500">Вы просмотрели все слова в словаре</p>
            </div>
        </div>

        <!-- Модальное окно добавления/редактирования -->
        <WordFormModal :is-open="showAddModal || showEditModal" :word="selectedWord" @close="closeModal"
            @saved="handleWordSaved" />
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useWordsStore } from '@/stores/words'
import { useToast } from '@/composables/useToast'
import WordCard from '@/components/WordCard.vue'
import WordFormModal from '@/components/WordFormModal.vue'
import SearchBar from '@/components/SearchBar.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'

const wordsStore = useWordsStore()
const { showSuccess, showError } = useToast()

const isCompactView = ref(true)
const showAddModal = ref(false)
const showEditModal = ref(false)
const selectedWord = ref(null)

// Бесконечная пагинация при скролле
const handleScroll = () => {
    const scrollPosition = window.innerHeight + window.scrollY
    const pageHeight = document.documentElement.offsetHeight - 200

    if (scrollPosition >= pageHeight &&
        !wordsStore.loading &&
        wordsStore.hasMore) {
        loadMoreWords()
    }
}

const loadMoreWords = () => {
    wordsStore.fetchWords()
}

const handleSearch = async (query) => {
    if (query.trim()) {
        await wordsStore.searchWords({ q: query })
    } else {
        wordsStore.fetchWords(20, true)
    }
}

const handleAdvancedSearch = async (params) => {
    await wordsStore.searchWords(params)
}

const handleEditWord = (word) => {
    selectedWord.value = word
    showEditModal.value = true
}

const handleWordSaved = async (wordData) => {
    try {
        let result

        if (selectedWord.value) {
            // Редактирование существующего слова
            result = await wordsStore.updateWord(selectedWord.value.id, wordData)
        } else {
            // Создание нового слова
            result = await wordsStore.createWord(wordData)
        }

        if (result.success) {
            showSuccess(
                selectedWord.value
                    ? 'Слово успешно обновлено'
                    : 'Слово успешно добавлено'
            )
        } else {
            showError(result.error)
        }
    } catch (error) {
        showError('Ошибка при сохранении слова')
    }
}

const closeModal = () => {
    showAddModal.value = false
    showEditModal.value = false
    selectedWord.value = null
}

// Инициализация
onMounted(() => {
    if (wordsStore.words.length === 0) {
        wordsStore.fetchWords()
    }

    // Добавляем обработчик скролла для бесконечной пагинации
    window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll)
})
</script>
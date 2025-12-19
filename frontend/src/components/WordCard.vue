<template>
    <div class="bg-white rounded-lg shadow-md p-4 mb-4 hover:shadow-lg transition-shadow duration-200 relative group"
        :class="{ 'border-l-4 border-primary': !isCompactView }">
        <!-- Действия при наведении -->
        <div
            class="absolute right-3 top-3 opacity-0 group-hover:opacity-100 transition-opacity duration-200 flex space-x-2">
            <button @click="handleEdit"
                class="p-2 bg-blue-50 hover:bg-blue-100 text-blue-600 rounded-full transition-colors"
                title="Редактировать">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
            </button>
            <button @click="handleDelete"
                class="p-2 bg-red-50 hover:bg-red-100 text-red-600 rounded-full transition-colors" title="Удалить">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
            </button>
        </div>

        <!-- Основное содержимое -->
        <div class="pr-12">
            <!-- Японские написания -->
            <div class="mb-3">
                <h3 class="text-lg font-semibold text-gray-800 mb-1 flex items-center">
                    <span class="jp-text text-xl mr-2">{{ word.jp[0] }}</span>
                    <span v-if="word.jp[1]" class="text-sm text-gray-600">({{ word.jp[1] }})</span>
                </h3>

                <!-- Чтения -->
                <div v-if="(word.on && word.on.length > 0) || (word.kun && word.kun.length > 0)"
                    class="flex flex-wrap gap-2 mt-1">
                    <span v-for="reading in word.on || []" :key="`on-${reading}`"
                        class="px-2 py-1 text-xs font-medium rounded-full bg-onyomi text-gray-800">
                        音: {{ reading }}
                    </span>
                    <span v-for="reading in word.kun || []" :key="`kun-${reading}`"
                        class="px-2 py-1 text-xs font-medium rounded-full bg-kunyomi text-gray-800">
                        訓: {{ reading }}
                    </span>
                </div>
            </div>

            <!-- Русские переводы -->
            <div class="mb-3">
                <div class="flex flex-wrap gap-2">
                    <span v-for="translation in word.ru" :key="translation"
                        class="px-3 py-1 bg-gray-100 text-gray-800 rounded-full text-sm">
                        {{ translation }}
                    </span>
                </div>
            </div>

            <!-- Дополнительная информация (только в развернутом виде) -->
            <div v-if="!isCompactView">
                <!-- Примеры -->
                <div v-if="word.ex_jp && word.ex_jp.length > 0" class="mt-4 pt-4 border-t">
                    <h4 class="text-sm font-medium text-gray-700 mb-2">Примеры:</h4>
                    <div class="space-y-2">
                        <div v-for="(example, index) in word.ex_jp" :key="`ex-jp-${index}`" class="text-sm">
                            <p class="jp-text text-gray-800 mb-1">{{ example }}</p>
                            <p v-if="word.ex_ru && word.ex_ru[index]" class="text-gray-600 italic">
                                {{ word.ex_ru[index] }}
                            </p>
                        </div>
                    </div>
                </div>

                <!-- Теги -->
                <div v-if="word.tags && word.tags.length > 0" class="mt-3">
                    <div class="flex flex-wrap gap-1">
                        <span v-for="tag in word.tags" :key="tag"
                            class="px-2 py-1 text-xs font-medium rounded-full bg-primary/10 text-primary">
                            {{ tag }}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useToast } from '@/composables/useToast'
import { useWordsStore } from '@/stores/words'

defineProps({
    word: {
        type: Object,
        required: true
    },
    isCompactView: {
        type: Boolean,
        default: true
    }
})

const emit = defineEmits(['edit', 'delete'])

const { showSuccess, showError } = useToast()
const wordsStore = useWordsStore()

const handleEdit = () => {
    emit('edit')
}

const handleDelete = async () => {
    if (!confirm('Вы уверены, что хотите удалить это слово?')) return

    try {
        const result = await wordsStore.deleteWord(props.word.id)

        if (result.success) {
            showSuccess('Слово успешно удалено')
        } else {
            showError(result.error)
        }
    } catch (error) {
        showError('Ошибка при удалении слова')
    }
}
</script>
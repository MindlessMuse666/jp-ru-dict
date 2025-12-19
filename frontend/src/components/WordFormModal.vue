<template>
    <div v-if="isOpen" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4">
        <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-hidden">
            <!-- Заголовок -->
            <div class="px-6 py-4 border-b">
                <h2 class="text-lg font-semibold text-gray-800">
                    {{ isEditing ? 'Редактировать слово' : 'Добавить новое слово' }}
                </h2>
            </div>

            <!-- Форма -->
            <div class="p-6 overflow-y-auto max-h-[calc(90vh-8rem)]">
                <form @submit.prevent="handleSubmit">
                    <!-- Основные поля -->
                    <div class="space-y-4">
                        <!-- Японские написания (обязательное) -->
                        <div>
                            <label class="block text-sm font-medium text-gray-700 mb-2">
                                Японские написания *
                                <span class="text-xs text-gray-500 ml-1">(минимум 1)</span>
                            </label>
                            <div class="space-y-2">
                                <div v-for="(item, index) in form.jp" :key="`jp-${index}`" class="flex gap-2">
                                    <input v-model="form.jp[index]" type="text" placeholder="Написание на японском"
                                        class="flex-1 px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent"
                                        :class="{ 'border-red-300': errors.jp && errors.jp[index] }" />
                                    <button v-if="form.jp.length > 1" type="button"
                                        @click="removeArrayItem('jp', index)"
                                        class="px-3 py-2 text-red-600 hover:text-red-800">
                                        Удалить
                                    </button>
                                </div>
                            </div>
                            <button type="button" @click="addArrayItem('jp')"
                                class="mt-2 text-sm text-primary hover:text-primary/80">
                                + Добавить написание
                            </button>
                            <p v-if="errors.jp" class="mt-1 text-sm text-red-600">{{ errors.jp }}</p>
                        </div>

                        <!-- Русские переводы (обязательное) -->
                        <div>
                            <label class="block text-sm font-medium text-gray-700 mb-2">
                                Русские переводы *
                                <span class="text-xs text-gray-500 ml-1">(минимум 1)</span>
                            </label>
                            <div class="space-y-2">
                                <div v-for="(item, index) in form.ru" :key="`ru-${index}`" class="flex gap-2">
                                    <input v-model="form.ru[index]" type="text" placeholder="Перевод на русский"
                                        class="flex-1 px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent"
                                        :class="{ 'border-red-300': errors.ru && errors.ru[index] }" />
                                    <button v-if="form.ru.length > 1" type="button"
                                        @click="removeArrayItem('ru', index)"
                                        class="px-3 py-2 text-red-600 hover:text-red-800">
                                        Удалить
                                    </button>
                                </div>
                            </div>
                            <button type="button" @click="addArrayItem('ru')"
                                class="mt-2 text-sm text-primary hover:text-primary/80">
                                + Добавить перевод
                            </button>
                            <p v-if="errors.ru" class="mt-1 text-sm text-red-600">{{ errors.ru }}</p>
                        </div>

                        <!-- Онъёми -->
                        <div>
                            <label class="block text-sm font-medium text-gray-700 mb-2">
                                Онъёми чтения
                                <span class="text-xs text-gray-500 ml-1">(по одному на строку)</span>
                            </label>
                            <div class="space-y-2">
                                <div v-for="(item, index) in form.on" :key="`on-${index}`" class="flex gap-2">
                                    <input v-model="form.on[index]" type="text" placeholder="音読み"
                                        class="flex-1 px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent" />
                                    <button v-if="form.on.length > 1" type="button"
                                        @click="removeArrayItem('on', index)"
                                        class="px-3 py-2 text-red-600 hover:text-red-800">
                                        Удалить
                                    </button>
                                </div>
                            </div>
                            <button type="button" @click="addArrayItem('on')"
                                class="mt-2 text-sm text-primary hover:text-primary/80">
                                + Добавить онъёми
                            </button>
                        </div>

                        <!-- Кунъёми -->
                        <div>
                            <label class="block text-sm font-medium text-gray-700 mb-2">
                                Кунъёми чтения
                                <span class="text-xs text-gray-500 ml-1">(по одному на строку)</span>
                            </label>
                            <div class="space-y-2">
                                <div v-for="(item, index) in form.kun" :key="`kun-${index}`" class="flex gap-2">
                                    <input v-model="form.kun[index]" type="text" placeholder="訓読み"
                                        class="flex-1 px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent" />
                                    <button v-if="form.kun.length > 1" type="button"
                                        @click="removeArrayItem('kun', index)"
                                        class="px-3 py-2 text-red-600 hover:text-red-800">
                                        Удалить
                                    </button>
                                </div>
                            </div>
                            <button type="button" @click="addArrayItem('kun')"
                                class="mt-2 text-sm text-primary hover:text-primary/80">
                                + Добавить кунъёми
                            </button>
                        </div>

                        <!-- Дополнительные поля (скрыты по умолчанию) -->
                        <div v-if="showAdvanced" class="space-y-4 pt-4 border-t">
                            <!-- Примеры на японском -->
                            <div>
                                <label class="block text-sm font-medium text-gray-700 mb-2">
                                    Примеры на японском
                                    <span class="text-xs text-gray-500 ml-1">(максимум 3)</span>
                                </label>
                                <div class="space-y-2">
                                    <div v-for="(item, index) in form.ex_jp" :key="`ex_jp-${index}`" class="flex gap-2">
                                        <textarea v-model="form.ex_jp[index]"
                                            placeholder="Пример предложения на японском" rows="2"
                                            class="flex-1 px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent"></textarea>
                                        <button v-if="form.ex_jp.length > 1" type="button"
                                            @click="removeArrayItem('ex_jp', index)"
                                            class="px-3 py-2 text-red-600 hover:text-red-800">
                                            Удалить
                                        </button>
                                    </div>
                                </div>
                                <button v-if="form.ex_jp.length < 3" type="button" @click="addArrayItem('ex_jp')"
                                    class="mt-2 text-sm text-primary hover:text-primary/80">
                                    + Добавить пример
                                </button>
                            </div>

                            <!-- Примеры на русском -->
                            <div>
                                <label class="block text-sm font-medium text-gray-700 mb-2">
                                    Переводы примеров
                                    <span class="text-xs text-gray-500 ml-1">(максимум 3)</span>
                                </label>
                                <div class="space-y-2">
                                    <div v-for="(item, index) in form.ex_ru" :key="`ex_ru-${index}`" class="flex gap-2">
                                        <textarea v-model="form.ex_ru[index]" placeholder="Перевод примера на русский"
                                            rows="2"
                                            class="flex-1 px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent"></textarea>
                                        <button v-if="form.ex_ru.length > 1" type="button"
                                            @click="removeArrayItem('ex_ru', index)"
                                            class="px-3 py-2 text-red-600 hover:text-red-800">
                                            Удалить
                                        </button>
                                    </div>
                                </div>
                                <button v-if="form.ex_ru.length < 3" type="button" @click="addArrayItem('ex_ru')"
                                    class="mt-2 text-sm text-primary hover:text-primary/80">
                                    + Добавить перевод примера
                                </button>
                            </div>

                            <!-- Теги -->
                            <div>
                                <label class="block text-sm font-medium text-gray-700 mb-2">
                                    Теги
                                    <span class="text-xs text-gray-500 ml-1">(максимум 5, через запятую)</span>
                                </label>
                                <input v-model="tagsInput" type="text" placeholder="n5, дом, мебель"
                                    class="w-full px-3 py-2 border rounded-md focus:ring-2 focus:ring-primary focus:border-transparent"
                                    @input="updateTagsFromInput" />
                                <div v-if="form.tags.length > 0" class="flex flex-wrap gap-2 mt-2">
                                    <span v-for="(tag, index) in form.tags" :key="`tag-${index}`"
                                        class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-primary/10 text-primary">
                                        {{ tag }}
                                        <button type="button" @click="removeTag(index)"
                                            class="ml-2 text-primary/70 hover:text-primary">
                                            &times;
                                        </button>
                                    </span>
                                </div>
                            </div>
                        </div>

                        <!-- Кнопка показать/скрыть дополнительные поля -->
                        <div class="pt-2">
                            <button type="button" @click="showAdvanced = !showAdvanced"
                                class="text-sm text-primary hover:text-primary/80 flex items-center">
                                <span>{{ showAdvanced ? 'Скрыть' : 'Показать' }} дополнительные поля</span>
                                <svg class="w-4 h-4 ml-1 transition-transform" :class="{ 'rotate-180': showAdvanced }"
                                    fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M19 9l-7 7-7-7" />
                                </svg>
                            </button>
                        </div>
                    </div>

                    <!-- Кнопки формы -->
                    <div class="mt-6 pt-4 border-t flex justify-end space-x-3">
                        <button type="button" @click="close"
                            class="px-4 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors">
                            Отмена
                        </button>
                        <button type="submit" :disabled="isSubmitting"
                            class="px-4 py-2 text-sm font-medium text-white bg-primary hover:bg-primary/90 rounded-md transition-colors disabled:opacity-50 disabled:cursor-not-allowed">
                            <span v-if="isSubmitting">Сохранение...</span>
                            <span v-else>{{ isEditing ? 'Сохранить' : 'Добавить' }}</span>
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { useToast } from '@/composables/useToast'

const props = defineProps({
    isOpen: {
        type: Boolean,
        default: false
    },
    word: {
        type: Object,
        default: null
    }
})

const emit = defineEmits(['close', 'saved'])

const { showSuccess, showError } = useToast()

const isEditing = ref(false)
const isSubmitting = ref(false)
const showAdvanced = ref(false)
const tagsInput = ref('')

const form = reactive({
    jp: [''],
    ru: [''],
    on: [''],
    kun: [''],
    ex_jp: [''],
    ex_ru: [''],
    tags: []
})

const errors = reactive({})

// Инициализация формы при открытии
watch(() => props.isOpen, (isOpen) => {
    if (isOpen) {
        resetForm()
        if (props.word) {
            isEditing.value = true
            loadWordData(props.word)
        } else {
            isEditing.value = false
        }
    }
})

const loadWordData = (word) => {
    form.jp = [...word.jp]
    form.ru = [...word.ru]
    form.on = word.on && word.on.length > 0 ? [...word.on] : ['']
    form.kun = word.kun && word.kun.length > 0 ? [...word.kun] : ['']
    form.ex_jp = word.ex_jp && word.ex_jp.length > 0 ? [...word.ex_jp] : ['']
    form.ex_ru = word.ex_ru && word.ex_ru.length > 0 ? [...word.ex_ru] : ['']
    form.tags = word.tags && word.tags.length > 0 ? [...word.tags] : []

    // Обновляем поле ввода тегов
    tagsInput.value = form.tags.join(', ')
}

const resetForm = () => {
    form.jp = ['']
    form.ru = ['']
    form.on = ['']
    form.kun = ['']
    form.ex_jp = ['']
    form.ex_ru = ['']
    form.tags = []
    tagsInput.value = ''
    showAdvanced.value = false
    Object.keys(errors).forEach(key => delete errors[key])
}

const validateForm = () => {
    let isValid = true
    Object.keys(errors).forEach(key => delete errors[key])

    // Проверка японских написаний
    const validJp = form.jp.filter(item => item.trim() !== '')
    if (validJp.length === 0) {
        errors.jp = 'Необходимо указать хотя бы одно японское написание'
        isValid = false
    }

    // Проверка русских переводов
    const validRu = form.ru.filter(item => item.trim() !== '')
    if (validRu.length === 0) {
        errors.ru = 'Необходимо указать хотя бы один русский перевод'
        isValid = false
    }

    return isValid
}

const addArrayItem = (field) => {
    // Проверка лимитов для массивов
    if (field === 'ex_jp' && form.ex_jp.length >= 3) return
    if (field === 'ex_ru' && form.ex_ru.length >= 3) return
    if (field === 'tags' && form.tags.length >= 5) return

    if (field === 'jp' || field === 'ru' || field === 'on' || field === 'kun') {
        form[field].push('')
    } else if (field === 'ex_jp' || field === 'ex_ru') {
        form[field].push('')
    }
}

const removeArrayItem = (field, index) => {
    form[field].splice(index, 1)
    // Если массив пуст, добавляем пустой элемент для полей чтений
    if ((field === 'on' || field === 'kun') && form[field].length === 0) {
        form[field].push('')
    }
}

const updateTagsFromInput = () => {
    const tags = tagsInput.value
        .split(',')
        .map(tag => tag.trim())
        .filter(tag => tag !== '')
        .slice(0, 5) // Ограничение на 5 тегов

    form.tags = tags
}

const removeTag = (index) => {
    form.tags.splice(index, 1)
    tagsInput.value = form.tags.join(', ')
}

const handleSubmit = async () => {
    if (!validateForm()) return

    isSubmitting.value = true

    try {
        // Подготавливаем данные для отправки
        const data = {
            jp: form.jp.filter(item => item.trim() !== ''),
            ru: form.ru.filter(item => item.trim() !== ''),
            on: form.on[0] ? form.on.filter(item => item.trim() !== '') : undefined,
            kun: form.kun[0] ? form.kun.filter(item => item.trim() !== '') : undefined,
            ex_jp: form.ex_jp[0] ? form.ex_jp.filter(item => item.trim() !== '') : undefined,
            ex_ru: form.ex_ru[0] ? form.ex_ru.filter(item => item.trim() !== '') : undefined,
            tags: form.tags.length > 0 ? form.tags : undefined
        }

        // Отправляем данные через emit
        emit('saved', data)

        // Закрываем модальное окно после успешного сохранения
        close()

    } catch (error) {
        showError('Ошибка при сохранении слова')
    } finally {
        isSubmitting.value = false
    }
}

const close = () => {
    resetForm()
    emit('close')
}
</script>
<template>
    <div class="mb-6">
        <!-- Основной поиск -->
        <div class="relative mb-4">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
            </div>
            <input v-model="searchQuery" type="text" placeholder="Поиск по японскому или русскому..."
                class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent transition-colors"
                @input="handleSearchInput" />
        </div>

        <!-- Кнопка расширенного поиска -->
        <div class="mb-4">
            <button @click="showAdvanced = !showAdvanced"
                class="text-sm text-primary hover:text-primary/80 flex items-center">
                <span>{{ showAdvanced ? 'Скрыть' : 'Показать' }} расширенный поиск</span>
                <svg class="w-4 h-4 ml-1 transition-transform" :class="{ 'rotate-180': showAdvanced }" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
            </button>
        </div>

        <!-- Расширенный поиск -->
        <Transition enter-active-class="transition-all duration-300 ease-out"
            enter-from-class="transform opacity-0 -translate-y-4" enter-to-class="transform opacity-100 translate-y-0"
            leave-active-class="transition-all duration-200 ease-in"
            leave-from-class="transform opacity-100 translate-y-0" leave-to-class="transform opacity-0 -translate-y-4">
            <div v-if="showAdvanced" class="bg-gray-50 rounded-lg p-4 border">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <!-- Теги -->
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-2">
                            Теги
                            <span class="text-xs text-gray-500 ml-1">(через запятую)</span>
                        </label>
                        <input v-model="advancedFilters.tags" type="text" placeholder="n5, дом, мебель"
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-primary focus:border-transparent" />
                    </div>

                    <!-- Онъёми -->
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-2">
                            Онъёми
                            <span class="text-xs text-gray-500 ml-1">(через запятую)</span>
                        </label>
                        <input v-model="advancedFilters.on" type="text" placeholder="き, しょう"
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-primary focus:border-transparent" />
                    </div>

                    <!-- Кунъёми -->
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-2">
                            Кунъёми
                            <span class="text-xs text-gray-500 ml-1">(через запятую)</span>
                        </label>
                        <input v-model="advancedFilters.kun" type="text" placeholder="つくえ, た"
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-primary focus:border-transparent" />
                    </div>
                </div>

                <!-- Кнопки расширенного поиска -->
                <div class="mt-4 flex justify-end space-x-3">
                    <button @click="clearAdvancedFilters"
                        class="px-4 py-2 text-sm text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-md transition-colors">
                        Сбросить
                    </button>
                    <button @click="applyAdvancedSearch"
                        class="px-4 py-2 text-sm font-medium text-white bg-primary hover:bg-primary/90 rounded-md transition-colors">
                        Применить фильтры
                    </button>
                </div>
            </div>
        </Transition>
    </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { debounce } from 'lodash-es'

const emit = defineEmits(['search', 'advanced-search'])

const showAdvanced = ref(false)
const searchQuery = ref('')

const advancedFilters = reactive({
    tags: '',
    on: '',
    kun: ''
})

// Дебаунс для основного поиска
const debouncedSearch = debounce(() => {
    emit('search', searchQuery.value.trim())
}, 500)

const handleSearchInput = () => {
    debouncedSearch()
}

const applyAdvancedSearch = () => {
    const params = {}

    if (advancedFilters.tags.trim()) {
        params.tags = advancedFilters.tags.split(',').map(tag => tag.trim())
    }

    if (advancedFilters.on.trim()) {
        params.on = advancedFilters.on.split(',').map(item => item.trim())
    }

    if (advancedFilters.kun.trim()) {
        params.kun = advancedFilters.kun.split(',').map(item => item.trim())
    }

    emit('advanced-search', params)
}

const clearAdvancedFilters = () => {
    advancedFilters.tags = ''
    advancedFilters.on = ''
    advancedFilters.kun = ''
    emit('advanced-search', {})
}

// При изменении основного поиска сбрасываем расширенные фильтры
watch(searchQuery, (newValue) => {
    if (newValue.trim() === '') {
        clearAdvancedFilters()
    }
})
</script>
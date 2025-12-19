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

        // Отправляем данные через emit вместе с флагом редактирования
        emit('saved', { data, isEditing: isEditing.value, word: props.word })

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
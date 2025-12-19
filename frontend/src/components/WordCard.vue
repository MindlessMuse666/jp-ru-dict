<script setup>
import { useToast } from '@/composables/useToast'
import { useWordsStore } from '@/stores/words'

const props = defineProps({
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
    emit('edit', props.word)
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
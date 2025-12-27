import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@/api/axios'

export const useImportStore = defineStore('import', () => {
    const isImporting = ref(false)
    const importResult = ref(null)
    const error = ref(null)

    const importCSV = async (content) => {
        isImporting.value = true
        error.value = null
        importResult.value = null

        try {
            const response = await api.post('/words/import', { content })
            importResult.value = response.data.data
            return { success: true, data: importResult.value }
        } catch (err) {
            error.value = err.response?.data?.error || 'Ошибка при импорте'
            return { success: false, error: error.value }
        } finally {
            isImporting.value = false
        }
    }

    const clearImportResult = () => {
        importResult.value = null
        error.value = null
    }

    return {
        isImporting,
        importResult,
        error,
        importCSV,
        clearImportResult
    }
})
<template>
    <div class="fixed bottom-4 right-4 z-50 flex flex-col space-y-2">
        <TransitionGroup name="toast">
            <div v-for="toast in toasts" :key="toast.id" :class="[
                'px-4 py-3 rounded-lg shadow-lg max-w-md',
                toastClasses[toast.type]
            ]">
                <div class="flex items-center justify-between">
                    <span>{{ toast.message }}</span>
                    <button @click="removeToast(toast.id)" class="ml-4 text-lg hover:opacity-70 transition-opacity">
                        &times;
                    </button>
                </div>
            </div>
        </TransitionGroup>
    </div>
</template>

<script setup>
import { useToast } from '@/composables/useToast'

const { toasts, removeToast } = useToast()

const toastClasses = {
    success: 'bg-green-100 text-green-800 border border-green-200',
    error: 'bg-red-100 text-red-800 border border-red-200',
    warning: 'bg-yellow-100 text-yellow-800 border border-yellow-200',
    info: 'bg-blue-100 text-blue-800 border border-blue-200'
}
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
    transition: all 0.3s ease;
}

.toast-enter-from,
.toast-leave-to {
    opacity: 0;
    transform: translateX(100%);
}
</style>
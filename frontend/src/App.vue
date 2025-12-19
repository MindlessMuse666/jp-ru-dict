<template>
    <div class="min-h-screen bg-background">
        <AppHeader />
        <main class="container mx-auto px-4 py-6">
            <RouterView />
        </main>
        <ToastContainer />
    </div>
</template>

<script setup>
import AppHeader from '@/components/AppHeader.vue'
import ToastContainer from '@/components/ToastContainer.vue'
</script>
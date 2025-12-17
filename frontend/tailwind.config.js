/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                'onyomi': '#F9E79F',
                'kunyomi': '#A2C2E0',
                'primary': '#3B82F6',
                'secondary': '#10B981',
                'danger': '#EF4444',
                'background': '#F8F8F8',
            },
            fontFamily: {
                'sans': ['Roboto', 'sans-serif'],
                'jp': ['Hiragino Sans', 'Yu Gothic', 'Meiryo', 'sans-serif'],
            },
        },
    },
    plugins: [],
}
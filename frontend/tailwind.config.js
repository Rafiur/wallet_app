/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,jsx}'],
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#eef4ff',
          100: '#d9e6ff',
          200: '#b9d0ff',
          300: '#8bb0ff',
          400: '#5a86ff',
          500: '#3a63f7',
          600: '#2647db',
          700: '#1f38b0',
          800: '#1f3390',
          900: '#1e2f73',
        },
        good: '#0ca30c',
        critical: '#d03b3b',
      },
    },
  },
  plugins: [],
}

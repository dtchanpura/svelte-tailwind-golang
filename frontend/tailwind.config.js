module.exports = {
  purge: {
    preserveHtmlElements: false,
    content: [
      './src/*.svelte',
      './public/**/*.html',
    ]
  },
  darkMode: 'media', // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
}

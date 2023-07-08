/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin')

module.exports = {
  content: ['./templates/index.html', './templates/**/*.html'],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    // unsure if this is needed, but could be handy
    // https://www.crocodile.dev/blog/css-transitions-with-tailwind-and-htmx
    // plugin(function({ addVariant }) {
    //   addVariant('htmx-settling', ['&.htmx-settling', '.htmx-settling &'])
    //   addVariant('htmx-request', ['&.htmx-request', '.htmx-request &'])
    //   addVariant('htmx-swapping', ['&.htmx-swapping', '.htmx-swapping &'])
    //   addVariant('htmx-added', ['&.htmx-added', '.htmx-added &'])
    // }),
  ],
}


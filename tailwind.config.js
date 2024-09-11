/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{go,html,templ}",],
  theme: {
    extend: {},
  },
  plugins: ["prettier-plugin-tailwindcss"],
}
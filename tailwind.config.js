/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.{html,js,templ}",
    "./components/**/*.{html,js,templ}",
    "./handlers/**/*.{html,js,templ}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require("daisyui")
  ],
  daisyui: {
    themes: ["dark"],
    darkTheme: "dark",
    base: true,
    styled: true,
    utils: true,
  }
}
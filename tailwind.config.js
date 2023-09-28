/** @type {import('tailwindcss').Config} */
export default {
  content: ["./build/views/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        shopify: {
          "primary": "#008663",
          "primary-focus": "#006d51",
          "secondary": "#1a1a1a",
          "accent": "#00946e",
          "accent-focus": "#00785a",
          "accent-content": "#00dfa7",
          "neutral": "#2a2a2a",
          "base-100": "#fff",
          "info": "#bae6fd",
          "success": "#d9f99d",
          "warning": "#fef08a",        
          "error": "#f43f5e",
        },
      },
      "dark",
      "cupcake",
      "light",
      "dark",
      "cupcake",
      "bumblebee",
      "emerald",
      "corporate",
      "synthwave",
      "retro",
      "cyberpunk",
      "valentine",
      "halloween",
      "garden",
      "forest",
      "aqua",
      "lofi",
      "pastel",
      "fantasy",
      "wireframe",
      "black",
      "luxury",
      "dracula",
      "cmyk",
      "autumn",
      "business",
      "acid",
      "lemonade",
      "night",
      "coffee",
      "winter",
    ],
  },
};

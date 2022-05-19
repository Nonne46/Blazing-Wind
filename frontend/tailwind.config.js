module.exports = {
    content: ["./public/**/*.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
    theme: {
        extend: {},
        minWidth: {
            "1/2": "50%",
        },
    },
    plugins: [require("@tailwindcss/typography"), require("flowbite/plugin")],
};

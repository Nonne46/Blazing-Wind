<template>
    <nav
        class="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-800"
    >
        <div
            class="container flex flex-wrap justify-between items-center mx-auto"
        >
            <a href="/" class="flex items-center">
                <span
                    class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
                    >Razing Wind</span
                >
            </a>
        </div>
    </nav>
    <div class="container mx-auto">
        <router-view />
    </div>
</template>

<script>
export default {
    computed: {},
    methods: {},
};
</script>

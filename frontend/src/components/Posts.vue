<template>
    <div class="">
        <div class="flex justify-center mr-6 my-2">
            <input
                class="bg-purple-white sm:w-1/2 lg:w-1/3 shadow rounded border"
                type="text"
                v-model="search"
                placeholder="Анимус умер?"
            />
        </div>

        <div class="">
            <div
                v-for="post in posts"
                :key="post"
                class="border sm:w-full lg:w-2/3 mx-auto border-gray-400 rounded-lg md:p-4 bg-white sm:py-3 py-4 px-2 my-5"
            >
                <div class="pl-12 md:pl-10 xs:pl-10">
                    <h2
                        class="text-2xl font-bold mb-2 hover:text-blue-600 leading-7"
                    >
                        <a
                            :href="
                                'https://forum.ss13.ru/index.php?showtopic=' +
                                post.topic_id +
                                '&\#entry' +
                                post.id
                            "
                        >
                            {{ post.title }}
                        </a>
                    </h2>
                    <div class="pr-52">
                        <div class="flex items-center">
                            <div class="mr-2">
                                <a href="">
                                    <img
                                        class="rounded-full w-8"
                                        src=""
                                        alt=""
                                    />
                                </a>
                            </div>
                            <div>
                                <p>
                                    <a
                                        :href="
                                            'https://forum.ss13.ru/index.php?showuser=' +
                                            post.user_id
                                        "
                                        class="text text-gray-700 text-sm hover:text-black"
                                    >
                                        {{ post.username }}</a
                                    >
                                </p>
                                <a
                                    href=""
                                    class="text-xs text-gray-600 hover:text-black"
                                >
                                    <time :datetime="post.created_at"
                                        >{{ post.created_at }}
                                    </time>
                                </a>
                            </div>
                        </div>
                    </div>
                    <div
                        class="mb-1 leading-6 prose prose-slate prose-a:text-blue-600 hover:prose-a:text-blue-500 sm:prose-sm lg:prose-xl"
                    >
                        <Markdown :source="limitString(post.text)" />
                        <!-- {{ post.text }} -->
                    </div>
                </div>
            </div>
        </div>

        <!-- <div class="flex mb-4 items-center" v-for="post in posts" :key="post">
            <p class="w-full text-grey-darkest">{{ post.title }}</p>

            <div class="flex mb-4 items-center">
                <Markdown :source="limitString(post.text)" />
            </div>
        </div> -->
    </div>
</template>

<script>
import SearchService from "../services/search.service";
import Markdown from "vue3-markdown-it";
export default {
    // eslint-disable-next-line vue/multi-word-component-names
    name: "Search",
    components: {
        Markdown,
    },
    data() {
        return {
            search: "Анимус умер~",
            posts: [],
            awaitingSearch: false,
            last_topic_id: 0,
            select_search: "posts",
            breaks: true,
            html: true,
            linkify: true,
            typographer: true,
            quotes: ["«\xA0", "\xA0»", "‹\xA0", "\xA0›"],
        };
    },
    mounted() {
        this.getPosts(this.search);
    },
    watch: {
        search: function () {
            if (!this.awaitingSearch) {
                setTimeout(() => {
                    this.getPosts(this.search);
                    this.awaitingSearch = false;
                }, 1000); // 1 sec delay
            }
            this.awaitingSearch = true;
        },
    },
    methods: {
        limitString(str) {
            let length = 500;
            if (str.length > length) {
                return str.substring(0, length) + "...";
            }
            return str;
        },
        xzPosts(topic_id) {
            if (this.last_topic_id == topic_id) {
                this.last_topic_id = topic_id;
                return false;
            }
            this.last_topic_id = topic_id;
            return true;
        },
        getPosts(search) {
            SearchService.getPosts(search).then(
                (response) => {
                    this.posts = response.data.results;
                },
                (error) => {
                    this.content =
                        (error.response &&
                            error.response.data &&
                            error.response.data.message) ||
                        error.message ||
                        error.toString();
                }
            );
        },
    },
};
</script>

<style>
blockquote.p {
    margin: 15px;
}
pre {
    white-space: pre-wrap;
}
</style>

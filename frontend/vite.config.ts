import {
	sveltekit
} from "@sveltejs/kit/vite";
import autoprefixer from "autoprefixer";
import postcssImport from "postcss-import";
import tailwind from "tailwindcss";
import tailwindNesting from "tailwindcss/nesting";
import {
	defineConfig
} from "vite";
import tailwindConfig from "./tailwind.config.js";

export default defineConfig({
	plugins: [sveltekit()],
	css: {
		postcss: {
			plugins: [postcssImport(), tailwindNesting, tailwind(tailwindConfig), autoprefixer()],
		}
	},
});

import { defineConfig } from "umi";
import proxy from "./proxy";

export default defineConfig({
    nodeModulesTransform: {
        type: 'none',
    },
    proxy: proxy['dev'],
    fastRefresh: {},
   
})
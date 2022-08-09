const glob = require("glob");
const markdownIt = require("markdown-it");
const meta = require("markdown-it-meta");
const fs = require("fs");
const _ = require("lodash");

const sidebar = (directory, array) => {
    return array.map(i => {
        const children = _.sortBy(
            glob
                .sync(`./${directory}/${i[1]}/*.md`)
                .map(path => {
                    const md = new markdownIt();
                    const file = fs.readFileSync(path, "utf8");
                    md.use(meta);
                    md.render(file);
                    const order = md.meta.order;
                    return { path, order };
                })
                .filter(f => f.order !== false),
            ["order", "path"]
        )
            .map(f => f.path)
            .filter(f => !f.match("README"));

        return {
            title: i[0],
            children
        };
    });
};

module.exports = {
    base: "/docs/",
    plugins: [
        ['@vuepress/search', {
            searchMaxSuggestions: 10
        }]
    ],
    locales: {
        "/": {
            lang: "en-US",
            title: "Spartan-Cosmos Documents",
            description: "Spartan-Cosmos Documents",
        },
        "/zh/": {
            lang: "简体中文",
            title: "Spartan-Cosmos 文档",
            description: "Spartan-Cosmos 文档",
        }
    },
    themeConfig: {
        repo: "bianjie/spartan-cosmos",
        docsDir: "docs",
        editLinks: true,
        docsBranch: "master",
        editLinkText: 'Help us improve this page!',
        locales: {
            "/": {
                selectText: 'Languages',
                label: 'English',
                editLinkText: 'Help us improve this page!',
                nav: [
                    {
                        text: 'Back to Spartan-Cosmos',
                        link: ''
                    }
                ],
                sidebar: sidebar("", [
                    ["Getting Started", "get-started"],
                    ["Concepts", "concepts"],
                    ["Features", "features"],
                    ["Daemon", "daemon"],
                    ["CLI Client", "cli-client"],
                    ["Endpoints", "endpoints"],
                    ["Tools", "tools"],
                    ["Migration", "migration"],
                    ["Resources", "resources"]
                ])
            },
            "/zh/": {
                selectText: '选择语言',
                label: '简体中文',
                editLinkText: '帮助我们完善此文档',
                nav: [{
                    text: 'Spartan-Cosmos 官网',
                    link: ''
                }],
                sidebar: sidebar("", [
                    ["快速开始", "/zh/get-started"],
                    ["概念", "/zh/concepts"],
                    ["功能模块", "/zh/features"],
                    ["守护进程", "/zh/daemon"],
                    ["命令行客户端", "/zh/cli-client"],
                    ["服务端点", "/zh/endpoints"],
                    ["工具", "/zh/tools"],
                    ["迁移", "/zh/migration"],
                    ["资源", "/zh/resources"]
                ])
            }
        },
    }
};
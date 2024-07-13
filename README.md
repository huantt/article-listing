## About
Collect your latest articles from sources such as [dev.to](https://dev.to), and then update the `README.md`.

## Use GitHub Action to update your README

**Step 1:** In your repository, create a file named `README.md.template`.

**Step 2:** Write anything you want within the `README.md.template` file.

**Step 3:** Embed one of the following entities within your `README.md.template`:

- **Article listing:**
```shell
{{ template "article-list" .Articles }}
```
- **Article table:**
```shell
{{ template "article-table" .Articles }}
```

If you are familiar with Go templates, you have access to the `root` variable, which includes the following fields:

- `Articles`: An array of Article. You can view the Article struct definition in [model/article.go](model/article.go).
- `Time`: Updated Time
- `Author`: Author of articles

**Step 4**: Register Github Action
- Create a file `.github/workflows/update-articles.yml` in your repository.
```yml
name: "Cronjob"
on:
schedule:
- cron: '15 0 * * *'

jobs:
    update-articles:
        permissions: write-all
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v3
            - name: Generate README
              uses: huantt/article-listing@v1.1.0
              with:
                username: YOUR_USERNAME_ON_DEV_TO                
                template-file: 'README.md.template'
                out-file: 'README.md'
                limit: 5
            - name: Commit
              run: |
                if git diff --exit-code; then
                    echo "No changes to commit."
                    exit 0
                else
                    git config user.name github-actions
                    git config user.email github-actions@github.com
                    git add .
                    git commit -m "update"
                    git push origin main
                fi
```

**Step 5**: Commit your change, then Github actions will run as your specified cron to update Articles into your README.md file

## Below is my recent articles JackTT collected from dev.to
### Table


<table>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/comparing-limit-offset-and-cursor-pagination-1n81">Comparing Limit-Offset and Cursor Pagination</a>
                <div>Comparing Limit-Offset and Cursor Pagination   There are two popular methods for pagination...</div>
                <div><i>10/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/showing-more-article-info-on-dailydev-239b">Showing more Article info on Daily.dev</a>
                <div>Daily.dev is a very good extension that helps us aggregate news from several sources.  When...</div>
                <div><i>09/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/showing-more-article-info-on-dailydev-5351">Showing more Article info on Daily.dev</a>
                <div>Daily.dev is a very good extension that helps us aggregate news from several sources.  When...</div>
                <div><i>08/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/article-as-code-syncing-articles-between-devto-and-multiple-blogging-platforms-4a7c">[Article as Code] Syncing Articles Between Dev.to and Multiple Blogging...</a>
                <div>In the world of content creation, each platform offers unique advantages. Publishing articles on...</div>
                <div><i>08/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/why-are-nosql-databases-beeter-at-horizontal-scaling-compared-to-sql-databases-1hk2">Why are NoSQL Databases beeter at Horizontal Scaling Compared to...</a>
                <div>The ability of NoSQL databases to horizontally scale more effectively than SQL databases is rooted in...</div>
                <div><i>01/07/2024</i></div>
            </td>
        </tr>
</table>


### List

- [Comparing Limit-Offset and Cursor Pagination](https://dev.to/jacktt/comparing-limit-offset-and-cursor-pagination-1n81) - 10/07/2024
- [Showing more Article info on Daily.dev](https://dev.to/jacktt/showing-more-article-info-on-dailydev-239b) - 09/07/2024
- [Showing more Article info on Daily.dev](https://dev.to/jacktt/showing-more-article-info-on-dailydev-5351) - 08/07/2024
- [[Article as Code] Syncing Articles Between Dev.to and Multiple Blogging...](https://dev.to/jacktt/article-as-code-syncing-articles-between-devto-and-multiple-blogging-platforms-4a7c) - 08/07/2024
- [Why are NoSQL Databases beeter at Horizontal Scaling Compared to...](https://dev.to/jacktt/why-are-nosql-databases-beeter-at-horizontal-scaling-compared-to-sql-databases-1hk2) - 01/07/2024

*Updated at: 2024-07-13T12:40:30Z*
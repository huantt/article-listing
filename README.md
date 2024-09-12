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
                <a href="https://dev.to/jacktt/snowflake-schema-vs-star-schema-pros-cons-and-use-cases-34p9">Snowflake Schema vs. Star Schema: Pros, Cons, and Use Cases</a>
                <div>Star Schema            Structure:     Central Fact Table: Contains quantitative data for...</div>
                <div><i>12/09/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/postgres-isolation-levels-72h">[Postgres] Isolation levels</a>
                <div>In PostgreSQL, isolation levels determine how transaction integrity is maintained when multiple...</div>
                <div><i>12/09/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/database-multi-version-concurrency-control-47hk">[Database] Multi-Version Concurrency Control</a>
                <div>What is Multi-Version Concurrency Control   Multi-Version Concurrency Control (MVCC) is a...</div>
                <div><i>06/09/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/database-multi-version-concurrency-control-l9j">[Database] Multi-Version Concurrency Control</a>
                <div>What is Multi-Version Concurrency Control   Multi-Version Concurrency Control (MVCC) is a...</div>
                <div><i>06/09/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/database-multi-version-concurrency-control-1fi8">[Database] Multi-Version Concurrency Control</a>
                <div>What is Multi-Version Concurrency Control   Multi-Version Concurrency Control (MVCC) is a...</div>
                <div><i>05/09/2024</i></div>
            </td>
        </tr>
</table>


### List

- [Snowflake Schema vs. Star Schema: Pros, Cons, and Use Cases](https://dev.to/jacktt/snowflake-schema-vs-star-schema-pros-cons-and-use-cases-34p9) - 12/09/2024
- [[Postgres] Isolation levels](https://dev.to/jacktt/postgres-isolation-levels-72h) - 12/09/2024
- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-47hk) - 06/09/2024
- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-l9j) - 06/09/2024
- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-1fi8) - 05/09/2024

*Updated at: 2024-09-12T12:48:05Z*
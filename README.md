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
                <a href="https://dev.to/jacktt/database-multi-version-concurrency-control-47hk">[Database] Multi-Version Concurrency Control</a>
                <div>What is Multi-Version Concurrency Control   Multi-Version Concurrency Control (MVCC) is a...</div>
                <div><i>06/09/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/database-multi-version-concurrency-control-ke1">[Database] Multi-Version Concurrency Control</a>
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
        <tr>
            <td width="300px"><img src="https://media.dev.to/cdn-cgi/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2F3uxjgm8eik6pm9pps6o5.jpg" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/acid-in-postgres-6h8">ACID in Postgres</a>
                <div>ACID is a set of properties that ensure reliable transactions in a database system. It stands for...</div>
                <div><i>15/08/2024</i></div>
            </td>
        </tr>
</table>


### List

- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-47hk) - 06/09/2024
- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-ke1) - 06/09/2024
- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-l9j) - 06/09/2024
- [[Database] Multi-Version Concurrency Control](https://dev.to/jacktt/database-multi-version-concurrency-control-1fi8) - 05/09/2024
- [ACID in Postgres](https://dev.to/jacktt/acid-in-postgres-6h8) - 15/08/2024

*Updated at: 2024-09-10T12:49:40Z*
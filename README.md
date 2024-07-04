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
                <a href="https://dev.to/jacktt/graphql-types-26og">GraphQL Types</a>
                <div>Field declaration   By default, it&#39;s valid for any field in your schema to return null...</div>
                <div><i>03/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/graphql-fundamental-236k">Graphql Fundamental</a>
                <div>GraphQL provides a complete and understandable description of the API, including both &#34;API endpoints&#34;...</div>
                <div><i>03/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/snowflake-schema-vs-star-schema-pros-cons-and-use-cases-2701">Snowflake Schema vs. Star Schema: Pros, Cons, and Use Cases</a>
                <div>Star Schema            Structure:     Central Fact Table: Contains quantitative data for...</div>
                <div><i>11/06/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/is-jwt-safe-when-anyone-can-decode-plain-text-claims-2j7o">Is JWT Safe When Anyone Can Decode Plain Text Claims</a>
                <div>If I get a JWT and can decode the payload, how is it secure? Why couldn&#39;t I just grab the token out...</div>
                <div><i>06/06/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/understanding-the-select-for-update-sql-statement-900">Understanding the &#34;SELECT FOR UPDATE&#34; SQL Statement</a>
                <div>What is &#34;SELECT FOR UPDATE&#34;?   SELECT FOR UPDATE is a clause in SQL that is appended to a...</div>
                <div><i>06/06/2024</i></div>
            </td>
        </tr>
</table>


### List

- [GraphQL Types](https://dev.to/jacktt/graphql-types-26og) - 03/07/2024
- [Graphql Fundamental](https://dev.to/jacktt/graphql-fundamental-236k) - 03/07/2024
- [Snowflake Schema vs. Star Schema: Pros, Cons, and Use Cases](https://dev.to/jacktt/snowflake-schema-vs-star-schema-pros-cons-and-use-cases-2701) - 11/06/2024
- [Is JWT Safe When Anyone Can Decode Plain Text Claims](https://dev.to/jacktt/is-jwt-safe-when-anyone-can-decode-plain-text-claims-2j7o) - 06/06/2024
- [Understanding the &#34;SELECT FOR UPDATE&#34; SQL Statement](https://dev.to/jacktt/understanding-the-select-for-update-sql-statement-900) - 06/06/2024

*Updated at: 2024-07-04T18:30:30Z*
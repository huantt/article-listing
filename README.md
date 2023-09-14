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

If you are familiar with Go templates, you have access to the `root` variable, which includes the following fields:

- `Articles`: An array of Article. You can view the Article struct definition in [model/article.go](model/article.go).
- `Time`: Updated Time

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
              uses: huantt/article-listing@v1.0.0
              with:
                username: YOUR_USERNAME_ON_DEV_TO                
                template-file: 'README.md.template'
                out-file: 'README.md'
                limit: 5
            - name: Commit
              run: |
                git config user.name github-actions
                git config user.email github-actions@github.com
                git add .
                git commit -m "update articles"
                git push origin main
```

**Step 5**: Commit your change, then Github actions will run as your specified cron to update Articles into your README.md file

## Below is my recent articles Jack collected from dev.to
### Table


<table>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--9aLNv3pz--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/urlpgle748e1db4sw81v.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/creating-dynamic-readmemd-file-388o">Creating Dynamic README.md File</a>
                <div>This is my Github Profile. The specific thing here is that the weather is updated every 6 hours...</div>
                <div><i>09/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/search-goole-like-a-pro-cheat-sheet-555g">Search Goole Like a Pro [Cheat sheet]</a>
                <div>Before reading my article, let&#39;s try searching the following input:    inurl:/jacktt/ site:dev.to    ...</div>
                <div><i>30/08/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/go-build-in-advance-4o8n">Advanced Go Build Techniques</a>
                <div>Table of contents   Build options Which file will be included Build tags Build contraints           ...</div>
                <div><i>30/08/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--ZviKv8F5--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/b5m3kjdhd57zgk0xpcdr.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/load-private-module-in-golang-project-122l">Load Private Module in Golang Project</a>
                <div>Table of Contents    I. How Does go get Work? II. How to Load Private Modules III. Build...</div>
                <div><i>12/08/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/speed-up-your-query-in-postgres-48e3">Speed up your query in Postgres</a>
                <div>Table of contents    Indexing Use EXPLAIN ANALYZE Use UNION Instead of OR to Use Index Use...</div>
                <div><i>23/06/2023</i></div>
            </td>
        </tr>
</table>


### List

- [Creating Dynamic README.md File](https://dev.to/jacktt/creating-dynamic-readmemd-file-388o) - 09/09/2023
- [Search Goole Like a Pro [Cheat sheet]](https://dev.to/jacktt/search-goole-like-a-pro-cheat-sheet-555g) - 30/08/2023
- [Advanced Go Build Techniques](https://dev.to/jacktt/go-build-in-advance-4o8n) - 30/08/2023
- [Load Private Module in Golang Project](https://dev.to/jacktt/load-private-module-in-golang-project-122l) - 12/08/2023
- [Speed up your query in Postgres](https://dev.to/jacktt/speed-up-your-query-in-postgres-48e3) - 23/06/2023

*Updated at: 2023-09-14T01:20:31Z*
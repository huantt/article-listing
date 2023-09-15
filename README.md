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

## Below is my recent articles Jack collected from dev.to
### Table


<table>
        <tr>
            <td width="300px">
                <a href="https://dev.to/jacktt/update-your-devto-articles-into-your-github-profile-4dpi"><img src="data/images/default-thumbnail.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/jacktt/update-your-devto-articles-into-your-github-profile-4dpi">Update your dev.to articles into your Github profile</a>
                <div>In this article, I&#39;ll share you how to update your latest articles on dev.to into your Github...</div>
                <div><i>14/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px">
                <a href="https://dev.to/jacktt/creating-dynamic-readmemd-file-388o"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--9aLNv3pz--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/urlpgle748e1db4sw81v.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/jacktt/creating-dynamic-readmemd-file-388o">Creating Dynamic README.md File</a>
                <div>This is my Github Profile. The specific thing here is that the weather is updated every 6 hours...</div>
                <div><i>09/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px">
                <a href="https://dev.to/jacktt/search-goole-like-a-pro-cheat-sheet-555g"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--EyxkUB5z--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/zbxjrjp6rqnuvkymwy5q.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/jacktt/search-goole-like-a-pro-cheat-sheet-555g">Search Goole Like a Pro [Cheat sheet]</a>
                <div>Before reading my article, let&#39;s try searching the following input:    inurl:/jacktt/ site:dev.to    ...</div>
                <div><i>30/08/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px">
                <a href="https://dev.to/jacktt/go-build-in-advance-4o8n"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--dwEQQuvx--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/w3qkrh0jj4qkeumkm880.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/jacktt/go-build-in-advance-4o8n">Advanced Go Build Techniques</a>
                <div>Table of contents   Build options Which file will be included Build tags Build contraints           ...</div>
                <div><i>30/08/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px">
                <a href="https://dev.to/jacktt/load-private-module-in-golang-project-122l"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--ZviKv8F5--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/b5m3kjdhd57zgk0xpcdr.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/jacktt/load-private-module-in-golang-project-122l">Load Private Module in Golang Project</a>
                <div>Load Private Module in Golang Project           Table of Contents    I. How Does go get Work? II. How...</div>
                <div><i>12/08/2023</i></div>
            </td>
        </tr>
</table>


### List

- [Update your dev.to articles into your Github profile](https://dev.to/jacktt/update-your-devto-articles-into-your-github-profile-4dpi) - 14/09/2023
- [Creating Dynamic README.md File](https://dev.to/jacktt/creating-dynamic-readmemd-file-388o) - 09/09/2023
- [Search Goole Like a Pro [Cheat sheet]](https://dev.to/jacktt/search-goole-like-a-pro-cheat-sheet-555g) - 30/08/2023
- [Advanced Go Build Techniques](https://dev.to/jacktt/go-build-in-advance-4o8n) - 30/08/2023
- [Load Private Module in Golang Project](https://dev.to/jacktt/load-private-module-in-golang-project-122l) - 12/08/2023

*Updated at: 2023-09-15T12:48:46&#43;07:00*
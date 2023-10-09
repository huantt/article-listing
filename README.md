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
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/showing-more-article-info-on-dailydev-277g">Showing more Article info on Daily.dev</a>
                <div>Daily.dev is a very good extension that helps us aggregate news from several sources.  When...</div>
                <div><i>05/10/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--l7K58RVg--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/0zvq5w0swapgj83xs5dd.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/article-as-code-syncing-articles-between-devto-and-multiple-blogging-platforms-3hib">[Article as Code] Syncing Articles Between Dev.to and Multiple Blogging...</a>
                <div>In the world of content creation, each platform offers unique advantages. Publishing articles on...</div>
                <div><i>03/10/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--MYhfJtI0--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/kj1g9zo5j4zpjfwmmy01.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/20x-faster-golang-docker-builds-289n">20X Faster Golang Docker Builds</a>
                <div>According to the Go command documentation:  &#34;The go command caches build outputs for reuse in future...</div>
                <div><i>25/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--9LZDzMJ5--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/nyr1vfd7w2jwn96uowfj.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/why-i-like-writing-technical-blogs-11nm">Why I Like Writing Technical Blogs</a>
                <div>Requires me to delve further into the topic   It forces me to learn more deeply about the...</div>
                <div><i>20/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--UWs3rGfN--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/54buhtjhbrjqmjwvp6do.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/makding-dynamic-website-thumbnail-412k">Making Dynamic Website Thumbnail</a>
                <div>Look at this article&#39;s thumbnail; it&#39;s generated dynamically when this post receives new reactions or...</div>
                <div><i>21/09/2023</i></div>
            </td>
        </tr>
</table>


### List

- [Showing more Article info on Daily.dev](https://dev.to/jacktt/showing-more-article-info-on-dailydev-277g) - 05/10/2023
- [[Article as Code] Syncing Articles Between Dev.to and Multiple Blogging...](https://dev.to/jacktt/article-as-code-syncing-articles-between-devto-and-multiple-blogging-platforms-3hib) - 03/10/2023
- [20X Faster Golang Docker Builds](https://dev.to/jacktt/20x-faster-golang-docker-builds-289n) - 25/09/2023
- [Why I Like Writing Technical Blogs](https://dev.to/jacktt/why-i-like-writing-technical-blogs-11nm) - 20/09/2023
- [Making Dynamic Website Thumbnail](https://dev.to/jacktt/makding-dynamic-website-thumbnail-412k) - 21/09/2023

*Updated at: 2023-10-09T12:44:37Z*
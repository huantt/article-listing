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
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--oAEGp2JN--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/zjbfi8ti7p4jfg7zuo9f.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/20x-faster-golang-docker-builds-289n">20X Faster Golang Docker Builds</a>
                <div>According to the Go command documentation:  &#34;The go command caches build outputs for reuse in future...</div>
                <div><i>25/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--6TB374Kl--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/6obe2ztk6bby7n0spoup.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/why-i-like-writing-technical-blogs-11nm">Why I Like Writing Technical Blogs</a>
                <div>Requires me to delve further into the topic   It forces me to learn more deeply about the...</div>
                <div><i>20/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--z1iJvsWa--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ocu5mnfjxsnp4fol02t2.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/makding-dynamic-website-thumbnail-412k">Making Dynamic Website Thumbnail</a>
                <div>Look at this article&#39;s thumbnail; it&#39;s generated dynamically when this post receives new reactions or...</div>
                <div><i>21/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--d1Ji2l0_--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/zbqqmm8z78o20v367tl2.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/survey-automatic-table-of-contents-in-devto-articles-4m1g">[Survey] Automatic Table of Contents in Dev.to Articles</a>
                <div>If you&#39;re an avid reader on Dev.to, you might have come across articles that feature a table of...</div>
                <div><i>20/09/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--IYXpxeIK--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/xgku9dux2s3sohohiy5n.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/is-jwt-safe-when-anyone-can-decode-plain-text-claims-3anc">Is JWT Safe When Anyone Can Decode Plain Text Claims</a>
                <div>If I get a JWT and can decode the payload, how is it secure? Why couldn&#39;t I just grab the token out...</div>
                <div><i>15/09/2023</i></div>
            </td>
        </tr>
</table>


### List

- [20X Faster Golang Docker Builds](https://dev.to/jacktt/20x-faster-golang-docker-builds-289n) - 25/09/2023
- [Why I Like Writing Technical Blogs](https://dev.to/jacktt/why-i-like-writing-technical-blogs-11nm) - 20/09/2023
- [Making Dynamic Website Thumbnail](https://dev.to/jacktt/makding-dynamic-website-thumbnail-412k) - 21/09/2023
- [[Survey] Automatic Table of Contents in Dev.to Articles](https://dev.to/jacktt/survey-automatic-table-of-contents-in-devto-articles-4m1g) - 20/09/2023
- [Is JWT Safe When Anyone Can Decode Plain Text Claims](https://dev.to/jacktt/is-jwt-safe-when-anyone-can-decode-plain-text-claims-3anc) - 15/09/2023

*Updated at: 2023-09-26T12:42:31Z*
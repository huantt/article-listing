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
                <a href="https://dev.to/jacktt/understanding-the-weighted-random-algorithm-581p">Understanding the Weighted Random Algorithm</a>
                <div>Imagine you have a collection of items, and each item has a different &#34;weight,&#34; or probability of...</div>
                <div><i>24/10/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/migrate-redis-to-aws-elasticache-8bl">Migrate Redis to AWS ElastiCache</a>
                <div>How to Migrate Redis from One Server to Another   Starting from Redis version 5.0.0, Redis...</div>
                <div><i>19/10/2023</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--4VPFPpha--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/9qo025bd43wsi3w0odc7.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/writing-resume-as-code-why-not-iab">Writing Resume as Code - Why not?</a>
                <div>Why I Developed Resume as Code   I&#39;ve explored various CV builder platforms over the...</div>
                <div><i>12/10/2023</i></div>
            </td>
        </tr>
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
</table>


### List

- [Understanding the Weighted Random Algorithm](https://dev.to/jacktt/understanding-the-weighted-random-algorithm-581p) - 24/10/2023
- [Migrate Redis to AWS ElastiCache](https://dev.to/jacktt/migrate-redis-to-aws-elasticache-8bl) - 19/10/2023
- [Writing Resume as Code - Why not?](https://dev.to/jacktt/writing-resume-as-code-why-not-iab) - 12/10/2023
- [Showing more Article info on Daily.dev](https://dev.to/jacktt/showing-more-article-info-on-dailydev-277g) - 05/10/2023
- [[Article as Code] Syncing Articles Between Dev.to and Multiple Blogging...](https://dev.to/jacktt/article-as-code-syncing-articles-between-devto-and-multiple-blogging-platforms-3hib) - 03/10/2023

*Updated at: 2023-10-25T18:26:30Z*
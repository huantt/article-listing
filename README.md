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
                <a href="https://dev.to/jacktt/kubernetes-scheduler-3b68">Kubernetes Scheduler</a>
                <div>I. Concepts            Node labels   When create a node, we can mark some labels for it. On...</div>
                <div><i>31/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/kubernetes-scheduler-1173">Kubernetes Scheduler</a>
                <div>I. Concepts            Node labels   When create a node, we can mark some labels for it. On...</div>
                <div><i>31/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/kubernetes-scheduler-3opf">Kubernetes Scheduler</a>
                <div>I. Concepts            Node labels   When create a node, we can mark some labels for it. On...</div>
                <div><i>31/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/exploring-some-powerful-features-of-golang-36h0">Exploring some Powerful Features of Golang</a>
                <div>Table of content    Goroutines Channel Buffered Channel Defer Select Plugin Binary...</div>
                <div><i>31/07/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/explaining-ab-testing-algorithm-3k6b">Explaining A/B testing algorithm</a>
                <div>We are developing a crypto news aggregator Bitesapp.co. A few days ago, we implemented A/B testing...</div>
                <div><i>30/07/2024</i></div>
            </td>
        </tr>
</table>


### List

- [Kubernetes Scheduler](https://dev.to/jacktt/kubernetes-scheduler-3b68) - 31/07/2024
- [Kubernetes Scheduler](https://dev.to/jacktt/kubernetes-scheduler-1173) - 31/07/2024
- [Kubernetes Scheduler](https://dev.to/jacktt/kubernetes-scheduler-3opf) - 31/07/2024
- [Exploring some Powerful Features of Golang](https://dev.to/jacktt/exploring-some-powerful-features-of-golang-36h0) - 31/07/2024
- [Explaining A/B testing algorithm](https://dev.to/jacktt/explaining-ab-testing-algorithm-3k6b) - 30/07/2024

*Updated at: 2024-08-06T18:32:33Z*
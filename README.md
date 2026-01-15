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
            <td width="300px"><img src="https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fb7wglhgavqdxj92rzgpj.jpeg" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/scaling-read-tracking-with-redis-bitmaps-3aip">Scaling Read Tracking with Redis Bitmaps</a>
                <div>A friend recently came to me with a problem. They had designed a feature to track whether each user...</div>
                <div><i>16/09/2025</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/aws-lambda-with-go-how-to-build-deploy-and-invoke-1p0o">AWS Lambda with Go - How to Build, Deploy, and...</a>
                <div>Build, Deploy, and Invoke AWS Lambda Functions in Golang            1. Initialize Lambda...</div>
                <div><i>20/08/2025</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/tracing-error-strack-in-golang-234o">Tracing error strack in Golang</a>
                <div>Problem: No Stack Trace in Native Errors   Consider this Go snippet:    func function3()...</div>
                <div><i>23/05/2025</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fz4mwymfyiy1d2jgrweca.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/a-better-pkggodev-hip">A better pkg.go.dev</a>
                <div>About pkgo.dev   I have never been able to read a package&#39;s documentation on pkg.go.dev...</div>
                <div><i>16/02/2025</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2F3wrh7ita355fq7dzzto1.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/pod-topology-spread-constraints-2pd7">Pod Topology Spread Constraints</a>
                <div>Pod Topology Spread Constraints   Pod Topology Spread Constraints - a.k.a...</div>
                <div><i>12/01/2025</i></div>
            </td>
        </tr>
</table>


### List

- [Scaling Read Tracking with Redis Bitmaps](https://dev.to/jacktt/scaling-read-tracking-with-redis-bitmaps-3aip) - 16/09/2025
- [AWS Lambda with Go - How to Build, Deploy, and...](https://dev.to/jacktt/aws-lambda-with-go-how-to-build-deploy-and-invoke-1p0o) - 20/08/2025
- [Tracing error strack in Golang](https://dev.to/jacktt/tracing-error-strack-in-golang-234o) - 23/05/2025
- [A better pkg.go.dev](https://dev.to/jacktt/a-better-pkggodev-hip) - 16/02/2025
- [Pod Topology Spread Constraints](https://dev.to/jacktt/pod-topology-spread-constraints-2pd7) - 12/01/2025

*Updated at: 2026-01-15T06:51:05Z*
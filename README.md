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
                <a href="https://dev.to/jacktt/kafka-dump-backup-restore-stream-40ai">Kafka dump: backup, restore, stream,...</a>
                <div>Source code   Kafka dump           Kafka data backup   Kafka dump is a tool to back up and...</div>
                <div><i>24/01/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/kafka-dump-backup-restore-stream-22hf">Kafka dump: backup, restore, stream,...</a>
                <div>Source code   Kafka dump           Kafka data backup   Kafka dump is a tool to back up and...</div>
                <div><i>24/01/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/saa-aws-storage-types-1noo">[SAA] AWS EBS, EFS and EC2 insntace storage</a>
                <div>EBS (a.k.a. Elastic Block Store)    They are bound to a specific AZ (in specific region) You...</div>
                <div><i>24/01/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/kubernetes-scheduler-129i">Kubernetes Scheduler</a>
                <div>I. Concepts            Node labels   When create a node, we can mark some labels for it. On...</div>
                <div><i>03/01/2024</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px"><img src="data/images/default-thumbnail.png" alt="thumbnail"></td>
            <td>
                <a href="https://dev.to/jacktt/multiple-git-configs-profiles-on-one-computer-2ik">Multiple git configs (profiles) on one computer</a>
                <div>How to let git know which profile should be used in specific folders?   Imagine that you’re...</div>
                <div><i>26/10/2023</i></div>
            </td>
        </tr>
</table>


### List

- [Kafka dump: backup, restore, stream,...](https://dev.to/jacktt/kafka-dump-backup-restore-stream-40ai) - 24/01/2024
- [Kafka dump: backup, restore, stream,...](https://dev.to/jacktt/kafka-dump-backup-restore-stream-22hf) - 24/01/2024
- [[SAA] AWS EBS, EFS and EC2 insntace storage](https://dev.to/jacktt/saa-aws-storage-types-1noo) - 24/01/2024
- [Kubernetes Scheduler](https://dev.to/jacktt/kubernetes-scheduler-129i) - 03/01/2024
- [Multiple git configs (profiles) on one computer](https://dev.to/jacktt/multiple-git-configs-profiles-on-one-computer-2ik) - 26/10/2023

*Updated at: 2024-04-26T18:27:37Z*
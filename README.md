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

## Below is my recent articles Sloan the DEV Moderator collected from dev.to


- [Welcome Thread - v242](https://dev.to/devteam/welcome-thread-v242-1b40) - 01/09/2023
- [Creating a Quantum Computing Portfolio Segment Using Variational Quantum Eigensolver](https://dev.to/aws/creating-a-quantum-computing-portfolio-segment-using-variational-quantum-eigensolver-ioc) - 13/09/2023
- [Silver Screen Code: What&#39;s Your Job&#39;s Movie Genre?](https://dev.to/codenewbieteam/silver-screen-code-whats-your-jobs-movie-genre-49l7) - 11/09/2023
- [How Do You Create Inclusive Workplaces?](https://dev.to/devteam/how-do-you-create-inclusive-workplaces-d5m) - 05/09/2023
- [What is your Why?](https://dev.to/acoh3n/what-is-your-why-j9b) - 13/09/2023

*Updated at: 2023-09-13T13:12:17Z*
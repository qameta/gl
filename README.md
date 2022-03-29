## Qameta GL (Gitlab CLI)

cli util designed to run Gitlab Pipelines with passing ENV variables (missing in glab cli)

Examples

```shell
## To set up
gl auth login --token=<your_gitlab_token> --project-id=<your_project_id>
## To run
gl pipeline run --env VERSION=1.1.2,ARCH=amd64 // DO NOT use spaces after commas
```
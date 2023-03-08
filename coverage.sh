commentJson='{"body":"Hello world!"}'

# PRのコメント一覧から既存のcommentを取得
# comment_id=$(curl \
#   -H "Accept: application/vnd.github+json" \
#   -H "Authorization: Bearer ${{ secrets.GHEC_MACHINE_ACCESS_TOKEN }}" \
#   https://api.github.com/repos/atsutama2/go_github_actions_test/issues/${BITRISE_PULL_REQUEST}/comments?per_page=100 \
#   | jq 'first(.[] | select(.body | startswith("## Code coverage")) | .id)')

#   curl -X POST \
#   -H "Accept: application/vnd.github.v3+json" \
#   -H "Authorization: Bearer <YOUR-TOKEN>" \
#   -d '{"body": "Hello from GitHub Actions!"}' \
#   https://api.github.com/repos/:owner/:repo/issues/:issue_number/comments

#   curl -L \
#   -H "Accept: application/vnd.github+json" \
#   -H "Authorization: Bearer <YOUR-TOKEN>"\
#   -H "X-GitHub-Api-Version: 2022-11-28" \
#   https://api.github.com/repos/OWNER/REPO/issues/comments

# Github PRにcomment
# api: https://docs.github.com/ja/rest/issues/comments#create-an-issue-comment
# curl \
#   -X POST \
#   -H "Accept: application/vnd.github.v3+json" \
#   -H "Authorization: token ${GHE_AUTH_TOKEN}" \
#   https://api.github.com/repos/atsutama2/go_github_actions_test/issues/${PULL_REQEST}/comments \
#   -d "$commentJson"
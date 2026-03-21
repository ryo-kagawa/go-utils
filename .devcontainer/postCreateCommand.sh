go_versions=$(curl -sSL 'https://go.dev/dl/?mode=json&include=all' | jq -r 'map(select(.stable and (.version | split(".")[1] // "0" | tonumber) >= 22)) | group_by(.version | split(".")[0:2] | join(".")) | map(sort_by(.version | split(".")[2] // "0" | tonumber) | last) | .[].version')
for version in $go_versions; do
  go install golang.org/dl/${version}@latest
  ${version} download
done

echo 'installed go list'

for cmd in $(ls $(go env GOPATH)/bin/go1.*); do
  $cmd version;
done

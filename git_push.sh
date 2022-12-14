go build

git add .
git commit -m "Include go original database sql"

git tag -a v0.0.14 $(git log --format="%H" -n 1) -m "database sql"

git push

git push origin v0.0.14
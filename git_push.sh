go build

git add .
git commit -m "Fix"

git tag -a v0.0.8 $(git log --format="%H" -n 1) -m "Scrollable cursor"

git push

git push origin v0.0.8
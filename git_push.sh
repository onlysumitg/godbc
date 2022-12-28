go build 

git add .
git commit -m "Column label"

git tag -a v0.0.27 $(git log --format="%H" -n 1) -m "Column type"

git push

git push origin v0.0.27
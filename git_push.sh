go build 

now=$(date)
git add .
git commit -m "$now"

git tag -a v0.0.39 $(git log --format="%H" -n 1) -m "$now"

git push

git push origin v0.0.39

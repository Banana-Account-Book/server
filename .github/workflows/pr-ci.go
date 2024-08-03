name: pr-ci
on:
  pull_request:
    types: [opened, synchronize, reopened]
jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        node-version: [20.x]
    steps:
      - name: Checkout
        uses: actions/checkout@v4
	  - name: Set up Go
		uses: actions/setup-go@v4
		with:
		  go-version: '1.22'
	  - name: Build
		run: go build -v
  sonarcloud:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: SonarCloud Scan
        uses: SonarSource/sonarcloud-github-action@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}

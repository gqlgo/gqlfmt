# gqlfmt

`gqlfmt` formats your GraphQL files.

## Install

```
go install github.com/gqlgo/gqlfmt@latest
```

## Usage

- Help
```
$ gqlfmt --help
usage: gqlfmt [flags] [path ...]
  -d    display diffs instead of rewriting files
  -l    list files whose formatting differs from gqlfmt's
  -v    verbose logging
  -w    write result to (source) file instead of stdout
```

- to stdout
```sh
$ cat *.graphql # display unformatted files
```

```graphql
query A1 {
     id
   test
}

query A2 {
      id
      test
}
query B1 {
     id
 hello
}

query B2 {
      id



      test
}
query C1 {
        id
}
```

```bash
$ gqlfmt *.graphql
```

```graphql
query A1 {
        id
        test
}
query A2 {
        id
        test
}
query B1 {
        id
        hello
}
query B2 {
        id
        test
}
query C1 {
        id
}
```

- Over write GraphQL files
```sh
$ gqlfmt -w *.graphql
```

- List files whose formatting differs from gqlfmt's
```sh
$ gqlfmt -l *.graphql
a.graphql
b.graphql
```

- Display diffs
```sh
$ gqlfmt -d *.graphql
```

```diff
diff -u a.graphql.orig a.graphql
--- a.graphql.orig      2021-04-06 14:26:26.000000000 +0900
+++ a.graphql   2021-04-06 14:26:26.000000000 +0900
@@ -1,9 +1,8 @@
 query A1 {
-     id
-   test
+       id
+       test
 }
-
 query A2 {
-      id
-      test
+       id
+       test
 }
diff -u b.graphql.orig b.graphql
--- b.graphql.orig      2021-04-06 14:26:26.000000000 +0900
+++ b.graphql   2021-04-06 14:26:26.000000000 +0900
@@ -1,12 +1,8 @@
 query B1 {
-     id
- hello
+       id
+       hello
 }
-
 query B2 {
-      id
-
-
-
-      test
+       id
+       test
 }
```

- from stdin

```sh
$ cat a.graphql | gqlfmt -d
```

```diff
diff -u stdin.go.orig stdin.go
--- stdin.go.orig       2021-04-06 15:03:59.000000000 +0900
+++ stdin.go    2021-04-06 15:03:59.000000000 +0900
@@ -1,9 +1,8 @@
 query A1 {
-     id
-   test
+       id
+       test
 }
-
 query A2 {
-      id
-      test
+       id
+       test
 }
```

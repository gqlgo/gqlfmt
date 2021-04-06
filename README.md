# gqlfmt

`gqlfmt` formats your GraphQL files.

## CAUTION!!

The current version is an alpha version.
gqlfmt has two issues. https://github.com/gqlgo/gqlfmt/issues/1 and https://github.com/gqlgo/gqlfmt/issues/2

## Install

```
go install github.com/gqlgo/gqlfmt
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

- Over write GraphQL files
```
$ glqfmt -w *.graphql
```

- List files whose formatting differs from gqlfmt's
```
$ glqfmt -l *.graphql
a.graphql
b.graphql
```

- Display diffs
```
$ glqfmt -d *.graphql
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


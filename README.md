# go-utils

個人的ライブラリとして作成したもの  
各ディレクトリ名が実際のライブラリとなっている

## 開発方針

- CGOを使用しない
- 比較演算子は次の物のみを使用する
  - `==`
  - `!=`
  - `<`
  - `<=`

## テスト対象

|GOOS/GOARCH|OS Version|Go Version|
|-|-|-|
|darwin/amd64|14, 15, 26|1.22 or later|
|darwin/arm64|14, 15, 26|1.22 or later|
|dragonfly/amd64|6.4.2|1.22 or later|
|freebsd/amd64|12.2, 12.4, 13.0, 13.1, 13.2, 13.3, 13.4, 13.5, 14.0, 14.1, 14.2, 14.3, 14.4, 15.0, 15.1|1.22 or later|
|freebsd/arm64|12.4, 13.0, 13.1, 13.2, 13.3, 13.4, 13.5, 14.0, 14.1, 14.2, 14.3, 14.4, 15.0, 15.1|1.22 or later|
|linux/386(CGO_ENABLED=0)||1.22 or later|
|linux/386(GNU C Library)||1.22 or later|
|linux/386(musl)||1.22 or later|
|linux/arm v5(CGO_ENABLED=0)||1.22 or later|
|linux/arm v5(GNU C Library)||1.22 or later|
|linux/arm v6(CGO_ENABLED=0)||1.22 or later|
|linux/arm v6(musl)||1.22 or later|
|linux/arm v7(CGO_ENABLED=0)||1.22 or later|
|linux/arm v7(GNU C Library)||1.22 or later|
|linux/arm v7(musl)||1.22 or later|
|linux/arm64(CGO_ENABLED=0)||1.22 or later|
|linux/arm64(GNU C Library)||1.22 or later|
|linux/arm64(musl)||1.22 or later|
|linux/mips64le(CGO_ENABLED=0)||1.23 or later|
|linux/mips64le(GNU C Library)||1.23 or later|
|linux/ppc64le(CGO_ENABLED=0)||1.22 or later|
|linux/ppc64le(GNU C Library)||1.22 or later|
|linux/riscv64(CGO_ENABLED=0)||1.22 or later|
|linux/riscv64(GNU C Library)||1.22 or later|
|linux/riscv64(musl)||1.22 or later|
|linux/s390x(CGO_ENABLED=0)||1.22 or later|
|linux/s390x(GNU C Library)||1.22 or later|
|linux/s390x(musl)||1.22 or later|
|netbsd/amd64|9.2, 9.3, 9.4, 10.0, 10.1|1.22 or later|
|netbsd/arm64|10.0, 10.1|1.22 or later|
|openbsd/amd64|6.8, 6.9, 7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8, 7.9|1.22 or later|
|openbsd/arm64|6.9, 7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8, 7.9|1.22 or later|
|windows/386|Windows Server 2022, Windows Server 2025|1.22 or later|
|windows/amd64|Windows Server 2022, Windows Server 2025|1.22 or later|
|windows/arm64|Windows 11|1.22 or later|

# go-github-release

Get latest tag/assets of GitHub Releases

# Usage

```go
import "github.com/skaji/go-github-release"

release := &github.Release{
  Owner:      "skaji",
  Repository: "relocatable-perl",
}

tag, _ := release.GetLatestTag()
// 5.36.0.1

assets, _ := release.GetLatestAssets()
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-darwin-amd64.tar.gz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-darwin-amd64.tar.xz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-darwin-arm64.tar.gz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-darwin-arm64.tar.xz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-linux-amd64.tar.gz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-linux-amd64.tar.xz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-linux-arm64.tar.gz
// https://github.com/skaji/relocatable-perl/releases/download/5.36.0.1/perl-linux-arm64.tar.xz
```

# Author

Shoichi Kaji

# License

MIT

<!-- togo-header -->
<div align="center">
  <picture><source media="(prefers-color-scheme: dark)" srcset=".github/assets/togo-mark-dark.svg" /><img src=".github/assets/togo-mark.svg" alt="ToGO" height="64" /></picture>
  <h1>togo-framework/sharing</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1F8A99" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/sharing"><img src="https://pkg.go.dev/badge/github.com/togo-framework/sharing.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/sharing
```

<!-- /togo-header -->

togo's zero-copy **data-sharing** capability — a client for the open
**OpenSharing / Delta Sharing** REST protocol. Consume datasets shared by
partners without copying data.

```bash
togo install togo-framework/sharing
togo config:set SHARING_ENDPOINT https://share.example.com/delta-sharing
```

```go
s := sharing.FromKernel(k)
shares, _ := s.Shares(ctx)
```

| Key | Purpose |
|---|---|
| `SHARING_ENDPOINT` | Delta Sharing endpoint |
| `SHARING_TOKEN` | bearer token (secret) |
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
